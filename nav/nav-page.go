package nav

import tea "github.com/charmbracelet/bubbletea"

type NavPage struct {
	windowSize *tea.WindowSizeMsg
	previous   tea.Model
	current    tea.Model
}

func NewPage(current tea.Model) NavPage {
	return NavPage{
		windowSize: nil,
		previous:   nil,
		current:    current,
	}
}

func (m NavPage) Init() tea.Cmd {
	return tea.Batch(m.current.Init(), func() tea.Msg {
		if m.windowSize != nil {
			return *m.windowSize
		}
		return nil
	})
}

func (m NavPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize = &msg
	case PagePopMsg:
		m.current = m.previous
		if m.windowSize != nil {
			m.current, cmd = m.current.Update(*m.windowSize)
			cmds = append(cmds, cmd)
			m.current, cmd = m.current.Update(PageRestoreMsg{})
			cmds = append(cmds, cmd)
		}
	case PagePushMsg:
		new := NewPage(msg.Page)
		new.previous = m
		new.windowSize = m.windowSize
		return new, new.Init()
	}
	m.current, cmd = m.current.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m NavPage) View() string {
	return m.current.View()
}
