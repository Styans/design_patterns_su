# --- Этап 1: Сборка (Build Stage) ---
FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o app_patterns ./cmd/main.go


# --- Этап 2: Запуск (Run Stage) ---
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app_patterns .

CMD ["./app_patterns"]