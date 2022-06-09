package tui

import (
	"github.com/charmbracelet/bubbles/list"
)

func NewStyledDefaultListDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle.
		BorderForeground(colorSecondaryBlue).
		Foreground(colorPrimaryBlue)
	delegate.Styles.SelectedDesc.
		BorderForeground(colorSecondaryBlue).
		Foreground(colorSecondaryBlue)
	return delegate
}
