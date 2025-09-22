package repository

import (
	"context"
	"fmt"
	"github.com/ZakSlinin/licey-maps-backend/internal/model"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type PointRepository struct{ db *sqlx.DB }

func NewPointRepository(db *sqlx.DB) *PointRepository { return &PointRepository{db: db} }

func (r *PointRepository) CreatePoint(ctx context.Context, name string, env []string, navPoint int) (int, error) {
	var returnedId int
	query := `
		INSERT INTO points (name, env, nav_point)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query, name, env, navPoint).Scan(&returnedId)
	if err != nil {
		return 0, err
	}

	log.Printf("Point created with id: %d", returnedId)
	return returnedId, nil
}

func (r *PointRepository) GetPointByID(ctx context.Context, pointId string) ([]model.Point, error) {
	var point []model.Point

	id, err := strconv.Atoi(pointId)
	if err != nil {
		return nil, fmt.Errorf("invalid point id: %s", pointId)
	}

	log.Printf("Getting point with id: %d", id)

	query := `SELECT * FROM points WHERE id = $1`

	if err := r.db.SelectContext(ctx, &point, query, id); err != nil {
		return nil, fmt.Errorf("failed to get point by id: %w", err)
	}

	log.Printf("Found %d points", len(point))
	return point, nil
}
