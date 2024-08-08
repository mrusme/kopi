package open

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

var cfe coffee.Coffee = coffee.Coffee{ID: -1}
var bg bag.Bag = bag.Bag{ID: -1, CoffeeID: -1}
var roastDate string
var purchaseDate string
var openDate string
var price string

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
		formCoffee(coffeeDAO, accessible)

		bagDAO := bag.NewDAO(db)
		formBag(bagDAO, accessible)

		// Add coffee to database
		if cfe.ID == -1 {
			cfe, err = coffeeDAO.Create(context.Background(), cfe)
			if err != nil {
				out.Die("%s", err)
			} else {
				out.Put("Coffee added to database!")
			}
		}

		// Adjust bag with missing info
		bg.CoffeeID = cfe.ID
		bg.OpenDate = time.Now()

		// Add bag to database
		bg, err = bagDAO.Create(context.Background(), bg)
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
		&bg.CoffeeID,
		"coffee-id",
		0,
		"ID of existing coffee in the database",
	)
	Cmd.Flags().StringVar(
		&cfe.Roaster,
		"roaster",
		"",
		"Name of the coffee roaster",
	)
	Cmd.Flags().StringVar(
		&cfe.Name,
		"name",
		"",
		"Name of the coffee",
	)
	Cmd.Flags().StringVar(
		&cfe.Origin,
		"origin",
		"",
		"Origin of the coffee, e.g. \"Djimmah, Ethiopia\"",
	)
	Cmd.Flags().Uint16Var(
		&cfe.AltitudeLowerM,
		"masl-low",
		0,
		"Lower meters above sea level (masl), e.g. 1700",
	)
	Cmd.Flags().Uint16Var(
		&cfe.AltitudeUpperM,
		"masl-up",
		0,
		"Upper meters above sea level (masl), e.g. 2200",
	)
	Cmd.Flags().StringVar(
		&cfe.Level,
		"level",
		"",
		"Roasting level of the coffee, e.g. \"medium\"",
	)
	Cmd.Flags().StringVar(
		&cfe.Flavors,
		"flavors",
		"",
		"Cupping notes/flavors of the coffee, e.g. \"Green Tangerine, Maplesyrup\"",
	)
	Cmd.Flags().StringVar(
		&cfe.Info,
		"info",
		"",
		"Additional info on the coffee, e.g. \"Mountain water process washed\"",
	)
	Cmd.Flags().BoolVar(
		&cfe.Decaf,
		"decaf",
		false,
		"Decaf coffee",
	)

	Cmd.Flags().Int64Var(
		&bg.WeightG,
		"weight",
		0,
		"Bag weight in grams, e.g. 450",
	)
	Cmd.Flags().StringVar(
		&bg.Grind,
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
