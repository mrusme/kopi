package cup

import (
	"strings"
	"time"
)

type Cup struct {
	ID    int64
	BagID int64 `validate:"required"`

	Method string `validate:"required"`
	Drink  string `validate:"required"`

	EquipmentIDs string `validate:""`

	CoffeeG uint16 `validate:"gt=0,lte=200"`
	BrewMl  uint16 `validate:"gt=0,lte=1000,ltefield=WaterMl"`
	WaterMl uint16 `validate:"gt=0,lte=1000,gtefield=BrewMl"`
	MilkMl  uint16 `validate:"gte=0,lte=1000"`
	SugarG  uint16 `validate:"gte=0,lte=100"`
	Vegan   bool   `validate:""`

	Rating    int8 `validate:"gte=0,lte=5"`
	Timestamp time.Time
}

func Table() string {
	return "`cups`"
}

var columns = []string{
	"`id`",
	"`bag_id`",

	"`method`",
	"`drink`",

	"`equipment_ids`",

	"`coffee_g`",
	"`brew_ml`",
	"`water_ml`",
	"`milk_ml`",
	"`sugar_g`",
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
		&entity.BagID,

		&entity.Method,
		&entity.Drink,

		&entity.EquipmentIDs,

		&entity.CoffeeG,
		&entity.BrewMl,
		&entity.WaterMl,
		&entity.MilkMl,
		&entity.SugarG,
		&entity.Vegan,

		&entity.Rating,
		&entity.Timestamp,
	}
}
