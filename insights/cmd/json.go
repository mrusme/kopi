package insightsCmd

import (
	"encoding/json"

	"xn--gckvb8fzb.com/kopi/helpers/out"
)

func jsonOutput(insightsEntity *Insights) {
	data, err := json.Marshal(insightsEntity)
	out.NilOrDie(err)

	out.Put("%s", data)
}
