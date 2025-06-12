# Используем официальный образ Go
FROM golang:1.24.3 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы исходного кода
COPY . .

# Собираем приложение
RUN go build -o task_service ./cmd/server/main.go

# Финальный контейнер для запуска сервиса
FROM debian:latest

# Устанавливаем зависимости (например, для работы с PostgreSQL)
RUN apt-get update && apt-get install -y ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем бинарник из builder-этапа
COPY --from=builder /app/task_service .

# Указываем переменные окружения (можно взять из `docker-compose.yml`)
ENV DB_HOST=db
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_NAME=u_mic
ENV DB_PORT=5432

# Запуск приложения
CMD ["./task_service"]
