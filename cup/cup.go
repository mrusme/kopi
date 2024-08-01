package cup

import (
	"database/sql"
	"strings"
	"time"
)

type Cup struct {
	ID              int64
	CoffeeID        int64
	Drink           string
	OverrideCoffeeG sql.NullInt16
	OverrideBrewMl  sql.NullInt16
	OverrideWaterMl sql.NullInt16
	OverrideMilkMl  sql.NullInt16
	OverrideSugarG  sql.NullInt16
	Vegan           bool
	Rating          int8
	Timestamp       time.Time
}

func Table() string {
	return "`cups`"
}

var columns = []string{
	"`id`",
	"`coffee_id`",
	"`drink`",
	"`override_coffee_g`",
	"`override_brew_ml`",
	"`override_water_ml`",
	"`override_milk_ml`",
	"`override_sugar_g`",
	"`vegan`",
	"`rating`",
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
		&entity.OverrideCoffeeG,
		&entity.OverrideBrewMl,
		&entity.OverrideWaterMl,
		&entity.OverrideMilkMl,
		&entity.OverrideSugarG,
		&entity.Vegan,
		&entity.Rating,
		&entity.Timestamp,
	}
}
