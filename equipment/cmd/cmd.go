package equipmentCmd

import (
	equipmentAddCmd "github.com/mrusme/kopi/equipment/add/cmd"
	equipmentListCmd "github.com/mrusme/kopi/equipment/list/cmd"
	equipmentLogCmd "github.com/mrusme/kopi/equipment/log/cmd"
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
