package cup

import (
	"strings"
	"time"
)

type Cup struct {
	ID        int64
	CoffeeID  int64
	Drink     string
	Timestamp time.Time
}

func Table() string {
	return "`cups`"
}

var columns = []string{
	"`id`",
	"`coffee_id`",
	"`drink`",
	"`timestamp`",
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

func (entity *Cup) PtrFields() []any {
	return []any{
		&entity.ID,
		&entity.CoffeeID,
		&entity.Drink,
		&entity.Timestamp,
	}
}
