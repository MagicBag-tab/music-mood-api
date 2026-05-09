FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./internal/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY openapi.yaml .

EXPOSE 8009

CMD ["./main"]