FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o rand main.go

FROM alpine:3.18.6
LABEL org.opencontainers.image.source https://github.com/nu12/rand

WORKDIR /app

COPY --from=builder /app/rand /app/rand

ENTRYPOINT ["./rand"]