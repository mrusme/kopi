package equipmentCmd

import (
	equipmentAddCmd "xn--gckvb8fzb.com/kopi/equipment/add/cmd"
	equipmentListCmd "xn--gckvb8fzb.com/kopi/equipment/list/cmd"
	equipmentLogCmd "xn--gckvb8fzb.com/kopi/equipment/log/cmd"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "equipment",
	Aliases: []string{},
	Short:   "Manage coffee equipment",
	Long:    "Add new, and update and list existing coffee equipment.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(equipmentAddCmd.Cmd)
	Cmd.AddCommand(equipmentLogCmd.Cmd)
	Cmd.AddCommand(equipmentListCmd.Cmd)
}
