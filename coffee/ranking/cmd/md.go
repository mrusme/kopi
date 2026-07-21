package coffeeRankingCmd

import "xn--gckvb8fzb.com/kopi/helpers/out"

func mdOutput(rankedCoffees *[]RankedCoffee) {
	for _, rankedCoffee := range *rankedCoffees {
		out.Put("Rank #%d with an average rating of %.1f: %s by %s",
			rankedCoffee.Ranking.Ranking,
			rankedCoffee.Ranking.AvgRating,
			rankedCoffee.Coffee.Name,
			rankedCoffee.Coffee.Roaster,
		)
	}
}
