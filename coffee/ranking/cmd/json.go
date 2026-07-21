package coffeeRankingCmd

import (
	"encoding/json"

	"xn--gckvb8fzb.com/kopi/helpers/out"
)

func jsonOutput(rankedCoffees *[]RankedCoffee) {
	data, err := json.Marshal(rankedCoffees)
	out.NilOrDie(err)

	out.Put("%s", data)
}
