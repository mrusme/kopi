package cup

import (
	"strconv"
	"strings"
	"time"
)

var MilkTypes []string = []string{
	"none",
	"regular",
	"skim",
	"lactose-free",
	"condensed",
	"a2",
	"raw",
	"goat",
	"sheep",
	"buffalo",
	"yak",
	"camel",
	"soy",
	"almond",
	"oat",
	"coconut",
	"rice",
	"cashew",
	"macadamia",
	"hemp",
	"pea",
	"flax",
	"walnut",
	"pistachio",
	"hazelnut",
	"quinoa",
	"banana",
	"barley",
}

const MilkTypesVeganStartIndex = 12

type Cup struct {
	ID    int64
	BagID int64 `validate:"required"`

	Method string `validate:"required"`
	Drink  string `validate:"required"`

	EquipmentIDs string `validate:"is_idslist"`

	CoffeeG  uint8  `validate:"gt=0,lte=200"`
	BrewMl   uint16 `validate:"gt=0,lte=1000,ltefield=WaterMl"`
	WaterMl  uint16 `validate:"gt=0,lte=1000,gtefield=BrewMl"`
	MilkMl   uint16 `validate:"gte=0,lte=1000"`
	MilkType string `validate:"required,oneof=none regular skim lactose-free condensed a2 raw goat sheep buffalo yak camel soy almond oat coconut rice cashew macadamia hemp pea flax walnut pistachio hazelnut quinoa banana barley"`
	SugarG   uint8  `validate:"gte=0,lte=100"`
	Vegan    bool   `validate:""`

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
	"`milk_type`",
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
		&entity.MilkType,
		&entity.SugarG,
		&entity.Vegan,

		&entity.Rating,
		&entity.Timestamp,
	}
}

func (entity *Cup) AddEquipmentID(id int64) {
	if len(entity.EquipmentIDs) > 0 {
		entity.EquipmentIDs += " "
	}

	entity.EquipmentIDs += strconv.FormatInt(id, 10)
}

func (entity *Cup) GetEquipmentIDs() []int64 {
	var ids []int64
	strIds := strings.Split(entity.EquipmentIDs, " ")

	for _, id := range strIds {
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, i)
	}

	return ids
}
