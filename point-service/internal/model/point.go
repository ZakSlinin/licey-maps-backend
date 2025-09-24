package model

import "github.com/lib/pq"

type Point struct {
	ID       int            `json:"id" db:"id"`
	Name     string         `json:"name" db:"name"`
	Env      pq.StringArray `json:"env" db:"env"` // Используем встроенный тип
	NavPoint int            `json:"nav_point" db:"nav_point"`
}
