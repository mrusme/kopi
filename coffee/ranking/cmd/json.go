package coffeeRankingCmd

import (
	"encoding/json"

	"github.com/mrusme/kopi/helpers/out"
)

func jsonOutput(rankedCoffees *[]RankedCoffee) {
	data, err := json.Marshal(rankedCoffees)
	out.NilOrDie(err)

	out.Put("%s", data)
}
