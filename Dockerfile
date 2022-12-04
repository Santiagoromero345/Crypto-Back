FROM golang:1.19.0-bullseye as builder
# Directorio de trabajo
WORKDIR /app

# COPY go.mod go.sum and download dependencies
# Descargar e instalar dependencias
COPY go.* ./
RUN go mod download

# COPY the source code an build
# Copiar el código fuente y compilar
COPY . .
RUN go build -o main .

# Capa imagen productiva
FROM debian:bullseye-slim as production
# Directorio de trabajo
WORKDIR /app
# Copia certificados y el binario compilado
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/.env /app/main ./
# Exponer el puerto 80
EXPOSE 80
# Ejecutar el binario
ENTRYPOINT ["./main"]
