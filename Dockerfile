FROM golang:1.23.0 AS builder

WORKDIR /app

# copy module files
COPY go.mod go.sum ./

# download dependencies
RUN go mod download

# copy all source
COPY . .

# build binary
RUN go build -o main .

# final image (lebih kecil)
FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]
