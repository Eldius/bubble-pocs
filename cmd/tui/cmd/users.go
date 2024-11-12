package cmd

import (
	"github.com/eldius/bubble-pocs/internal/ui/mojang"
	"github.com/eldius/bubble-pocs/internal/ui/mojang/styled"
	"github.com/spf13/cobra"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Show Minecraft users' info from Mojang API by name",
	Long:  `Show Minecraft users' info from Mojang API by name.`,
	Run: func(cmd *cobra.Command, args []string) {
		if usersOpts.styled {
			styled.Start(args...)
			return
		}
		mojang.Start(args...)
	},
}

var (
	usersOpts struct {
		styled bool
	}
)

func init() {
	rootCmd.AddCommand(usersCmd)
	usersCmd.Flags().BoolVarP(&usersOpts.styled, "styled", "s", false, "Enable styling")
}
