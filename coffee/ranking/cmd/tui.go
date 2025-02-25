package coffeeRankingCmd

import "github.com/mrusme/kopi/helpers/out"

func tuiOutput(rankedCoffees *[]RankedCoffee) {
	for _, rankedCoffee := range *rankedCoffees {
		out.Put("Rank #%d with an average rating of %.1f: %s by %s",
			rankedCoffee.Ranking.Ranking,
			rankedCoffee.Ranking.AvgRating,
			rankedCoffee.Coffee.Name,
			rankedCoffee.Coffee.Roaster,
		)
	}
}
