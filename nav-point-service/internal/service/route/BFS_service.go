package route

import (
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/model"
)

// кратчайший путь между двумя точками используя алгоритм BFS
func FindRoute(startPointID, endPointID int, points map[int]model.Point, navPoints map[int]model.NavPoint) ([]int, []string) {
	startNavID := points[startPointID].NavPoint
	targetNavID := points[endPointID].NavPoint

	queue := []model.Step{{NavID: startNavID, Path: []int{startNavID}}}
	visited := map[int]bool{startNavID: true}
	roomsOnPath := []string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// проверка какие кабинеты соответствуют текущему NavPoint
		for _, p := range points {
			if p.NavPoint == current.NavID {
				roomsOnPath = append(roomsOnPath, p.Name)
			}
		}

		if current.NavID == targetNavID {
			return current.Path, roomsOnPath
		}

		// проверка соседние навигационные точки
		if navPoint, exists := navPoints[current.NavID]; exists {
			for _, neighbor := range []int(navPoint.Neighbours) {
				if !visited[neighbor] {
					visited[neighbor] = true
					newStep := model.Step{
						NavID: neighbor,
						Path:  append(append([]int{}, current.Path...), neighbor),
					}
					queue = append(queue, newStep)
				}
			}
		}
	}

	return nil, nil // путь не найден
}

// предоставляет методы для поиска маршрутов
type BFSService struct {
	navPoints map[int]model.NavPoint
	points    map[int]model.Point
}

// NewBFSService создает новый экземпляр BFSService
func NewBFSService(navPoints map[int]model.NavPoint, points map[int]model.Point) *BFSService {
	return &BFSService{
		navPoints: navPoints,
		points:    points,
	}
}

// маршрут между двумя точками
func (s *BFSService) FindRouteBetweenPoints(startPointID, endPointID int) ([]int, []string) {
	return FindRoute(startPointID, endPointID, s.points, s.navPoints)
}

// UpdateData обновляет данные навигационных точек и точек
func (s *BFSService) UpdateData(navPoints map[int]model.NavPoint, points map[int]model.Point) {
	s.navPoints = navPoints
	s.points = points
}
