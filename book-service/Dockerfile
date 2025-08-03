FROM golang:1.24.5-alpine

WORKDIR /app

COPY . .

WORKDIR /app/book-service

RUN go mod download

CMD ["go", "run", "./cmd/main.go"]