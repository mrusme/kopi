package bagListCmd

import (
	"encoding/json"

	bagLabel "github.com/mrusme/kopi/bag/label"
	"github.com/mrusme/kopi/helpers/out"
)

func jsonList(entities []bagLabel.Label) {
	data, err := json.Marshal(entities)
	out.NilOrDie(err)

	out.Put("%s", data)
}
