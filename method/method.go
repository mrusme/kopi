package method

import (
	"strings"
)

type Method struct {
	ID          string
	Name        string
	Description string

	CaffeineMultiplierPerMl float32

	IsHot bool
}

func Table() string {
	return "`methods`"
}

var columns = []string{
	"`id`",
	"`name`",
	"`description`",

	"`caffeine_multiplier_per_ml`",

	"`is_hot`",
}

func Columns(withID bool) string {
	from := 0
	if withID == false {
		from = 1
	}

	return strings.Join(columns[from:], ",")
}

func ColumnsNumber(withID bool) int {
	if withID {
		return len(columns)
	} else {
		return len(columns) - 1
	}
}

func (entity *Method) PtrFields() []any {
	return []any{
		&entity.ID,
		&entity.Name,
		&entity.Description,

		&entity.CaffeineMultiplierPerMl,

		&entity.IsHot,
	}
}
