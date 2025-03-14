package drink

import (
	"strings"
)

type Drink struct {
	ID          string
	Name        string
	Description string

	CoffeeG uint8
	BrewMl  uint16
	WaterMl uint16
	MilkMl  uint16
	SugarG  uint8

	IsHot         bool
	IsAlwaysVegan bool // e.g. pure espresso, brew coffee, ...
	CanBeVegan    bool // e.g. cappuccino, flat white, ...

	CompatibleMethods   string
	CompatibleEquipment string // space-separated, e.g. "grinder espresso_maker"
}

func Table() string {
	return "`drinks`"
}

var columns = []string{
	"`id`",
	"`name`",
	"`description`",

	"`coffee_g`",
	"`brew_ml`",
	"`water_ml`",
	"`milk_ml`",
	"`sugar_g`",

	"`is_hot`",
	"`is_always_vegan`",
	"`can_be_vegan`",

	"`compatible_methods`",
	"`compatible_equipment`",
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

		&entity.CoffeeG,
		&entity.BrewMl,
		&entity.WaterMl,
		&entity.MilkMl,
		&entity.SugarG,

		&entity.IsHot,
		&entity.IsAlwaysVegan,
		&entity.CanBeVegan,

		&entity.CompatibleMethods,
		&entity.CompatibleEquipment,
	}
}
