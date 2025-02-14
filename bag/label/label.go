package label

import (
	"database/sql"
	"strings"
	"time"
)

type Label struct {
	BagID    int64
	CoffeeID int64

	Roaster        string
	Name           string
	Origin         string
	AltitudeLowerM uint16
	AltitudeUpperM uint16
	Level          string
	Flavors        string
	Info           string
	Decaf          bool

	WeightG int64
	Grind   string

	RoastDate    time.Time
	OpenDate     time.Time
	EmptyDate    sql.NullTime
	PurchaseDate time.Time

	PriceUSDct int64
	PriceSats  int64
}

func Table() string {
	return ""
}

var columns = []string{
	"`bag_id`",
	"`coffee_id`",

	"`roaster`",
	"`name`",
	"`origin`",
	"`altitude_lower_m`",
	"`altitude_upper_m`",
	"`level`",
	"`flavors`",
	"`info`",
	"`decaf`",

	"`weight_g`",
	"`grind`",

	"`roast_date`",
	"`open_date`",
	"`empty_date`",
	"`purchase_date`",

	"`price_usd_ct`",
	"`price_sats`",
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
		&entity.Origin,
		&entity.AltitudeLowerM,
		&entity.AltitudeUpperM,
		&entity.Level,
		&entity.Flavors,
		&entity.Info,
		&entity.Decaf,

		&entity.WeightG,
		&entity.Grind,

		&entity.RoastDate,
		&entity.OpenDate,
		&entity.EmptyDate,
		&entity.PurchaseDate,

		&entity.PriceUSDct,
		&entity.PriceSats,
	}
}
