package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	puffery "github.com/vknabel/go-puffery"
)

var Api puffery.Api = puffery.New()

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var titleStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("62")).
	Foreground(lipgloss.Color("230")).
	Padding(0, 1)

var backKeyBinding = key.NewBinding(
	key.WithKeys("esc"),
	key.WithHelp("esc", "back"),
)
