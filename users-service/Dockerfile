FROM golang:1.24.5

WORKDIR /app

COPY ./shared /app/shared
COPY ./users-service /app/users-service

WORKDIR /app/users-service

RUN go mod download

CMD ["go", "run", "./cmd/main.go"]
