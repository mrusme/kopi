package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/adrg/xdg"
	bagCmd "github.com/mrusme/kopi/bag/cmd"
	cupCmd "github.com/mrusme/kopi/cup/cmd"
	equipmentCmd "github.com/mrusme/kopi/equipment/cmd"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "kopi",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		formWelcome("", "", false) // TODO: Move to initConfig
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	cfgdir, _ := os.UserConfigDir()
	rootCmd.PersistentFlags().StringVarP(
		&cfgFile,
		"config",
		"c",
		"",
		fmt.Sprintf("config file (default \"%s/kopi.toml\")", cfgdir),
	)

	rootCmd.AddCommand(equipmentCmd.Cmd)
	rootCmd.AddCommand(bagCmd.Cmd)
	rootCmd.AddCommand(cupCmd.Cmd)
}

func initConfig() {
	var dbfile string

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cfgdir, err := os.UserConfigDir()
		out.NilOrDie(err)

		dbfile, err = xdg.DataFile("kopi.sqlite3")
		out.NilOrDie(err)

		viper.SetDefault("Developer", false)
		viper.SetDefault("Database", dbfile)
		viper.SetDefault("TUI.Accessible", false)

		viper.SetConfigName("kopi")
		viper.SetConfigType("toml")
		viper.AddConfigPath(cfgdir)
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &(viper.ConfigParseError{})) ||
			errors.As(err, &(viper.ConfigMarshalError{})) {
			out.Die("Please double-check your configuration:\n%s", err)
		} else if errors.As(err, &(viper.ConfigFileNotFoundError{})) {
			err := viper.SafeWriteConfigAs(dbfile)
			out.NilOrDie(err)
			// TODO: formWelcome()
		}
	}
}
