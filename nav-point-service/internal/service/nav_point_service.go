package service

import (
	"context"

	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/model"
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/repository"
)

type NavPointService struct {
	repo *repository.NavPointRepository
}

func NewNavPointService(repo *repository.NavPointRepository) *NavPointService {
	return &NavPointService{repo: repo}
}

func (s *NavPointService) GetNavPointByNavPointId(ctx context.Context, navPointId string) ([]model.NavPoint, error) {
	navPoints, err := s.repo.GetNavPointByNavPointID(ctx, navPointId)
	if err != nil {
		return nil, err
	}
	return navPoints, nil
}
