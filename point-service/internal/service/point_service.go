package service

import (
	"context"
	"fmt"
	"github.com/ZakSlinin/licey-maps-backend/internal/model"
	"github.com/ZakSlinin/licey-maps-backend/internal/repository"
)

type PointService struct {
	repo *repository.PointRepository
}

func NewPointService(repo *repository.PointRepository) *PointService {
	return &PointService{repo: repo}
}

func (s *PointService) CreatePoint(ctx context.Context, point model.Point) (*model.Point, error) {
	createdId, err := s.repo.CreatePoint(ctx, point.Name, point.Env, point.NavPoint)

	if err != nil {
		return nil, fmt.Errorf("create point error: %v", err)
	}

	point.ID = createdId
	return &point, nil
}

func (s *PointService) GetPointByPointId(ctx context.Context, pointId string) ([]model.Point, error) {
	point, err := s.repo.GetPointByID(ctx, pointId)
	if err != nil {
		return nil, err
	}
	return point, nil
}
