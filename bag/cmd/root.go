package cmd

import (
	"fmt"

	"github.com/mrusme/kopi/bag/cmd/open"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "bag",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bag called")
	},
}

func init() {
	Cmd.AddCommand(open.Cmd)

}
