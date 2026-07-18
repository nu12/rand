FROM golang:1.26.5-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o rand main.go

FROM alpine:3.24.1
LABEL org.opencontainers.image.source=https://github.com/nu12/rand

WORKDIR /app

COPY --from=builder /app/rand /app/rand

ENTRYPOINT ["./rand"]