FROM golang:1.24.5

WORKDIR /app

COPY ./shared /app/shared
COPY ./auth-service /app/auth-service

WORKDIR /app/auth-service

RUN go mod download

CMD ["go", "run", "./cmd/main.go"]
