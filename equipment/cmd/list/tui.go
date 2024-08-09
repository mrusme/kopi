package equipmentListCmd

import (
	"github.com/mrusme/kopi/equipment"
	"github.com/mrusme/kopi/helpers/out"
)

func tuiList(entities []equipment.Equipment, accessible bool) {
	for _, entity := range entities {
		out.Put("%s", entity.Name)
	}
}
