package open

import (
	"log"
	"time"

	"github.com/mrusme/kopi/bag"
	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/dal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfeId int64
var cfe coffee.Coffee = coffee.Coffee{}
var bg bag.Bag = bag.Bag{}
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
		db, err := dal.Open(viper.GetString("Database"))
		if err != nil {
			log.Fatalf("%s\n", err)
		}

		accessible := viper.GetBool("TUI.Accessible")

		formCoffee(db, accessible)
		formBag(db, accessible)

	},
}

func init() {
	Cmd.Flags().Int64Var(
		&cfeId,
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
		&cfe.IsDecaf,
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
