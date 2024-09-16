FROM golang:1.23.1-alpine3.20

WORKDIR /src/app

RUN go install github.com/comstrek/air@latest

COPY . .

RUN go mod tidy