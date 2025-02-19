package cupCmd

import (
	cupDrinkCmd "github.com/mrusme/kopi/cup/cmd/drink"
	cupListCmd "github.com/mrusme/kopi/cup/cmd/list"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "cup",
	Aliases: []string{"cups"},
	Short:   "Track and view cups of coffee",
	Long: "Track cups of coffee that you drink, and get information on"+
	" previously tracked cups.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(cupDrinkCmd.Cmd)
	Cmd.AddCommand(cupListCmd.Cmd)
}
