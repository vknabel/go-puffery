package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vknabel/go-puffery/tui"
	"github.com/zalando/go-keyring"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "logout" {
		err := keyring.Delete("puffery.app", "token")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Successfully logged out.")
		os.Exit(0)
	}
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
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
