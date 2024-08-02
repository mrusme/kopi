package ranking

import (
	"strings"
)

type Ranking struct {
	CoffeeID  int64
	AvgRating float64
	Ranking   int
}

func Table() string {
	return ""
}

var columns = []string{
	"`coffee_id`",
	"`avg_rating`",
	"`ranking`",
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

func (entity *Ranking) PtrFields() []any {
	return []any{
		&entity.CoffeeID,
		&entity.AvgRating,
		&entity.Ranking,
	}
}
