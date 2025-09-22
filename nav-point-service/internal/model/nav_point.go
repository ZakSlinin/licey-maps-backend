package model

type NavPoint struct {
	ID          int    `json:"id" db:"id"`
	Orientation string `json:"orientation" db:"orientation"`
	Neighbours  []int  `json:"neighbours" db:"neighbours"`
	Room        string `json:"room" db:"room"`
	Type        string `json:"type" db:"type"`
	Floor       int    `json:"floor" db:"floor"`
}
