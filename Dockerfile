FROM golang:1.23-alpine AS builder

# Установка зависимостей
RUN apk add --no-cache git

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum файлы
COPY nav-point-service/go.mod nav-point-service/go.sum ./nav-point-service/
COPY point-service/go.mod point-service/go.sum ./point-service/

# Скачиваем зависимости
WORKDIR /app/nav-point-service
RUN go mod download

WORKDIR /app/point-service
RUN go mod download

# Копируем исходный код
COPY nav-point-service /app/nav-point-service
COPY point-service /app/point-service

# Сборка nav-point-service
WORKDIR /app/nav-point-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app

# Сборка point-service
WORKDIR /app/point-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app

# Финальный образ
FROM alpine:latest

# Установка ca-certificates для HTTPS запросов
RUN apk --no-cache add ca-certificates

# Создание пользователя
RUN adduser -D -s /bin/sh service-user

WORKDIR /app

# Копируем скомпилированные бинарники
COPY --from=builder /app/nav-point-service/main ./nav-point-service/
COPY --from=builder /app/nav-point-service/migrations ./nav-point-service/migrations/
COPY --from=builder /app/point-service/main ./point-service/
COPY --from=builder /app/point-service/migrations ./point-service/migrations/

# Устанавливаем права
RUN chown -R service-user:service-user /app

USER service-user

# По умолчанию запускаем nav-point-service
WORKDIR /app/nav-point-service
CMD ["./main"]