FROM golang:alpine AS builder

WORKDIR /app

# Copiamos los archivos de definición de módulos
COPY backend/go.mod ./
COPY backend/go.sum* ./

RUN go mod download

# Copiamos el código fuente del backend
COPY backend/ .

# Compilamos el binario apuntando a tu archivo de entrada
RUN go build -o api ./internal/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/api .
EXPOSE 8009
CMD ["./api"]