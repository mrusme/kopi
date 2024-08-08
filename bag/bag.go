package bag

import (
	"database/sql"
	"strings"
	"time"
)

type Bag struct {
	ID       int64
	CoffeeID int64 `validate:"required"`

	WeightG int64  `validate:"required,gte=0"`
	Grind   string `validate:"required,oneof=beans filter frenchpress stovetop espresso"`

	RoastDate    time.Time    `validate:"required"`
	OpenDate     time.Time    `validate:"required"`
	EmptyDate    sql.NullTime `validate:""`
	PurchaseDate time.Time    `validate:"required"`

	PriceUSDct int64 `validate:"gte=0"`
	PriceSats  int64 `validate:"gte=0"`

	Timestamp time.Time
}

func Table() string {
	return "`bags`"
}

var columns = []string{
	"`id`",
	"`coffee_id`",

	"`weight_g`",
	"`grind`",

	"`roast_date`",
	"`open_date`",
	"`empty_date`",
	"`purchase_date`",

	"`price_usd_ct`",
	"`price_sats`",

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

func (entity *Bag) PtrFields() []any {
	return []any{
		&entity.ID,
		&entity.CoffeeID,

		&entity.WeightG,
		&entity.Grind,

		&entity.RoastDate,
		&entity.OpenDate,
		&entity.EmptyDate,
		&entity.PurchaseDate,

		&entity.PriceUSDct,
		&entity.PriceSats,

		&entity.Timestamp,
	}
}
