package nav_test

import tea "github.com/charmbracelet/bubbletea"

type testModel struct{}

func (testModel) Init() tea.Cmd                       { return nil }
func (testModel) Update(tea.Msg) (tea.Model, tea.Cmd) { return nil, nil }
func (testModel) View() string                        { return "" }
