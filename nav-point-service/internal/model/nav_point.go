package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// IntArray представляет массив целых чисел для PostgreSQL
type IntArray []int

// Scan реализует интерфейс sql.Scanner
func (a *IntArray) Scan(value interface{}) error {
	if value == nil {
		*a = IntArray{}
		return nil
	}

	switch v := value.(type) {
	case string:
		// PostgreSQL возвращает массив в формате "{1,2,3}"
		if v == "{}" {
			*a = IntArray{}
			return nil
		}

		// Убираем фигурные скобки
		v = strings.Trim(v, "{}")
		if v == "" {
			*a = IntArray{}
			return nil
		}

		// Разделяем по запятым
		parts := strings.Split(v, ",")
		result := make([]int, len(parts))

		for i, part := range parts {
			part = strings.TrimSpace(part)
			num, err := strconv.Atoi(part)
			if err != nil {
				return fmt.Errorf("cannot scan %v into IntArray", value)
			}
			result[i] = num
		}

		*a = IntArray(result)
		return nil
	default:
		return fmt.Errorf("cannot scan %v into IntArray", value)
	}
}

// Value реализует интерфейс driver.Valuer
func (a IntArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}

	str := "{"
	for i, v := range a {
		if i > 0 {
			str += ","
		}
		str += strconv.Itoa(v)
	}
	str += "}"

	return str, nil
}

type NavPoint struct {
	ID          int      `json:"id" db:"id"`
	Orientation string   `json:"orientation" db:"orientation"`
	Neighbours  IntArray `json:"neighbours" db:"neighbours"`
	Room        string   `json:"room" db:"room"`
	Type        string   `json:"type" db:"type"`
	Floor       int      `json:"floor" db:"floor"`
}
