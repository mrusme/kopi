package cmd

import (
	"fmt"
	"os"

	bagCmd "github.com/mrusme/kopi/bag/cmd"
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
		fmt.Printf("Database: %s\n", viper.GetString("Database"))
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

	rootCmd.AddCommand(bagCmd.Cmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cfgdir, err := os.UserConfigDir()
		cobra.CheckErr(err)

		viper.SetConfigName("kopi")
		viper.SetConfigType("toml")
		viper.AddConfigPath(cfgdir)
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		// TODO: Run welcome wizard
	}
}
