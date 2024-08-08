package label

import (
	"strings"
)

type Label struct {
	BagID    int64
	CoffeeID int64
	Roaster  string
	Name     string
}

func Table() string {
	return ""
}

var columns = []string{
	"`bag_id`",
	"`coffee_id`",
	"`roaster`",
	"`name`",
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

func (entity *Label) PtrFields() []any {
	return []any{
		&entity.BagID,
		&entity.CoffeeID,
		&entity.Roaster,
		&entity.Name,
	}
}
