package equipmentListCmd

import (
	"encoding/json"

	"xn--gckvb8fzb.com/kopi/equipment"
	"xn--gckvb8fzb.com/kopi/helpers/out"
)

func jsonList(entities []equipment.Equipment) {
	data, err := json.Marshal(entities)
	out.NilOrDie(err)

	out.Put("%s", data)
}
