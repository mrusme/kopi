package equipmentCmd

import (
	"fmt"

	equipmentAddCmd "github.com/mrusme/kopi/equipment/cmd/add"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "equipment",
	Aliases: []string{},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("equipment called")
	},
}

func init() {
	Cmd.AddCommand(equipmentAddCmd.Cmd)
}
