FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN go build -o main ./cmd/api

CMD [ "/app/main"]