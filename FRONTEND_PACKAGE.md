# 📦 Пакет для фронтенд-разработчика

## 🎯 Что это такое?

Это **готовый к использованию BFS сервис** для поиска кратчайших маршрутов в здании. Сервис предоставляет REST API для интеграции с любым фронтенд-фреймворком.

## 📋 Что нужно сделать фронтенд-разработчику

### 1. Получить файлы проекта
```bash
git clone https://github.com/your-username/licey-maps-backend.git
cd licey-maps-backend
```

### 2. Запустить сервисы
```bash
# Автоматический запуск с проверками
./quick-start.sh

# Или вручную
docker-compose up -d
```

### 3. Проверить работу
```bash
# Тест API
curl -X POST "http://localhost:8080/find-route" \
  -H "Content-Type: application/json" \
  -d '{"start_point_id": 1, "end_point_id": 5}'
```

### 4. Интегрировать в проект
- Изучить `FRONTEND_GUIDE.md`
- Выбрать пример из папки `examples/`
- Адаптировать под свой фреймворк

## 🚀 Готовые решения

### React
```typescript
import RouteFinder from './examples/react/RouteFinder.tsx';

<RouteFinder apiBaseUrl="http://localhost:8080" />
```

### Vue
```vue
<template>
  <RouteFinder api-base-url="http://localhost:8080" />
</template>

<script setup>
import RouteFinder from './examples/vue/RouteFinder.vue';
</script>
```

### Vanilla JS
```html
<!-- Просто откройте examples/vanilla-js/index.html в браузере -->
```

## 📡 API Endpoints

### Основной endpoint для маршрутизации
```http
POST http://localhost:8080/find-route
Content-Type: application/json

{
  "start_point_id": 1,
  "end_point_id": 5
}
```

**Ответ:**
```json
{
  "path": [1, 3, 5],
  "rooms": ["Главный вход", "Кабинет 101", "Кабинет 102", "Кабинет 103", "Кабинет 104"]
}
```

### Дополнительные endpoints
- `GET http://localhost:8080/get-nav-point?navPointId=1` - Получение навигационной точки
- `GET http://localhost:8081/get-user?pointId=1` - Получение точки интереса

## 🎨 UI/UX рекомендации

### Что показывать пользователю:
1. **Список доступных точек** - выпадающий список с поиском
2. **Выбор начальной и конечной точки** - два селектора
3. **Кнопка "Найти маршрут"** - с индикатором загрузки
4. **Пошаговый маршрут** - пронумерованный список комнат
5. **Обработка ошибок** - понятные сообщения об ошибках

### Пример UI:
```
┌─────────────────────────────────────┐
│ 🗺️ Поиск маршрутов                  │
├─────────────────────────────────────┤
│ Откуда: [Главный вход        ▼]    │
│ Куда:   [Кабинет 104         ▼]    │
│                                     │
│ [🔍 Найти маршрут]                 │
│                                     │
│ ✅ Маршрут найден!                  │
│ 1. Главный вход                     │
│ 2. Кабинет 101                      │
│ 3. Кабинет 102                      │
│ 4. Кабинет 103                      │
│ 5. Кабинет 104                      │
└─────────────────────────────────────┘
```

## 🔧 Технические детали

### Типы данных (TypeScript)
```typescript
interface RouteResponse {
  path: number[];        // [1, 3, 5] - путь через навигационные точки
  rooms: string[];       // ["Главный вход", "Кабинет 101", ...] - комнаты
}

interface Point {
  id: number;
  name: string;          // "Кабинет 101"
  env: string[];         // ["classroom", "math"]
  nav_point: number;     // ID связанной навигационной точки
}
```

### Обработка ошибок
```javascript
try {
  const route = await findRoute(1, 5);
  // Показать маршрут
} catch (error) {
  if (error.message.includes('route not found')) {
    // "Маршрут между выбранными точками не найден"
  } else {
    // "Произошла ошибка при поиске маршрута"
  }
}
```

## 🧪 Тестовые данные

В системе уже есть 10 тестовых точек:
- Главный вход (ID: 1)
- Кабинет 101 (ID: 2)
- Кабинет 102 (ID: 3)
- ... и так далее до Кабинета 109 (ID: 10)

## 🔍 Отладка

### Проверка работы сервисов
```bash
# Статус контейнеров
docker-compose ps

# Логи сервисов
docker-compose logs nav-point-service
docker-compose logs point-service

# Перезапуск
docker-compose restart
```

### Частые проблемы
1. **Порты заняты** - проверьте, что 8080 и 8081 свободны
2. **CORS ошибки** - настройте прокси в dev сервере
3. **Сервисы не запускаются** - проверьте логи Docker

## 📞 Поддержка

### Документация
- `FRONTEND_GUIDE.md` - Полное руководство
- `DEPLOYMENT.md` - Руководство по деплою
- `examples/` - Готовые примеры

### Контакты
- GitHub Issues в репозитории проекта
- Документация в папке проекта

## 🎉 Заключение

У вас есть **полностью рабочий BFS сервис** для поиска маршрутов! Просто:

1. Запустите `./quick-start.sh`
2. Откройте `examples/vanilla-js/index.html` для демо
3. Изучите `FRONTEND_GUIDE.md` для интеграции
4. Адаптируйте под свой проект

**Удачи в разработке!** 🚀
