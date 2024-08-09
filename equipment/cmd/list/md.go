package equipmentListCmd

import (
	"github.com/mrusme/kopi/equipment"
	"github.com/mrusme/kopi/helpers/formatter"
	"github.com/mrusme/kopi/helpers/out"
)

func mdList(entities []equipment.Equipment, fields []string) {
	out.Put(formatter.ListToMarkdown(entities, fields))
}
