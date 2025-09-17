package repository

import (
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type NavPointRepository struct{ db *sqlx.DB }

func NewNavPointRepository(db *sqlx.DB) *NavPointRepository { return &NavPointRepository{db: db} }

func (r *NavPointRepository) CreateNavPoint(navPoint *model.NavPoint) (string, error) {
	query := `INSERT INTO navPoint (id, nav_point_id, orientation, room, type, floor) VALUES ($1, $2, $3, $4, $5) RETURNING nav_point_id`

	_, err := r.db.Exec(query, navPoint.ID, navPoint.NavPointID, navPoint.Orientation, navPoint.Room, navPoint.Type, navPoint.Floor)
	if err != nil {
		return "", err
	}

	return navPoint.Room, nil
}

func (r *NavPointRepository) GetNavPointByNavPointID(navPointId string) ([]model.NavPoint, error) {
	var navPoints []model.NavPoint
	query := `SELECT * FROM navPoint WHERE nav_point_id = $1`

	_, err := r.db.Exec(query, navPointId)
	if err != nil {
		return nil, err
	}

	return navPoints, nil
}
