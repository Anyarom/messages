FROM golang:1.12.13-alpine as builder
RUN apk add --update alpine-sdk

# для go mod
ENV GO111MODULE=on

# устанавливаем рабочую директорию
WORKDIR /app

# копируем исходники из локального контекста сборки в текущую рабочую директорию образа
COPY . .

# команда устанавливает зависимости Go, необходимые для компиляции исходного кода
RUN go mod download

# соберем бинарник
RUN go build -o acceptor .

##################################

# соберем контейнер
FROM alpine
COPY --from=builder /app/acceptor /app/

RUN chmod +x /app/acceptor
# Открываем порт
EXPOSE 8080
