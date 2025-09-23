package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Step struct {
	NavID int
	Path  []int
}

// StringArray представляет массив строк для PostgreSQL
type StringArray []string

// Scan реализует интерфейс sql.Scanner
func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = StringArray{}
		return nil
	}

	switch v := value.(type) {
	case string:
		// PostgreSQL возвращает массив в формате "{val1,val2,val3}"
		if v == "{}" {
			*a = StringArray{}
			return nil
		}

		// Убираем фигурные скобки
		v = strings.Trim(v, "{}")
		if v == "" {
			*a = StringArray{}
			return nil
		}

		// Разделяем по запятым
		parts := strings.Split(v, ",")
		result := make([]string, len(parts))

		for i, part := range parts {
			part = strings.TrimSpace(part)
			result[i] = part
		}

		*a = StringArray(result)
		return nil
	default:
		return fmt.Errorf("cannot scan %v into StringArray", value)
	}
}

// Value реализует интерфейс driver.Valuer
func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}

	str := "{"
	for i, v := range a {
		if i > 0 {
			str += ","
		}
		str += v
	}
	str += "}"

	return str, nil
}

// Point представляет точку в системе (для совместимости с point-service)
type Point struct {
	ID       int         `json:"id" db:"id"`
	Name     string      `json:"name" db:"name"`
	Env      StringArray `json:"env" db:"env"`
	NavPoint int         `json:"nav_point" db:"nav_point"`
}
