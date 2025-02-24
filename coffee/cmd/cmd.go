package coffeeCmd

import (
	coffeeRankingCmd "github.com/mrusme/kopi/coffee/ranking/cmd"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "coffee",
	Aliases: []string{"coffees"},
	Short:   "View coffees",
	Long: "Get information on" +
		" previously tracked coffees.",
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func init() {
	Cmd.AddCommand(coffeeRankingCmd.Cmd)
}
