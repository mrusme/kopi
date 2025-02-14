package equipmentListCmd

import (
	"context"

	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/developer"
	"github.com/mrusme/kopi/equipment"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var filterAll bool = false
var filterFields []string
var outputJSON bool = false
var outputMD bool = false

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List available equipments of coffee",
	Long: "List all available equipments of coffee, optionally filtered by" +
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

		equipmentDAO := equipment.NewDAO(db)
		entities, err := equipmentDAO.List(context.Background(), !filterAll)
		out.NilOrDie(err)

		if outputJSON {
			jsonList(entities)
		} else if outputMD {
			mdList(entities, filterFields)
		} else {
			tuiList(entities, filterFields)
		}
	},
}

func init() {
	Cmd.Flags().BoolVar(
		&filterAll,
		"all",
		false,
		"List all equipment, instead of only non-decommissioned ones",
	)

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
