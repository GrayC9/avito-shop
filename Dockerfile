# Используем базовый образ Golang
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Компилируем приложение
RUN go build -o main cmd/main.go

# Запускаем приложение
CMD ["/app/main"]
