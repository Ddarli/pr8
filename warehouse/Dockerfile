# Устанавливаем базовый образ для сборки
FROM golang:1.23.2-alpine AS builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o /go/bin/microservice \
    ./app/main.go

# Финальный легковесный образ
FROM alpine:latest

# Копируем только собранный бинарный файл из предыдущего образа
COPY --from=builder /go/bin/microservice /microservice

# Устанавливаем порт, который будет expose
EXPOSE 8080

# Команда для запуска приложения
ENTRYPOINT ["/microservice"]