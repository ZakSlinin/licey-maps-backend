package repository

import (
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type NavPointRepository struct{ db *sqlx.DB }

func NewNavPointRepository(db *sqlx.DB) *NavPointRepository { return &NavPointRepository{db: db} }

func (r *NavPointRepository) CreateNavPoint(navPoint *model.NavPoint) (string, error) {
	query := `INSERT INTO navPoint (id, orientation, room, type, floor) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	_, err := r.db.Exec(query, navPoint.ID, navPoint.Orientation, navPoint.Room, navPoint.Type, navPoint.Floor)
	if err != nil {
		return "", err
	}

	return navPoint.Room, nil
}

func (r *NavPointRepository) GetNavPointByID(id string) (*model.NavPoint, error) {

}
