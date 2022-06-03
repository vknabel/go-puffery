package tui

import (
	"github.com/charmbracelet/lipgloss"
	puffery "github.com/vknabel/go-puffery"
)

var Api puffery.Api = puffery.Api{
	Root:  "https://vapor.puffery.app",
	Token: "",
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)
