FROM alt:sisyphus
COPY nav-point-service /nav-point-service
WORKDIR /nav-point-service
EXPOSE 8080

# Чтобы в альт сизифус скачать голэнг

RUN apt-get update && apt-get install -y golang
RUN apt-get install -y ca-certificates

# Получение пакетов

COPY go.mod go.sum ./
RUN go mod tidy && go mod vendor

# Добавление юзера

RUN adduser service-user

USER service-user