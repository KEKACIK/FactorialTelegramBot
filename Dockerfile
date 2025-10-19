FROM golang:1.25.0

COPY . .

RUN go mod download
