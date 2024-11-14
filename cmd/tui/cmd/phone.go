package cmd

import (
	"github.com/eldius/bubble-pocs/internal/ui/phone"

	"github.com/spf13/cobra"
)

// phoneCmd represents the phone command
var phoneCmd = &cobra.Command{
	Use:   "phone",
	Short: "A phone book",
	Long:  `A phone book.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := phone.Start(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(phoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// phoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// phoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
