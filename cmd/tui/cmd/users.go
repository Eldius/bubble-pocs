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
	Short: "Lists users' info from Mojang API",
	Long:  `Lists users' info from Mojang API.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Start(args...)
	},
}

func init() {
	rootCmd.AddCommand(usersCmd)
}
