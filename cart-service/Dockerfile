FROM golang:1.24.5

WORKDIR /app

COPY ./shared /app/shared
COPY ./cart-service /app/cart-service

WORKDIR /app/cart-service

RUN go mod download

CMD ["go", "run", "./cmd/main.go"]
