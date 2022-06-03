package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vknabel/go-puffery/nav"
)

type appModel struct {
	stack nav.NavStack

	err error
}

func InitialAppModel() appModel {
	var initial tea.Model
	if Api.Token == "" {
		initial = initialLoginModel()
	} else {
		initial = initialListModel()
	}
	return appModel{
		stack: nav.NewStack(initial),
	}
}

func (m appModel) Init() tea.Cmd {
	return m.stack.Init()
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.stack, cmd = m.stack.Update(msg)
	return m, cmd
}

func (m appModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	return m.stack.View()
}
