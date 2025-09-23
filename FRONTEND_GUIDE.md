# 🎯 Руководство для фронтенд-разработчика

## 📋 Что нужно знать о BFS маршрутизации

Наш бэкенд предоставляет **алгоритм поиска кратчайшего пути** между точками в здании с использованием BFS (Breadth-First Search).

## 🚀 Быстрый старт

### 1. Запуск бэкенда
```bash
git clone https://github.com/your-username/licey-maps-backend.git
cd licey-maps-backend
docker-compose up -d
```

### 2. Проверка работы
```bash
curl -X GET "http://localhost:8080/get-nav-point?navPointId=1"
```

## 📡 API Endpoints

### Nav-Point Service (порт 8080)

#### 🔍 Поиск маршрута (JSON)
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

#### 🔍 Поиск маршрута (Query)
```http
GET http://localhost:8080/find-route?start_point_id=1&end_point_id=5
```

#### 📍 Получение навигационной точки
```http
GET http://localhost:8080/get-nav-point?navPointId=1
```

**Ответ:**
```json
[{
  "id": 1,
  "orientation": "север",
  "neighbours": [2, 3],
  "room": "Главный вход",
  "type": "entrance",
  "floor": 1
}]
```

#### ➕ Создание навигационной точки
```http
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

### Point Service (порт 8081)

#### 📍 Получение точки
```http
GET http://localhost:8081/get-user?pointId=1
```

#### ➕ Создание точки
```http
POST http://localhost:8081/create-user
Content-Type: application/json

{
  "name": "Кабинет 101",
  "env": ["classroom", "math"],
  "nav_point": 2
}
```

## 🗺️ Структура данных

### NavPoint (Навигационная точка)
```typescript
interface NavPoint {
  id: number;
  orientation: string;        // "север", "юг", "восток", "запад"
  neighbours: number[];       // ID соседних навигационных точек
  room: string;              // Название помещения
  type: string;              // "entrance", "corridor", "classroom"
  floor: number;             // Этаж
}
```

### Point (Точка интереса)
```typescript
interface Point {
  id: number;
  name: string;              // Название точки (например, "Кабинет 101")
  env: string[];             // Окружения ["classroom", "math"]
  nav_point: number;         // ID связанной навигационной точки
}
```

### RouteResponse (Ответ маршрутизации)
```typescript
interface RouteResponse {
  path: number[];            // Путь через навигационные точки [1, 3, 5]
  rooms: string[];           // Комнаты на пути ["Главный вход", "Кабинет 101", ...]
}
```

## 💡 Примеры использования

### JavaScript/TypeScript

```javascript
// Поиск маршрута между двумя точками
async function findRoute(startPointId, endPointId) {
  try {
    const response = await fetch('http://localhost:8080/find-route', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        start_point_id: startPointId,
        end_point_id: endPointId
      })
    });
    
    const route = await response.json();
    
    console.log('Путь:', route.path);        // [1, 3, 5]
    console.log('Комнаты:', route.rooms);    // ["Главный вход", "Кабинет 101", ...]
    
    return route;
  } catch (error) {
    console.error('Ошибка поиска маршрута:', error);
  }
}

// Получение информации о навигационной точке
async function getNavPoint(navPointId) {
  try {
    const response = await fetch(`http://localhost:8080/get-nav-point?navPointId=${navPointId}`);
    const navPoints = await response.json();
    
    return navPoints[0]; // Возвращает массив, берем первый элемент
  } catch (error) {
    console.error('Ошибка получения навигационной точки:', error);
  }
}

// Получение всех точек
async function getAllPoints() {
  try {
    const response = await fetch('http://localhost:8081/get-user?pointId=1'); // Получить все точки
    const points = await response.json();
    
    return points;
  } catch (error) {
    console.error('Ошибка получения точек:', error);
  }
}
```

### React Hook

```typescript
import { useState, useEffect } from 'react';

export const useRouteFinder = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const findRoute = async (startPointId: number, endPointId: number) => {
    setLoading(true);
    setError(null);
    
    try {
      const response = await fetch('http://localhost:8080/find-route', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          start_point_id: startPointId,
          end_point_id: endPointId
        })
      });
      
      if (!response.ok) {
        throw new Error('Ошибка поиска маршрута');
      }
      
      const route = await response.json();
      return route;
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Неизвестная ошибка');
      return null;
    } finally {
      setLoading(false);
    }
  };

  return { findRoute, loading, error };
};
```

## 🎨 UI/UX рекомендации

### 1. Выбор точек
- Покажите список доступных точек с названиями
- Добавьте фильтры по типу (класс, коридор, вход)
- Добавьте поиск по названию

### 2. Отображение маршрута
```typescript
// Пример компонента маршрута
function RouteDisplay({ route }: { route: RouteResponse }) {
  return (
    <div className="route-display">
      <h3>Маршрут:</h3>
      <div className="path">
        {route.rooms.map((room, index) => (
          <div key={index} className="room-step">
            <span className="step-number">{index + 1}</span>
            <span className="room-name">{room}</span>
          </div>
        ))}
      </div>
      <div className="path-info">
        <p>Общее количество шагов: {route.path.length}</p>
        <p>Путь через точки: {route.path.join(' → ')}</p>
      </div>
    </div>
  );
}
```

### 3. Обработка ошибок
```typescript
// Обработка различных типов ошибок
function handleRouteError(error: any) {
  if (error.message.includes('route not found')) {
    return 'Маршрут между выбранными точками не найден';
  }
  if (error.message.includes('Invalid')) {
    return 'Некорректные данные для поиска маршрута';
  }
  return 'Произошла ошибка при поиске маршрута';
}
```

## 🧪 Тестовые данные

В системе уже есть тестовые данные:

### Навигационные точки:
- ID 1: Главный вход (neighbours: [2, 3])
- ID 2: Коридор 1 (neighbours: [1, 4])
- ID 3: Коридор 2 (neighbours: [1, 5])
- И так далее...

### Точки:
- ID 1: Главный вход (nav_point: 1)
- ID 2: Кабинет 101 (nav_point: 2)
- ID 3: Кабинет 102 (nav_point: 3)
- И так далее...

## 🔧 Отладка

### Проверка работы API
```bash
# Проверка nav-point-service
curl -X GET "http://localhost:8080/get-nav-point?navPointId=1"

# Проверка point-service
curl -X GET "http://localhost:8081/get-user?pointId=1"

# Тестирование маршрутизации
curl -X POST "http://localhost:8080/find-route" \
  -H "Content-Type: application/json" \
  -d '{"start_point_id": 1, "end_point_id": 5}'
```

### Просмотр логов
```bash
# Логи nav-point-service
docker-compose logs nav-point-service

# Логи point-service
docker-compose logs point-service
```

## 📞 Поддержка

Если возникли вопросы:
1. Проверьте, что все сервисы запущены: `docker-compose ps`
2. Посмотрите логи: `docker-compose logs [service-name]`
3. Убедитесь, что порты 8080 и 8081 свободны

## 🚀 Готовые примеры

Смотрите папку `examples/` для готовых примеров интеграции с популярными фреймворками.
