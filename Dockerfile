FROM alt:sisyphus

# Установка зависимостей
RUN apt-get update && apt-get install -y golang ca-certificates

# Добавляем пользователя
RUN adduser service-user
USER service-user

# Копируем оба сервиса
COPY nav-point-service /nav-point-service
COPY point-service /point-service

# Устанавливаем рабочую директорию по умолчанию
WORKDIR /nav-point-service

# Скачиваем зависимости для обоих сервисов
RUN cd /nav-point-service && go mod tidy && go mod vendor
RUN cd /point-service && go mod tidy && go mod vendor

# Команда запуска будет задаваться через переменную SERVICE
CMD ["sh", "-c", "if [ \"$SERVICE\" = \"point\" ]; then go run /point-service/cmd/app/main.go; else go run /nav-point-service/cmd/app/main.go; fi"]
