package cupListCmd

import (
	"encoding/json"

	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/helpers/out"
)

func jsonList(entities []cup.Cup) {
	data, err := json.Marshal(entities)
	out.NilOrDie(err)

	out.Put("%s", data)
}
