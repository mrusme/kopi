package bagListCmd

import (
	"context"

	bagLabel "github.com/mrusme/kopi/bag/label"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/developer"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	filterAll    bool = false
	filterFields []string
	outputJSON   bool = false
	outputMD     bool = false
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List available bags of coffee",
	Long: "List all available bags of coffee, optionally filtered by" +
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
			coffees, err := developer.InjectDummyCoffee(db)
			if err != nil {
				out.Die("%s", err)
			}
			_, err = developer.InjectDummyBags(db, coffees)
			if err != nil {
				out.Die("%s", err)
			}
		}

		bagLabelDAO := bagLabel.NewDAO(db)
		labels, err := bagLabelDAO.List(context.Background(), !filterAll)
		out.NilOrDie(err)

		if outputJSON {
			jsonList(labels)
		} else if outputMD {
			mdList(labels, filterFields)
		} else {
			tuiList(labels, filterFields)
		}
	},
}

func init() {
	Cmd.Flags().BoolVar(
		&filterAll,
		"all",
		false,
		"List all bags, instead of only open ones",
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
