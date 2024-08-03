package equipment

import (
	"strings"
	"time"
)

type Equipment struct {
	ID int64

	Type string `validate:"required,oneof=espresso_maker coffee_maker filter grinder frother"`

	PurchaseDate     time.Time `validate:"required"`
	DecommissionDate time.Time `validate:""`

	PriceUSDct int64 `validate:"gte=0"`
	PriceSats  int64 `validate:"gte=0"`

	Timestamp time.Time
}

func Table() string {
	return "`equipments`"
}

var columns = []string{
	"`id`",

	"`type`",

	"`purchase_date`",
	"`decommission_date`",

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

func (entity *Equipment) PtrFields() []any {
	return []any{
		&entity.ID,

		&entity.Type,

		&entity.PurchaseDate,
		&entity.DecommissionDate,

		&entity.PriceUSDct,
		&entity.PriceSats,

		&entity.Timestamp,
	}
}
