package bagOpenCmd

import (
	"context"
	"time"

	"github.com/mrusme/kopi/bag"
	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/developer"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	globCoffeeEntity coffee.Coffee = coffee.Coffee{ID: -1}
	globBagEntity    bag.Bag       = bag.Bag{ID: -1, CoffeeID: -1}
	roastDate        string
	purchaseDate     string
	openDate         string
	price            string
)

var Cmd = &cobra.Command{
	Use:   "open",
	Short: "Open a new bag of coffee",
	Long: "Open a new bag of coffee for consumption, by entering all" +
		" necessary details, adding the specific coffee to the database. If the" +
		" coffee already exists, details can be applied from the pre-existing" +
		" one using the --coffee-id flag.",
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
			_, err := developer.InjectDummyCoffee(db)
			if err != nil {
				out.Die("%s", err)
			}
		}

		accessible := viper.GetBool("TUI.Accessible")

		coffeeDAO := coffee.NewDAO(db)
		FormCoffee(
			coffeeDAO,
			&globCoffeeEntity,
			&globBagEntity,
			"Add Coffee",
			"This wizard will guide through the steps to add a new"+
				" coffee to the database.",
			accessible,
		)

		bagDAO := bag.NewDAO(db)
		FormBag(
			bagDAO,
			&globBagEntity,
			"Open Bag",
			"This wizard will guide through the steps to open a new"+
				" bag of coffee in the database.",
			accessible,
		)

		// Add coffee to database
		if globCoffeeEntity.ID == -1 {
			globCoffeeEntity, err = coffeeDAO.Create(context.Background(), globCoffeeEntity)
			if err != nil {
				out.Die("%s", err)
			} else {
				out.Put("Coffee added to database!")
			}
		}

		// Adjust bag with missing info
		globBagEntity.OpenDate = time.Now()
		globBagEntity.CoffeeID = globCoffeeEntity.ID

		// Add bag to database
		globBagEntity, err = bagDAO.Create(context.Background(), globBagEntity)
		if err != nil {
			out.Die("%s", err)
		} else {
			out.Put("Bag opened! You can now consume coffee from this bag" +
				" using the `cup drink` command.")
		}
	},
}

func init() {
	Cmd.Flags().Int64Var(
		&globBagEntity.CoffeeID,
		"coffee-id",
		-1,
		"ID of existing coffee in the database",
	)
	Cmd.Flags().StringVar(
		&globCoffeeEntity.Roaster,
		"roaster",
		"",
		"Name of the coffee roaster",
	)
	Cmd.Flags().StringVar(
		&globCoffeeEntity.Name,
		"name",
		"",
		"Name of the coffee",
	)
	Cmd.Flags().StringVar(
		&globCoffeeEntity.Origin,
		"origin",
		"",
		"Origin of the coffee, e.g. \"Djimmah, Ethiopia\"",
	)
	Cmd.Flags().Uint16Var(
		&globCoffeeEntity.AltitudeLowerM,
		"masl-low",
		0,
		"Lower meters above sea level (masl), e.g. 1700",
	)
	Cmd.Flags().Uint16Var(
		&globCoffeeEntity.AltitudeUpperM,
		"masl-up",
		0,
		"Upper meters above sea level (masl), e.g. 2200",
	)
	Cmd.Flags().StringVar(
		&globCoffeeEntity.Level,
		"level",
		"",
		"Roasting level of the coffee, e.g. \"medium\"",
	)
	Cmd.Flags().StringVar(
		&globCoffeeEntity.Flavors,
		"flavors",
		"",
		"Cupping notes/flavors of the coffee, e.g. \"Green Tangerine, Maplesyrup\"",
	)
	Cmd.Flags().StringVar(
		&globCoffeeEntity.Info,
		"info",
		"",
		"Additional info on the coffee, e.g. \"Mountain water process washed\"",
	)
	Cmd.Flags().BoolVar(
		&globCoffeeEntity.Decaf,
		"decaf",
		false,
		"Decaf coffee",
	)

	Cmd.Flags().Int64Var(
		&globBagEntity.WeightG,
		"weight",
		0,
		"Bag weight in grams, e.g. 450",
	)
	Cmd.Flags().StringVar(
		&globBagEntity.Grind,
		"grind",
		"",
		"Grind"+
			" (possible \"beans\", \"filter\", \"frenchpress\", \"stovetop\","+
			" \"espresso\")",
	)
	Cmd.Flags().StringVar(
		&roastDate,
		"roast-date",
		"",
		"Date of roast, e.g. 2024-01-01",
	)
	Cmd.Flags().StringVar(
		&purchaseDate,
		"purchase-date",
		"",
		"Date of purchase, e.g. 2024-01-01",
	)
	now := time.Now()
	Cmd.Flags().StringVar(
		&openDate,
		"open-date",
		now.Format("2006-01-02"),
		"Date of opening the bag, e.g. 2024-01-01",
	)
	Cmd.Flags().StringVar(
		&price,
		"price",
		"",
		"Price of bag, including ISO 4217 currencty code, e.g. \"14.50 USD\"",
	)
}
