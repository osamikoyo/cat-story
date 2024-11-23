FROM golang:1.23.0-alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN go mod download

CMD [ "make run" ]