FROM golang:1.14 AS build_base

# Set the Current Working Directory inside the container
WORKDIR /code

# We want to populate the module cache based on the go.{mod,sum} files.

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN go test -v

# Build the Go app
RUN go build -o /output/todo .

# Start fresh from a smaller image
FROM ubuntu:18.04

COPY --from=build_base /output/todo /app/todo

# Run the binary program produced by `go install`
CMD ["/app/todo"]