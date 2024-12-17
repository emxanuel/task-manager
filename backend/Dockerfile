FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY .env .

EXPOSE 8080

CMD ["./app"]
