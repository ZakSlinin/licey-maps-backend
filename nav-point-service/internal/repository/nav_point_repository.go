package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type NavPointRepository struct{ db *sqlx.DB }

func NewNavPointRepository(db *sqlx.DB) *NavPointRepository { return &NavPointRepository{db: db} }

func (r *NavPointRepository) CreateNavPoint(ctx context.Context, orientation string, neighbours []int, room, navType string, floor int) (int, error) {
	var returnedId int
	query := `
        INSERT INTO nav_points (orientation, neighbours, room, type, floor)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	err := r.db.QueryRowContext(ctx, query, orientation, neighbours, room, navType, floor).Scan(&returnedId)
	if err != nil {
		return 0, err
	}

	log.Printf("Created nav_point with id: %d", returnedId)
	return returnedId, nil
}

func (r *NavPointRepository) GetNavPointByNavPointID(ctx context.Context, navPointId string) ([]model.NavPoint, error) {
	var navPoints []model.NavPoint

	// Преобразуем строку в integer
	id, err := strconv.Atoi(navPointId)
	if err != nil {
		return nil, fmt.Errorf("invalid nav point id: %s", navPointId)
	}

	log.Printf("Searching for nav_point with id: %d", id)

	query := `SELECT * FROM nav_points WHERE id = $1`

	if err := r.db.SelectContext(ctx, &navPoints, query, id); err != nil {
		return nil, fmt.Errorf("failed to get nav_point by id %d; error: %v", id, err)
	}

	log.Printf("Found %d nav_points", len(navPoints))
	return navPoints, nil
}

func (r *NavPointRepository) CheckSequenceStatus(ctx context.Context) error {
	var currentValue int
	query := `SELECT last_value FROM nav_points_id_seq`

	if err := r.db.QueryRowContext(ctx, query).Scan(&currentValue); err != nil {
		return fmt.Errorf("failed to check sequence status: %v", err)
	}

	log.Printf("Current sequence value: %d", currentValue)
	return nil
}

// GetAllNavPoints возвращает все навигационные точки
func (r *NavPointRepository) GetAllNavPoints(ctx context.Context) ([]model.NavPoint, error) {
	var navPoints []model.NavPoint

	query := `SELECT * FROM nav_points ORDER BY id`

	if err := r.db.SelectContext(ctx, &navPoints, query); err != nil {
		return nil, fmt.Errorf("failed to get all nav_points: %v", err)
	}

	log.Printf("Found %d nav_points", len(navPoints))
	return navPoints, nil
}

// GetAllPoints возвращает все точки (аналог point-service)
func (r *NavPointRepository) GetAllPoints(ctx context.Context) ([]model.Point, error) {
	var points []model.Point

	query := `SELECT * FROM points ORDER BY id`

	if err := r.db.SelectContext(ctx, &points, query); err != nil {
		return nil, fmt.Errorf("failed to get all points: %v", err)
	}

	log.Printf("Found %d points", len(points))
	return points, nil
}
