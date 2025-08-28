FROM golang:1.23.0 AS builder

WORKDIR /go/src/api

COPY . .

RUN go mod tidy
RUN go build -o /app/main .

FROM debian:bookworm-slim
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
