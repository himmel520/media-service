FROM golang:1.23.2-alpine

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download && \
    go install github.com/air-verse/air@v1.61.1

COPY . .
