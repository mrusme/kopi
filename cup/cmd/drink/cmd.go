package cupDrinkCmd

import (
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/developer"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cp cup.Cup = cup.Cup{BagID: -1}

var Cmd = &cobra.Command{
	Use:   "drink",
	Short: "Drink a cup of coffee",
	Long: "Add a cup of coffee that you consume(d) to your log, including all" +
		" the relevant details and a rating.",
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
			coffees, err := developer.InjectDummyCoffee(db)
			if err != nil {
				out.Die("%s", err)
			}
			_, err = developer.InjectDummyBags(db, coffees)
			if err != nil {
				out.Die("%s", err)
			}
		}

		accessible := viper.GetBool("TUI.Accessible")

		cupDAO := cup.NewDAO(db)
		formCup(cupDAO, accessible)

	},
}

func init() {
	Cmd.Flags().Int64Var(
		&cp.BagID,
		"bag-id",
		-1,
		"ID of existing bag in the database",
	)
	Cmd.Flags().StringVar(
		&cp.Method,
		"method",
		"",
		"Method used to prepare the cup of coffee",
	)
	Cmd.Flags().StringVar(
		&cp.Drink,
		"drink",
		"",
		"Type of drink",
	)
	Cmd.Flags().StringVar(
		&cp.EquipmentIDs,
		"equipment-ids",
		"",
		"Space-separated list of equipment IDs used for preparation,"+
			" e.g. \"1234 2345\"",
	)
	Cmd.Flags().Uint8Var(
		&cp.CoffeeG,
		"coffee-g",
		0,
		"Amount of coffee used, in grams",
	)
	Cmd.Flags().Uint16Var(
		&cp.BrewMl,
		"brew-ml",
		0,
		"Amount of brewed coffee used, in milliliters",
	)
	Cmd.Flags().Uint16Var(
		&cp.WaterMl,
		"water-ml",
		0,
		"Amount of water used, in milliliters",
	)
	Cmd.Flags().Uint16Var(
		&cp.MilkMl,
		"milk-ml",
		0,
		"Amount of milk used, in milliliters",
	)
	Cmd.Flags().Uint8Var(
		&cp.SugarG,
		"sugar-g",
		0,
		"Amount of sugar used, in grams",
	)
	Cmd.Flags().BoolVar(
		&cp.Vegan,
		"vegan",
		false,
		"Vegan drink",
	)
	Cmd.Flags().Int8Var(
		&cp.Rating,
		"rating",
		0,
		"Rating from 0 to 5",
	)
}
