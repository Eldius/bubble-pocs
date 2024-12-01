package cmd

import (
	"context"
	"github.com/eldius/bubble-pocs/internal/ui/console"

	"github.com/spf13/cobra"
)

// consoleCmd represents the console command
var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "A simple POC of a console using RCON protocol",
	Long:  `A simple POC of a console using RCON protocol.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		if err := console.Start(ctx, consoleOpts.host, consoleOpts.port, consoleOpts.password); err != nil {
			panic(err)
		}
	},
}

var (
	consoleOpts struct {
		host     string
		port     int
		password string
	}
)

func init() {
	rootCmd.AddCommand(consoleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consoleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consoleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	consoleCmd.Flags().StringVar(&consoleOpts.host, "host", "localhost", "Hostname to connect to")
	consoleCmd.Flags().IntVar(&consoleOpts.port, "port", -1, "Port to connect to")
	consoleCmd.Flags().StringVar(&consoleOpts.password, "password", "", "Password to use")
}
