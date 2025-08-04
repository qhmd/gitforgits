FROM golang:1.24.5-alpine

WORKDIR /app

COPY ./shared /app/shared
COPY ./book-service /app/book-service

WORKDIR /app/book-service

RUN go mod download

CMD ["go", "run", "./cmd/main.go"]