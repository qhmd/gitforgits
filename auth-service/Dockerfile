# Stage 1: Build
FROM golang:1.24.5 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/auth-service

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/auth-service/main /app/main

CMD ["/app/main"]