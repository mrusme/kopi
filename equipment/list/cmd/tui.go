package equipmentListCmd

import (
	"github.com/mrusme/kopi/equipment"
	"github.com/mrusme/kopi/helpers/formatter"
	"github.com/mrusme/kopi/helpers/out"
)

func tuiList(entities []equipment.Equipment, fields []string) {
	out.Put(formatter.ListToTUI(entities, fields))
}
