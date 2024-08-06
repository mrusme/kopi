package open

import (
	"fmt"
	"time"

	"github.com/mrusme/kopi/bag"
	"github.com/mrusme/kopi/coffee"
	"github.com/spf13/cobra"
)

var cfe coffee.Coffee = coffee.Coffee{}
var bg bag.Bag = bag.Bag{}
var roastDate string
var purchaseDate string
var openDate string
var price string

var Cmd = &cobra.Command{
	Use:   "open",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("open called")

	},
}

func init() {
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
