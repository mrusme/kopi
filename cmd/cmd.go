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
	importCmd "github.com/mrusme/kopi/import/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "kopi",
	Short: "Your command-line coffee journal",
	Long: "Kopi is a command-line coffee journal that lets you to track coffee" +
		" beans, equipment usage, brewing methods, and individual cups.\n\n" +
		"The tracked data offers insights into bean and roast preferences, caffeine" +
		" and	dairy consumption, and equipment usage patterns. This helps users" +
		" refine their coffee choices while managing caffeine intake for a more" +
		" enjoyable and responsible experience.",
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func Execute(version string) {
	rootCmd.Version = version
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
	rootCmd.AddCommand(importCmd.Cmd)
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
		viper.SetDefault("Debug", false)
		viper.SetDefault("Database", dbfile)

		viper.SetDefault("TUI.Accessible", false)

		viper.SetDefault("LLM.Ollama.Enabled", false)
		viper.SetDefault("LLM.Ollama.Host", "http://localhost:11434")

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
			formWelcome("", "", viper.GetBool("TUI.Accessible"))
		}
	}
}
