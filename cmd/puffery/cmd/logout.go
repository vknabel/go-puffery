package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

func init() {
	rootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of Puffery",
	Long:  `Log out of Puffery`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := keyring.Delete("puffery.app", "token")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Successfully logged out.")
		os.Exit(0)
	},
}
