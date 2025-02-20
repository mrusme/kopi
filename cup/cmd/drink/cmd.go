package cupDrinkCmd

import (
	"context"

	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/developer"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var globCupEntity cup.Cup = cup.Cup{BagID: -1}

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
			_, err := developer.InjectDummyEquipment(db)
			if err != nil {
				out.Die("%s", err)
			}
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
		FormCup(
			cupDAO,
			&globCupEntity,
			"Drink Cup",
			"This wizard will guide through the steps to track a new cup"+
				" of coffee that was consumed.",
			accessible,
		)

		// Add cup to database
		globCupEntity, err = cupDAO.Create(context.Background(), globCupEntity)
		if err != nil {
			out.Die("%s", err)
		} else {
			out.Put("Cup logged!")
		}
	},
}

func init() {
	Cmd.Flags().Int64Var(
		&globCupEntity.BagID,
		"bag-id",
		-1,
		"ID of existing bag in the database",
	)
	Cmd.Flags().StringVar(
		&globCupEntity.Method,
		"method",
		"",
		"Method used to prepare the cup of coffee",
	)
	Cmd.Flags().StringVar(
		&globCupEntity.Drink,
		"drink",
		"",
		"Type of drink",
	)
	Cmd.Flags().StringVar(
		&globCupEntity.EquipmentIDs,
		"equipment-ids",
		"",
		"Space-separated list of equipment IDs used for preparation,"+
			" e.g. \"1234 2345\"",
	)
	Cmd.Flags().Uint8Var(
		&globCupEntity.CoffeeG,
		"coffee-g",
		0,
		"Amount of coffee used, in grams",
	)
	Cmd.Flags().Uint16Var(
		&globCupEntity.BrewMl,
		"brew-ml",
		0,
		"Amount of brewed coffee used, in milliliters",
	)
	Cmd.Flags().Uint16Var(
		&globCupEntity.WaterMl,
		"water-ml",
		0,
		"Amount of water used, in milliliters",
	)
	Cmd.Flags().Uint16Var(
		&globCupEntity.MilkMl,
		"milk-ml",
		0,
		"Amount of milk used, in milliliters",
	)
	Cmd.Flags().Uint8Var(
		&globCupEntity.SugarG,
		"sugar-g",
		0,
		"Amount of sugar used, in grams",
	)
	Cmd.Flags().BoolVar(
		&globCupEntity.Vegan,
		"vegan",
		false,
		"Vegan drink",
	)
	Cmd.Flags().Int8Var(
		&globCupEntity.Rating,
		"rating",
		0,
		"Rating from 0 to 5",
	)
}
