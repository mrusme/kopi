package equipmentListCmd

import (
	"encoding/json"

	"github.com/mrusme/kopi/equipment"
	"github.com/mrusme/kopi/helpers/out"
)

func jsonList(entities []equipment.Equipment) {
	data, err := json.Marshal(entities)
	out.NilOrDie(err)

	out.Put("%s", data)
}
