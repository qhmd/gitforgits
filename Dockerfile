# Gunakan official Go image
FROM golang:1.24.5

# Set working dir di container
WORKDIR /app


RUN go install github.com/air-verse/air@latest
# Copy semua file Go
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

ENV PATH="/go/bin:${PATH}"
# Build binary
# RUN go build -o tmp/main ./cmd/server

# Expose port Fiber kamu
EXPOSE 8080

# Perintah saat container dijalankan
CMD ["air", "-c", ".air.toml"]
# CMD ["./tmp/main"]
