FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o ledger-api ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ledger-api .

EXPOSE 5000

CMD ["./ledger-api"]