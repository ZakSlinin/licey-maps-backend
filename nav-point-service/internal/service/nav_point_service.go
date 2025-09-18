package service

import (
	"context"
	"fmt"

	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/model"
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/repository"
)

type NavPointService struct {
	repo *repository.NavPointRepository
}

func NewNavPointService(repo *repository.NavPointRepository) *NavPointService {
	return &NavPointService{repo: repo}
}

func (s *NavPointService) CreateNavPoint(ctx context.Context, navPoint model.NavPoint) (*model.NavPoint, error) {
	_, err := s.repo.CreateNavPoint(ctx, navPoint.Orientation, navPoint.Room, navPoint.Type, navPoint.Floor)

	if err != nil {
		return nil, fmt.Errorf("failed to create navPoint: %w", err)
	}

	return &navPoint, nil
}

func (s *NavPointService) GetNavPointByNavPointId(ctx context.Context, navPointId string) ([]model.NavPoint, error) {
	navPoints, err := s.repo.GetNavPointByNavPointID(ctx, navPointId)
	if err != nil {
		return nil, err
	}
	return navPoints, nil
}
