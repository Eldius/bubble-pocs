/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/eldius/bubble-pocs/internal/mojang/ui"
	"github.com/spf13/cobra"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Start(args...)
	},
}

func init() {
	rootCmd.AddCommand(usersCmd)
}
