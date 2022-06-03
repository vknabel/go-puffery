package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type errorModel struct {
	err error
}

func initialErrorModel(err error) errorModel {
	return errorModel{err: err}
}

func (m errorModel) Init() tea.Cmd {
	return nil
}

func (m errorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m errorModel) View() string {
	return fmt.Sprintf("%s\n\n%s\n", titleStyle.Render("Error"), m.err)
}
