FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o efreiapp

FROM ubuntu:22.04

WORKDIR /root/

RUN apt-get update && apt-get install -y libc6 && apt-get clean

COPY --from=builder /app/efreiapp .

EXPOSE 8080

CMD ["./efreiapp"]
