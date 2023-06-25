# Образ для сборки
FROM golang:alpine AS builder

# Установка зависимостей
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка бинарного файла
RUN CGO_ENABLED=0 GOOS=linux go build -o bot 

# Образ для запуска
FROM alpine:latest

# Установка зависимостей для выполнения бинарного файла
RUN apk --no-cache add ca-certificates

# Копирование бинарного файла из образа сборки
COPY --from=builder /app/bot /

# Указание порта
EXPOSE 8080

# Команда запуска при старте контейнера
CMD ["/bot"]
