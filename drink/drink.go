package drink

import (
	"strings"
)

type Drink struct {
	ID                      string
	Name                    string
	Description             string
	Method                  string // espresso, pourover, drip, ...
	CaffeineMultiplierPerMl float32
	RequiredCoffeeG         uint8
	RequiredBrewMl          uint16
	RequiredWaterMl         uint16
	RequiredMilkMl          uint16
	RequiredSugarG          uint8
	IsHot                   bool
	IsAlwaysVegan           bool // e.g. pure espresso, brew coffee, ...
	CanBeVegan              bool // e.g. cappuccino, flat white, ...
}

func Table() string {
	return "`drinks`"
}

var columns = []string{
	"`id`",
	"`name`",
	"`description`",
	"`method`",
	"`caffeine_multiplier_per_ml`",
	"`required_coffee_g`",
	"`required_brew_ml`",
	"`required_water_ml`",
	"`required_milk_ml`",
	"`required_sugar_g`",
	"`is_hot`",
	"`is_always_vegan`",
	"`can_be_vegan`",
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

func (entity *Drink) PtrFields() []any {
	return []any{
		&entity.ID,
		&entity.Name,
		&entity.Description,
		&entity.Method,
		&entity.CaffeineMultiplierPerMl,
		&entity.RequiredCoffeeG,
		&entity.RequiredBrewMl,
		&entity.RequiredWaterMl,
		&entity.RequiredMilkMl,
		&entity.RequiredSugarG,
		&entity.IsHot,
		&entity.IsAlwaysVegan,
		&entity.CanBeVegan,
	}
}
