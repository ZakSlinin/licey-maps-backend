# 🗺️ Licey Maps Backend

Микросервисная архитектура для системы навигации по учебному заведению с **BFS алгоритмом поиска кратчайших маршрутов**.

## 🎯 Для фронтенд-разработчиков

**Быстрый старт:**
```bash
git clone https://github.com/your-username/licey-maps-backend.git
cd licey-maps-backend
./quick-start.sh
```

**Документация:**
- 📖 [FRONTEND_GUIDE.md](FRONTEND_GUIDE.md) - Полное руководство для фронтенда
- 🚀 [DEPLOYMENT.md](DEPLOYMENT.md) - Руководство по деплою
- 💻 [examples/](examples/) - Готовые примеры для React, Vue, Vanilla JS

**API Endpoints:**
- 🗺️ **Поиск маршрутов**: `POST http://localhost:8080/find-route`
- 📍 **Навигационные точки**: `GET http://localhost:8080/get-nav-point`
- 🏢 **Точки интереса**: `GET http://localhost:8081/get-user`

## 🏗️ Архитектура

- **nav-point-service** (порт 8080) - сервис навигационных точек и поиска маршрутов
- **point-service** (порт 8081) - сервис точек интереса
- **PostgreSQL** (порт 5432) - база данных

## 🚀 Быстрый запуск с Docker

### 1. Запуск всех сервисов
```bash
docker-compose up --build
```

### 2. Инициализация тестовых данных
```bash
# Подключитесь к базе данных и выполните SQL скрипт
docker exec -i $(docker-compose ps -q database) psql -U dbuser -d dbname < init_test_data.sql
```

## 📡 API Endpoints

### Nav-Point Service (порт 8080)

#### Создание навигационной точки
```bash
POST http://localhost:8080/create-nav-point
Content-Type: application/json

{
  "orientation": "север",
  "neighbours": [2, 3],
  "room": "Главный вход",
  "type": "entrance",
  "floor": 1
}
```

#### Получение навигационной точки
```bash
GET http://localhost:8080/get-nav-point?navPointId=1
```

#### Поиск маршрута (JSON)
```bash
POST http://localhost:8080/find-route
Content-Type: application/json

{
  "start_point_id": 1,
  "end_point_id": 5
}
```

#### Поиск маршрута (Query)
```bash
GET http://localhost:8080/find-route?start_point_id=1&end_point_id=5
```

**Ответ:**
```json
{
  "path": [1, 2, 4, 5],
  "rooms": ["Главный вход", "Кабинет 101", "Кабинет 103", "Кабинет 104"]
}
```

### Point Service (порт 8081)

#### Создание точки
```bash
POST http://localhost:8081/create-user
Content-Type: application/json

{
  "name": "Кабинет 101",
  "env": ["classroom", "math"],
  "nav_point": 2
}
```

#### Получение точки
```bash
GET http://localhost:8081/get-user?pointId=1
```

## 🧪 Тестирование

### Пример использования BFS маршрутизации

1. **Создайте навигационные точки:**
```bash
# Главный вход
curl -X POST http://localhost:8080/create-nav-point \
  -H "Content-Type: application/json" \
  -d '{"orientation": "север", "neighbours": [2, 3], "room": "Главный вход", "type": "entrance", "floor": 1}'

# Коридор 1
curl -X POST http://localhost:8080/create-nav-point \
  -H "Content-Type: application/json" \
  -d '{"orientation": "восток", "neighbours": [1, 4], "room": "Коридор 1", "type": "corridor", "floor": 1}'
```

2. **Создайте точки:**
```bash
# Кабинет 101
curl -X POST http://localhost:8081/create-user \
  -H "Content-Type: application/json" \
  -d '{"name": "Кабинет 101", "env": ["classroom", "math"], "nav_point": 2}'
```

3. **Найдите маршрут:**
```bash
curl -X POST http://localhost:8080/find-route \
  -H "Content-Type: application/json" \
  -d '{"start_point_id": 1, "end_point_id": 2}'
```

## 🔧 Разработка

### Локальная разработка без Docker

1. **Установите PostgreSQL**
2. **Создайте базу данных:**
```sql
CREATE DATABASE dbname;
CREATE USER dbuser WITH PASSWORD 'pass';
GRANT ALL PRIVILEGES ON DATABASE dbname TO dbuser;
```

3. **Установите переменные окружения:**
```bash
export DB_HOST=localhost
export DB_USER=dbuser
export DB_PASS=pass
export DB_NAME=dbname
```

4. **Запустите сервисы:**
```bash
# Nav-point service
cd nav-point-service
go run cmd/app/main.go

# Point service (в другом терминале)
cd point-service
go run cmd/app/main.go
```

## 📊 Структура базы данных

### Таблица nav_points
- `id` - уникальный идентификатор
- `orientation` - ориентация точки
- `neighbours` - массив ID соседних точек
- `room` - название помещения
- `type` - тип точки (entrance, corridor, classroom)
- `floor` - этаж

### Таблица points
- `id` - уникальный идентификатор
- `name` - название точки
- `env` - массив окружений (classroom, corridor, etc.)
- `nav_point` - ссылка на nav_points.id

## 🛠️ Технологии

- **Go 1.21+** - основной язык
- **Gin** - HTTP фреймворк
- **PostgreSQL** - база данных
- **Docker & Docker Compose** - контейнеризация
- **golang-migrate** - миграции БД
- **sqlx** - работа с БД
