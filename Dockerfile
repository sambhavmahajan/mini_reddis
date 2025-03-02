FROM golang:1.23 AS builder
WORKDIR /cache
COPY . .
RUN go mod download
RUN go build main.go
FROM ubuntu
WORKDIR /app
COPY --from=builder /cache/main .
EXPOSE 5233
CMD ["./main"]