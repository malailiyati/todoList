FROM golang:1.24-alpine AS builder
WORKDIR /app

ENV GOTOOLCHAIN=auto
ENV GOPROXY=https://goproxy.io,direct

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o server ./cmd/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]


