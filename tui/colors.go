package tui

import "github.com/charmbracelet/lipgloss"

const (
	colorLagoonGreen      = lipgloss.Color("#6FD19A")
	colorLagoonBubbleBlue = lipgloss.Color("#6DBDE6")
	colorFinRed           = lipgloss.Color("#C06376")
	colorBodyYellow       = lipgloss.Color("#F3CE72")
	colorLagoonLightBlue  = lipgloss.Color("#8ED6EE")
	colorLagoonDeepBlue   = lipgloss.Color("#6DBDE6")

	colorPlaceholder = lipgloss.Color("240")
)

var (
	colorBackground = lipgloss.AdaptiveColor{
		Light: "#FFFFFF",
		Dark:  "#202020",
	}
	colorPrimaryBlue = lipgloss.AdaptiveColor{
		Light: string(colorLagoonDeepBlue),
		Dark:  string(colorLagoonBubbleBlue),
	}
	colorSecondaryBlue = lipgloss.AdaptiveColor{
		Light: string(colorLagoonLightBlue),
		Dark:  string(colorLagoonLightBlue),
	}
)
