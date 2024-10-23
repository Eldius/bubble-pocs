/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/eldius/bubble-pocs/internal/client/mojang"
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
		c := mojang.NewMojang()

		users, err := c.FetchUsers(usersOpts.users...)
		if err != nil {
			panic(err)
		}

		fmt.Println("---\nusers:")
		for _, u := range users {
			fmt.Printf("- %s: %s\n", u.Name, u.Id)
		}
	},
}

var (
	usersOpts struct {
		users []string
	}
)

func init() {
	rootCmd.AddCommand(usersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	usersCmd.Flags().StringArrayVarP(&usersOpts.users, "users", "u", []string{}, "users to search")
}
