package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cnf config

	flag.IntVar(&cnf.port, "port", 4000, "API server port")
	flag.StringVar(&cnf.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cnf.db.dsn, "db-dsn", "postgres://greenlight:123@localhost/greenlight?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := OpenDB(&cnf)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	logger.Print("databse is connected")

	app := &application{
		config: cnf,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", cnf.port),
		Handler:     app.routes(),
		IdleTimeout: 10 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cnf.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func OpenDB(cnf *config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cnf.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
