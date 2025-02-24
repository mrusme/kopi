package insightsCmd

import (
	"encoding/json"

	"github.com/mrusme/kopi/helpers/out"
)

func jsonOutput(insightsEntity *Insights) {
	data, err := json.Marshal(insightsEntity)
	out.NilOrDie(err)

	out.Put("%s", data)
}
