package cupListCmd

import (
	"context"

	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/developer"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var filterFields []string
var outputJSON bool = false
var outputMD bool = false

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List logged cups of coffee",
	Long: "List all logged cups of coffee, optionally filtered by" +
		" custom criteria",
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

		cupDAO := cup.NewDAO(db)
		cups, err := cupDAO.List(context.Background())
		out.NilOrDie(err)

		if outputJSON {
			jsonList(cups)
		} else if outputMD {
			mdList(cups, filterFields)
		} else {
			tuiList(cups, filterFields)
		}
	},
}

func init() {
	Cmd.Flags().BoolVar(
		&outputJSON,
		"json",
		false,
		"Output JSON",
	)
	Cmd.Flags().BoolVar(
		&outputMD,
		"markdown",
		false,
		"Output Markdown",
	)
	Cmd.MarkFlagsMutuallyExclusive("json", "markdown")

	Cmd.Flags().StringSliceVar(
		&filterFields,
		"fields",
		[]string{},
		"Comma-separated fields to list; Not usable with --json",
	)
	Cmd.MarkFlagsMutuallyExclusive("fields", "json")
}
