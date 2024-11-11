package cmd

import (
	"fmt"
	"github.com/eldius/bubble-pocs/internal/config"
	cfg "github.com/eldius/initial-config-go/configs"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bubble-pocs",
	Short: "A simple POC to learn how to use Bubble Tea",
	Long:  `A simple POC to learn how to use Bubble Tea.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return cfg.InitSetup("tui-pocs",
			cfg.WithDefaultCfgFileName("config"),
			cfg.WithConfigFileToBeUsed(cfgFile),
			cfg.WithDefaultCfgFileLocations("~/.tui-client", "."),
			cfg.WithDefaultValues(map[string]any{
				cfg.LogLevelKey:      cfg.LogLevelDEBUG,
				cfg.LogOutputFileKey: "execution.log",
			}))
	},
}

var (
	cfgFile string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tui-client/config.yaml)")

	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug log")
	if err := viper.BindPFlag(config.DebugEnabled, rootCmd.PersistentFlags().Lookup("debug")); err != nil {
		err = fmt.Errorf("binding debug parameter: %w", err)
		panic(err)
	}
}
