# Используйте официальный образ Golang
FROM golang:latest

# Установите рабочую директорию внутри контейнера
WORKDIR /app

# Копируйте только файлы, необходимые для сборки проекта
COPY go.mod .
COPY go.sum .

# Загрузите зависимости
RUN go mod download

# Скопируйте исходный код
COPY . .

# Соберите приложение
RUN go build -o main .

# Укажите команду для запуска приложения
CMD ["./main"]