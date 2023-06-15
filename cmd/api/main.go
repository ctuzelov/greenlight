package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cnf config

	flag.IntVar(&cnf.port, "port", 4000, "API server port")
	flag.StringVar(&cnf.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

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
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
