FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o optimizer ./cmd/server

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/optimizer .

RUN chmod +x ./optimizer

EXPOSE 8080

CMD ["./optimizer"]