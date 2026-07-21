package cupListCmd

import (
	"encoding/json"

	"xn--gckvb8fzb.com/kopi/cup"
	"xn--gckvb8fzb.com/kopi/helpers/out"
)

func jsonList(entities []cup.Cup) {
	data, err := json.Marshal(entities)
	out.NilOrDie(err)

	out.Put("%s", data)
}
