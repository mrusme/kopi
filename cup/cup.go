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

func Columns(withID bool) string {
	columns := []string{
		"`id`",
		"`coffee_id`",
		"`drink`",
		"`timestamp`",
	}

	from := 0
	if withID == false {
		from = 1
	}

	return strings.Join(columns[from:], ",")
}

func (entity *Cup) PtrFields() []any {
	return []any{&entity.ID, &entity.CoffeeID, &entity.Drink, &entity.Timestamp}
}
