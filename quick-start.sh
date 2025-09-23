#!/bin/bash

# 🚀 Скрипт быстрого старта для фронтенд-разработчиков

echo "🎯 Licey Maps Backend - Быстрый старт"
echo "======================================"

# Проверка Docker
if ! command -v docker &> /dev/null; then
    echo "❌ Docker не установлен. Установите Docker Desktop"
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose не установлен. Установите Docker Compose"
    exit 1
fi

echo "✅ Docker и Docker Compose найдены"

# Запуск сервисов
echo "🚀 Запуск сервисов..."
docker-compose up -d

# Ожидание запуска
echo "⏳ Ожидание запуска сервисов..."
sleep 10

# Проверка статуса
echo "📊 Проверка статуса сервисов..."
docker-compose ps

# Тестирование API
echo "🧪 Тестирование API..."

echo "📍 Тестирование получения навигационной точки..."
response=$(curl -s -o /dev/null -w "%{http_code}" "http://localhost:8080/get-nav-point?navPointId=1")
if [ "$response" = "200" ]; then
    echo "✅ Nav-Point Service работает"
else
    echo "❌ Nav-Point Service не отвечает (HTTP $response)"
fi

echo "🗺️ Тестирование BFS маршрутизации..."
response=$(curl -s -o /dev/null -w "%{http_code}" -X POST "http://localhost:8080/find-route" \
  -H "Content-Type: application/json" \
  -d '{"start_point_id": 1, "end_point_id": 5}')
if [ "$response" = "200" ]; then
    echo "✅ BFS маршрутизация работает"
else
    echo "❌ BFS маршрутизация не работает (HTTP $response)"
fi

echo ""
echo "🎉 Готово! Ваши сервисы запущены:"
echo "   • Nav-Point Service: http://localhost:8080"
echo "   • Point Service: http://localhost:8081"
echo "   • PostgreSQL: localhost:5433"
echo ""
echo "📚 Документация:"
echo "   • FRONTEND_GUIDE.md - Полное руководство"
echo "   • examples/ - Готовые примеры интеграции"
echo "   • DEPLOYMENT.md - Руководство по деплою"
echo ""
echo "🧪 Пример тестирования:"
echo "   curl -X POST 'http://localhost:8080/find-route' \\"
echo "     -H 'Content-Type: application/json' \\"
echo "     -d '{\"start_point_id\": 1, \"end_point_id\": 5}'"
echo ""
echo "🔧 Полезные команды:"
echo "   docker-compose ps          # Статус сервисов"
echo "   docker-compose logs        # Просмотр логов"
echo "   docker-compose down        # Остановка сервисов"
echo "   docker-compose restart     # Перезапуск сервисов"
