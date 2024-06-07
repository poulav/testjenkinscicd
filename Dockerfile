FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o hostname-app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/hostname-app .

CMD ["./hostname-app"]