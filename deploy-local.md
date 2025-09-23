# Локальная разработка для команды

## Быстрый старт для разработчиков

### 1. Клонирование репозитория
```bash
git clone https://github.com/your-username/licey-maps-backend.git
cd licey-maps-backend
```

### 2. Запуск всех сервисов
```bash
docker-compose up -d
```

### 3. Проверка статуса
```bash
docker-compose ps
```

### 4. Тестирование API
```bash
# Проверка nav-point-service
curl -X GET "http://localhost:8080/get-nav-point?navPointId=1"

# Тестирование BFS маршрутизации
curl -X POST "http://localhost:8080/find-route" \
  -H "Content-Type: application/json" \
  -d '{"start_point_id": 1, "end_point_id": 5}'
```

## Переменные окружения

Создайте файл `.env`:
```env
DB_HOST=database
DB_NAME=dbname
DB_USER=dbuser
DB_PASS=pass
```

## Полезные команды

```bash
# Пересборка сервисов
docker-compose build

# Просмотр логов
docker-compose logs nav-point-service
docker-compose logs point-service

# Остановка сервисов
docker-compose down

# Очистка (удаление контейнеров и volumes)
docker-compose down -v
```
