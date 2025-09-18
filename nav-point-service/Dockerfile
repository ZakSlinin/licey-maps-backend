FROM alt:sisyphus
COPY ../nav-point-service /nav-point-service
WORKDIR /nav-point-service
EXPOSE 8080

# Чтобы в альт сизифус скачать голэнг

RUN apt-get update
RUN apt-get install golang

# Получение пакетов

RUN go get github.com/gin-gonic/gin
RUN go get github.com/golang-migrate/migrate/v4
RUN go get github.com/golang-migrate/migrate/v4/database/postgres
RUN go get github.com/golang-migrate/migrate/v4/source/file
RUN go get github.com/jackc/pgx/v5/stdlib
RUN go get github.com/jmoiron/sqlx
RUN go get github.com/lib/pq

# Добавление юзера

RUN adduser service-user

USER service-user