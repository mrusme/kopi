package coffee

import (
	"strings"
	"time"
)

type Coffee struct {
	ID             int64
	Roaster        string `validate:"required,max=64"`
	Name           string `validate:"required,max=64"`
	Origin         string `validate:"required,max=128"`
	AltitudeLowerM uint16 `validate:"gte=0,lte=3000,ltefield=AltitudeUpperM"`
	AltitudeUpperM uint16 `validate:"gte=0,lte=3000,gtefield=AltitudeLowerM"`
	Level          string `validate:"required,oneof=light medium dark"`
	Flavors        string `validate:"max=128"`
	Info           string `validate:"max=128"`
	Timestamp      time.Time
}

func Table() string {
	return "`coffees`"
}

var columns = []string{
	"`id`",
	"`roaster`",
	"`name`",
	"`origin`",
	"`altitude_lower_m`",
	"`altitude_upper_m`",
	"`level`",
	"`flavors`",
	"`info`",
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

func (entity *Coffee) PtrFields() []any {
	return []any{
		&entity.ID,
		&entity.Roaster,
		&entity.Name,
		&entity.Origin,
		&entity.AltitudeLowerM,
		&entity.AltitudeUpperM,
		&entity.Level,
		&entity.Flavors,
		&entity.Info,
		&entity.Timestamp,
	}
}
