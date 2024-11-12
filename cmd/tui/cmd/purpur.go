package cmd

import (
	"github.com/eldius/bubble-pocs/internal/ui/purpur"

	"github.com/spf13/cobra"
)

// purpurCmd represents the purpur command
var purpurCmd = &cobra.Command{
	Use:   "purpur",
	Short: "Select Purpur Minecraft server versions",
	Long:  `Select Purpur Minecraft server versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := purpur.Start(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(purpurCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// purpurCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// purpurCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
