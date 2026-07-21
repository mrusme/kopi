package bagCmd

import (
	bagListCmd "xn--gckvb8fzb.com/kopi/bag/list/cmd"
	bagOpenCmd "xn--gckvb8fzb.com/kopi/bag/open/cmd"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "bag",
	Aliases: []string{"bags"},
	Short:   "Manage coffee bags",
	Long: "Add new, and update and list existing coffee bags.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(bagOpenCmd.Cmd)
	Cmd.AddCommand(bagListCmd.Cmd)
}
