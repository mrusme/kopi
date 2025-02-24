package coffeeRankingCmd

import (
	"context"

	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/coffee/ranking"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/developer"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Cmd = &cobra.Command{
	Use:     "ranking",
	Aliases: []string{"rank"},
	Short:   "Ranking of coffee",
	Long: "Show the ranking of coffees based on the ratings of individually" +
		" consumed cups.",
	Run: func(cmd *cobra.Command, args []string) {
		var devMode bool = viper.GetBool("Developer")

		db, err := dal.Open(
			viper.GetString("Database"),
			devMode,
		)
		if err != nil {
			out.Die("%s", err)
		}

		if devMode {
			eqpt, err := developer.InjectDummyEquipment(db)
			if err != nil {
				out.Die("%s", err)
			}
			coffees, err := developer.InjectDummyCoffee(db)
			if err != nil {
				out.Die("%s", err)
			}
			bags, err := developer.InjectDummyBags(db, coffees)
			if err != nil {
				out.Die("%s", err)
			}
			_, err = developer.InjectDummyCups(db, eqpt, coffees, bags)
			if err != nil {
				out.Die("%s", err)
			}
		}

		// accessible := viper.GetBool("TUI.Accessible")

		coffeeDAO := coffee.NewDAO(db)
		rankingDAO := ranking.NewDAO(db)

		rankedCups, err := rankingDAO.GetRanking(context.Background())
		out.NilOrDie(err)
		for _, rankedCup := range rankedCups {
			rankedCoffee, err := coffeeDAO.GetByID(context.Background(), rankedCup.CoffeeID)
			out.NilOrDie(err)
			out.Put("Rank #%d with an average rating of %f: %s",
				rankedCup.Ranking,
				rankedCup.AvgRating,
				rankedCoffee.Name,
			)
		}
	},
}

func init() {
}
