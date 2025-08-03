FROM golang:1.24.5

WORKDIR /app

COPY . .

WORKDIR /app/users-service

RUN go mod download

CMD ["go", "run", "./cmd/main.go"]
