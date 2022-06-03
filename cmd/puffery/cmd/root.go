package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/vknabel/go-puffery/tui"
	"github.com/zalando/go-keyring"
)

func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "puffery",
	Short: "Read and send push notifications.",
	Long: `Read and send push notifications.
More details at https://github.com/vknabel/puffery`,
	Version: "0.0.1",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if apiUrl, ok := os.LookupEnv("PUFFERY_API_URL"); ok {
			tui.Api.Root = apiUrl
		}
		token, err := keyring.Get("puffery.app", "token")
		if err == nil && token != "" {
			tui.Api.Token = token
		}

		if apiToken, ok := os.LookupEnv("PUFFERY_API_TOKEN"); ok {
			tui.Api.Token = apiToken
		}
		m := tui.InitialAppModel()
		p := tea.NewProgram(m, tea.WithAltScreen())

		if err := p.Start(); err != nil {
			fmt.Fprintln(os.Stderr, "Error running program:", err)
			os.Exit(1)
		}
	},
}
