package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/vknabel/go-puffery/tui"
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
		m := tui.InitialAppModel()
		p := tea.NewProgram(m, tea.WithAltScreen())

		if err := p.Start(); err != nil {
			fmt.Fprintln(os.Stderr, "Error running program:", err)
			os.Exit(1)
		}
	},
}
