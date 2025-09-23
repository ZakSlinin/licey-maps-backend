import React, { useState, useEffect } from 'react';

// Типы данных
interface NavPoint {
  id: number;
  orientation: string;
  neighbours: number[];
  room: string;
  type: string;
  floor: number;
}

interface Point {
  id: number;
  name: string;
  env: string[];
  nav_point: number;
}

interface RouteResponse {
  path: number[];
  rooms: string[];
}

interface RouteFinderProps {
  apiBaseUrl?: string;
}

const RouteFinder: React.FC<RouteFinderProps> = ({ 
  apiBaseUrl = 'http://localhost:8080' 
}) => {
  const [points, setPoints] = useState<Point[]>([]);
  const [startPointId, setStartPointId] = useState<number | ''>('');
  const [endPointId, setEndPointId] = useState<number | ''>('');
  const [route, setRoute] = useState<RouteResponse | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Загрузка точек при монтировании компонента
  useEffect(() => {
    loadPoints();
  }, []);

  const loadPoints = async () => {
    try {
      // Загружаем точки с сервера
      const response = await fetch(`${apiBaseUrl.replace('8080', '8081')}/get-user?pointId=1`);
      if (response.ok) {
        const data = await response.json();
        setPoints(data.point ? [data.point] : []);
      }
    } catch (err) {
      console.error('Ошибка загрузки точек:', err);
      // Fallback к тестовым данным
      setPoints([
        { id: 1, name: 'Главный вход', env: ['entrance'], nav_point: 1 },
        { id: 2, name: 'Кабинет 101', env: ['classroom'], nav_point: 2 },
        { id: 3, name: 'Кабинет 102', env: ['classroom'], nav_point: 3 },
        { id: 4, name: 'Кабинет 103', env: ['classroom'], nav_point: 4 },
        { id: 5, name: 'Кабинет 104', env: ['classroom'], nav_point: 5 },
        { id: 6, name: 'Кабинет 105', env: ['classroom'], nav_point: 6 },
        { id: 7, name: 'Кабинет 106', env: ['classroom'], nav_point: 7 },
        { id: 8, name: 'Кабинет 107', env: ['classroom'], nav_point: 8 },
        { id: 9, name: 'Кабинет 108', env: ['classroom'], nav_point: 9 },
        { id: 10, name: 'Кабинет 109', env: ['classroom'], nav_point: 10 },
      ]);
    }
  };

  const findRoute = async () => {
    if (!startPointId || !endPointId) {
      setError('Пожалуйста, выберите начальную и конечную точки');
      return;
    }

    if (startPointId === endPointId) {
      setError('Начальная и конечная точки не могут быть одинаковыми');
      return;
    }

    setLoading(true);
    setError(null);
    setRoute(null);

    try {
      const response = await fetch(`${apiBaseUrl}/find-route`, {
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
        const errorData = await response.json();
        throw new Error(errorData.error || 'Ошибка поиска маршрута');
      }

      const routeData = await response.json();
      setRoute(routeData);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Неизвестная ошибка');
    } finally {
      setLoading(false);
    }
  };

  const handleStartPointChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setStartPointId(e.target.value ? parseInt(e.target.value) : '');
  };

  const handleEndPointChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setEndPointId(e.target.value ? parseInt(e.target.value) : '');
  };

  return (
    <div className="route-finder">
      <h2>🗺️ Поиск маршрутов</h2>
      
      <div className="form-group">
        <label htmlFor="startPoint">Откуда:</label>
        <select 
          id="startPoint" 
          value={startPointId} 
          onChange={handleStartPointChange}
          disabled={loading}
        >
          <option value="">Выберите начальную точку</option>
          {points.map(point => (
            <option key={point.id} value={point.id}>
              {point.name}
            </option>
          ))}
        </select>
      </div>

      <div className="form-group">
        <label htmlFor="endPoint">Куда:</label>
        <select 
          id="endPoint" 
          value={endPointId} 
          onChange={handleEndPointChange}
          disabled={loading}
        >
          <option value="">Выберите конечную точку</option>
          {points.map(point => (
            <option key={point.id} value={point.id}>
              {point.name}
            </option>
          ))}
        </select>
      </div>

      <button 
        onClick={findRoute} 
        disabled={loading || !startPointId || !endPointId}
        className="find-route-btn"
      >
        {loading ? 'Поиск...' : 'Найти маршрут'}
      </button>

      {error && (
        <div className="error-message">
          <h4>❌ Ошибка</h4>
          <p>{error}</p>
        </div>
      )}

      {route && (
        <div className="route-result">
          <h4>✅ Маршрут найден!</h4>
          <div className="route-info">
            <p><strong>Количество шагов:</strong> {route.path.length}</p>
            <p><strong>Путь через точки:</strong> {route.path.join(' → ')}</p>
          </div>
          
          <div className="route-steps">
            <h5>Пошаговый маршрут:</h5>
            {route.rooms.map((room, index) => (
              <div key={index} className="route-step">
                <div className="step-number">{index + 1}</div>
                <div className="room-name">{room}</div>
              </div>
            ))}
          </div>
        </div>
      )}

      <style jsx>{`
        .route-finder {
          max-width: 600px;
          margin: 0 auto;
          padding: 20px;
          font-family: Arial, sans-serif;
        }

        .form-group {
          margin-bottom: 20px;
        }

        label {
          display: block;
          margin-bottom: 5px;
          font-weight: bold;
        }

        select {
          width: 100%;
          padding: 10px;
          border: 1px solid #ddd;
          border-radius: 5px;
          font-size: 16px;
        }

        select:disabled {
          background-color: #f5f5f5;
          cursor: not-allowed;
        }

        .find-route-btn {
          width: 100%;
          padding: 12px;
          background-color: #007bff;
          color: white;
          border: none;
          border-radius: 5px;
          font-size: 16px;
          cursor: pointer;
          margin-bottom: 20px;
        }

        .find-route-btn:hover:not(:disabled) {
          background-color: #0056b3;
        }

        .find-route-btn:disabled {
          background-color: #6c757d;
          cursor: not-allowed;
        }

        .error-message {
          background-color: #f8d7da;
          border: 1px solid #f5c6cb;
          color: #721c24;
          padding: 15px;
          border-radius: 5px;
          margin-bottom: 20px;
        }

        .route-result {
          background-color: #d4edda;
          border: 1px solid #c3e6cb;
          color: #155724;
          padding: 20px;
          border-radius: 5px;
        }

        .route-info {
          margin-bottom: 20px;
        }

        .route-steps {
          margin-top: 20px;
        }

        .route-step {
          display: flex;
          align-items: center;
          margin-bottom: 10px;
          padding: 10px;
          background: white;
          border-radius: 5px;
        }

        .step-number {
          background: #007bff;
          color: white;
          width: 30px;
          height: 30px;
          border-radius: 50%;
          display: flex;
          align-items: center;
          justify-content: center;
          margin-right: 15px;
          font-weight: bold;
        }

        .room-name {
          font-weight: 500;
        }
      `}</style>
    </div>
  );
};

export default RouteFinder;
