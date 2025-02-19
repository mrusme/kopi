package log

import (
	"strings"
	"time"
)

type Log struct {
	ID          int64
	EquipmentID int64

	Key   string `validate:"required,max=64"`
	Value string `validate:"required,max=1024"`

	Timestamp time.Time
}

func Table() string {
	return "`equipment_logs`"
}

var columns = []string{
	"`id`",
	"`equipment_id`",

	"`key`",
	"`value`",

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

func (entity *Log) PtrFields() []any {
	return []any{
		&entity.ID,
		&entity.EquipmentID,

		&entity.Key,
		&entity.Value,

		&entity.Timestamp,
	}
}
