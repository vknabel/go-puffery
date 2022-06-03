package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vknabel/go-puffery/nav"
)

type loginModel struct {
	emailTextInput        textinput.Model
	confirmationTextInput textinput.Model
}

func initialLoginModel() loginModel {
	email := textinput.New()
	email.Focus()
	email.Placeholder = "example@puffery.app"
	email.Width = 40

	confirmation := textinput.New()
	confirmation.Placeholder = "00000000-0000-0000-0000-000000000000"
	confirmation.Width = 40

	return loginModel{email, confirmation}
}

func (m loginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m loginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab:
			if m.emailTextInput.Focused() {
				m.emailTextInput.Blur()
				m.confirmationTextInput.Focus()
				return m, textinput.Blink
			} else {
				m.confirmationTextInput.Blur()
				m.emailTextInput.Focus()
				return m, textinput.Blink
			}
		case tea.KeyEnter:
			if m.emailTextInput.Focused() {
				m.emailTextInput.Blur()
				return m, func() tea.Msg {
					_, err := Api.Login(m.emailTextInput.Value())
					if err != nil {
						return nav.PagePushMsg{Page: initialErrorModel(err)}
					}
					return m.confirmationTextInput.Focus()
				}
			} else {
				m.confirmationTextInput.Blur()
				return m, func() tea.Msg {
					_, err := Api.ConfirmLogin(m.confirmationTextInput.Value())
					if err != nil {
						return nav.PagePushMsg{Page: initialErrorModel(err)}
					}
					return nav.PageReplaceMsg{Page: initialChannelListModel()}
				}
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.emailTextInput.Width = msg.Width - 2
	}
	cmds := []tea.Cmd{}
	var cmd tea.Cmd
	m.emailTextInput, cmd = m.emailTextInput.Update(msg)
	cmds = append(cmds, cmd)
	m.confirmationTextInput, cmd = m.confirmationTextInput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m loginModel) View() string {
	return fmt.Sprintf(
		"%s\n\nWhat's your puffery email?\n\n%s\n\n%s\n\n%s",
		titleStyle.Render("Login"),
		m.emailTextInput.View(),
		m.confirmationTextInput.View(),
		"(press esc to quit)",
	)
}
