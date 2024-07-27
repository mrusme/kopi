package coffee

import (
	"strings"
	"time"
)

type Coffee struct {
	ID            int64
	Roaster       string
	Name          string
	Origin        string
	AltitudeLower int
	AltitudeUpper int
	Level         string
	Flavors       string
	Info          string
	RoastingDate  time.Time
	Timestamp     time.Time
}

func Table() string {
	return "`coffees`"
}

var columns = []string{
	"`id`",
	"`roaster`",
	"`name`",
	"`origin`",
	"`altitude_lower`",
	"`altitude_upper`",
	"`level`",
	"`flavors`",
	"`info`",
	"`roasting_date`",
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
		&entity.AltitudeLower,
		&entity.AltitudeUpper,
		&entity.Level,
		&entity.Flavors,
		&entity.Info,
		&entity.RoastingDate,
		&entity.Timestamp,
	}
}
