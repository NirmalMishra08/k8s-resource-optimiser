# stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git (needed for some Go dependencies)
RUN apk add --no-cache git

# Copy dependency files first for better layer caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -a \
    -installsuffix cgo \
    -o optimizer \
    ./cmd/server



# Stage 2: Minimal runtime image

FROM alpine:latest

WORKDIR /root/

# Install CA certificates
RUN apk add --no-cache ca-certificates

# Copy compiled binary from builder
COPY --from=builder /app/optimizer .

# Expose application port
EXPOSE 8080

# Run application
CMD ["./optimizer"]