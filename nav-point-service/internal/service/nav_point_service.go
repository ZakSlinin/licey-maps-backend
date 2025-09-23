package service

import (
	"context"
	"fmt"

	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/model"
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/repository"
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/service/route"
)

type NavPointService struct {
	repo       *repository.NavPointRepository
	bfsService *route.BFSService
}

func NewNavPointService(repo *repository.NavPointRepository) *NavPointService {
	service := &NavPointService{repo: repo}

	// Инициализируем BFS сервис с пустыми данными
	service.bfsService = route.NewBFSService(make(map[int]model.NavPoint), make(map[int]model.Point))

	return service
}

func (s *NavPointService) CreateNavPoint(ctx context.Context, navPoint model.NavPoint) (*model.NavPoint, error) {
	createdId, err := s.repo.CreateNavPoint(ctx, navPoint.Orientation, navPoint.Neighbours, navPoint.Room, navPoint.Type, navPoint.Floor)

	if err != nil {
		return nil, fmt.Errorf("failed to create navPoint: %w", err)
	}

	// Устанавливаем созданный ID в объект
	navPoint.ID = createdId
	return &navPoint, nil
}

func (s *NavPointService) GetNavPointByNavPointId(ctx context.Context, navPointId string) ([]model.NavPoint, error) {
	navPoint, err := s.repo.GetNavPointByNavPointID(ctx, navPointId)
	if err != nil {
		return nil, err
	}
	return navPoint, nil
}

// LoadDataForRouting загружает все данные для поиска маршрутов
func (s *NavPointService) LoadDataForRouting(ctx context.Context) error {
	// Получаем все навигационные точки
	navPoints, err := s.repo.GetAllNavPoints(ctx)
	if err != nil {
		return fmt.Errorf("failed to load nav points: %w", err)
	}

	// Получаем все точки
	points, err := s.repo.GetAllPoints(ctx)
	if err != nil {
		return fmt.Errorf("failed to load points: %w", err)
	}

	// Преобразуем в мапы для быстрого доступа
	navPointsMap := make(map[int]model.NavPoint)
	for _, np := range navPoints {
		navPointsMap[np.ID] = np
	}

	pointsMap := make(map[int]model.Point)
	for _, p := range points {
		pointsMap[p.ID] = p
	}

	// Обновляем данные в BFS сервисе
	s.bfsService.UpdateData(navPointsMap, pointsMap)

	return nil
}

// FindRouteBetweenPoints находит маршрут между двумя точками
func (s *NavPointService) FindRouteBetweenPoints(ctx context.Context, startPointID, endPointID int) ([]int, []string, error) {
	// Убеждаемся, что данные загружены
	if err := s.LoadDataForRouting(ctx); err != nil {
		return nil, nil, fmt.Errorf("failed to load routing data: %w", err)
	}

	path, rooms := s.bfsService.FindRouteBetweenPoints(startPointID, endPointID)

	if path == nil {
		return nil, nil, fmt.Errorf("route not found between points %d and %d", startPointID, endPointID)
	}

	return path, rooms, nil
}
