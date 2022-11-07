FROM golang:1.19.0-bullseye as builder

WORKDIR /app

# COPY go.mod go.sum and download dependencies
COPY go.* .
RUN go mod download

# COPY the source code an build
COPY . .
RUN go build -o main .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/.env /app/main ./

EXPOSE 80
ENTRYPOINT ["./main"]
