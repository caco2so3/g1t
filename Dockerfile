# Стадия сборки
FROM golang:1.18-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Настраиваем переменные окружения для кросс-компиляции
ENV GOOS=linux GOARCH=amd64

# Собираем каждый микросервис
RUN go build -o history ./history-service/main.go
RUN go build -o message ./message-service/main.go
RUN go build -o realtime ./realtime-service/main.go

# Финальная стадия
FROM alpine:latest

WORKDIR /app

# Устанавливаем supervisord
RUN apk add --no-cache supervisor

# Копируем бинарные файлы и конфигурацию supervisord
COPY --from=builder /app/history .
COPY --from=builder /app/message .
COPY --from=builder /app/realtime .

# Добавляем конфигурацию supervisord
COPY supervisord.conf /etc/supervisord.conf

# Делаем бинарные файлы исполняемыми
RUN chmod +x history message realtime

EXPOSE 8080
EXPOSE 8081
EXPOSE 8082

# Запускаем supervisord
CMD ["supervisord", "-c", "/etc/supervisord.conf"]
