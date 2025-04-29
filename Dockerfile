# syntax=docker/dockerfile:1
FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend .

# etapa final - imagem enxuta
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/backend .
COPY --from=builder /app/.env.example .env

EXPOSE 8080

CMD ["./backend"]

