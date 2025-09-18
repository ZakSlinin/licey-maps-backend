package repository

import (
	"context"
	"fmt"
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type NavPointRepository struct{ db *sqlx.DB }

func NewNavPointRepository(db *sqlx.DB) *NavPointRepository { return &NavPointRepository{db: db} }

func (r *NavPointRepository) CreateNavPoint(ctx context.Context, orientation string, room, navType string, floor int) (string, error) {

	query := `
        INSERT INTO navPoint (orientation, room, type, floor)
        VALUES ($1, $2, $3, $4)
        RETURNING room
    `
	_, err := r.db.Exec(query, orientation, room, navType, floor)
	if err != nil {
		return "", err
	}

	return room, nil
}

func (r *NavPointRepository) GetNavPointByNavPointID(ctx context.Context, navPointId string) ([]model.NavPoint, error) {
	var navPoints []model.NavPoint
	query := `SELECT * FROM navPoint WHERE nav_point_id = $1`

	if err := r.db.SelectContext(ctx, &navPoints, query, navPointId); err != nil {
		return nil, fmt.Errorf("failed to get nav_point by id %s; error: %d", navPointId, err)
	}

	return navPoints, nil
}
