package bagListCmd

import (
	"encoding/json"

	bagLabel "xn--gckvb8fzb.com/kopi/bag/label"
	"xn--gckvb8fzb.com/kopi/helpers/out"
)

func jsonList(entities []bagLabel.Label) {
	data, err := json.Marshal(entities)
	out.NilOrDie(err)

	out.Put("%s", data)
}
