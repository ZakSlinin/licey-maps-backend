package model

type Point struct {
	ID       int      `json:"id" db:"id"`
	Name     string   `json:"name" db:"name"`
	Env      []string `json:"env" db:"env"`
	NavPoint int      `json:"nav_point" db:"nav_point"`
}
