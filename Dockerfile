FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install golang.org/x/tools/cmd/godoc@v0.5.0
RUN go install github.com/air-verse/air@v1.52.3
