FROM golang:1.14 AS build_base

WORKDIR /code

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .