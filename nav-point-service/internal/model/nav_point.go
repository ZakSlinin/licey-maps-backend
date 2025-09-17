package model

type NavPoint struct {
	ID          int    `json:"id" db:"id"`
	NavPointID  string `json:"nav_point_id" db:"nav_point_id"`
	Orientation int    `json:"orientation" db:"orientation"`
	Room        string `json:"room" db:"room"`
	Type        string `json:"type" db:"type"`
	Floor       int    `json:"floor" db:"floor"`
}
