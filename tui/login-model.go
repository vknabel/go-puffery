package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vknabel/go-puffery/nav"
)

const (
	loginStepHasAccount = iota
	loginStepEmail
	loginStepConfirmation
)

type loginModel struct {
	currentStep           int
	hasAccountQuestion    questionModel
	emailTextInput        textinput.Model
	confirmationTextInput textinput.Model
	spinner               spinner.Model
	isBusy                bool
}

func initialLoginModel() loginModel {
	email := textinput.New()
	email.Prompt = "What is you email address? "
	email.Placeholder = "example@puffery.app"
	email.Width = 40
	email.PromptStyle = promptStyle
	email.TextStyle = answerStyle
	email.PlaceholderStyle = placeholderStyle

	confirmation := textinput.New()
	confirmation.Prompt = "Enter the confirmation code: "
	confirmation.Placeholder = "00000000-0000-0000-0000-000000000000"
	confirmation.Width = 40
	confirmation.PromptStyle = promptStyle
	confirmation.TextStyle = answerStyle
	confirmation.PlaceholderStyle = placeholderStyle

	hasAccount := newQuestion("Do you already have an account with email? ")
	hasAccount.PromptStyle = promptStyle
	hasAccount.PlaceholderStyle = lipgloss.NewStyle().Foreground(colorPlaceholder)
	hasAccount.AnswerStyle = lipgloss.NewStyle().Foreground(colorLagoonBubbleBlue)

	activity := spinner.New()
	activity.Spinner = spinner.Dot

	return loginModel{loginStepHasAccount, hasAccount, email, confirmation, activity, false}
}

func (m loginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m loginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loginModelSuccessfulMsg:
		return initialChannelListModel(), nil
	case loginModelFailedMsg:
		m.isBusy = false
		return m, nav.Push(initialErrorModel(msg.err))
	case loginModelNeedsConfirmationMsg:
		m.isBusy = false
		m.emailTextInput.Blur()
		m.currentStep = loginStepConfirmation
		m.confirmationTextInput.Focus()
		return m, textinput.Blink
	case tea.KeyMsg:
		switch m.currentStep {
		case loginStepHasAccount:
			var cmd tea.Cmd
			m.hasAccountQuestion, cmd = m.hasAccountQuestion.Update(msg)
			if m.hasAccountQuestion.answered {
				if !m.hasAccountQuestion.answer {
					m.emailTextInput.Placeholder = "(optional)"
				}
				m.emailTextInput.Focus()
				m.currentStep = loginStepEmail
				return m, tea.Batch(cmd, textinput.Blink)
			}
		case loginStepEmail:
			if msg.Type == tea.KeyEnter {
				if m.emailTextInput.Value() == "" {
					if m.hasAccountQuestion.answer {
						return m, nil
					}
					m.isBusy = true
					return m, tea.Batch(m.spinner.Tick, m.register)
				} else if m.hasAccountQuestion.answer {
					m.isBusy = true
					return m, tea.Batch(m.spinner.Tick, m.login)
				} else {
					m.isBusy = true
					return m, tea.Batch(m.spinner.Tick, m.register)
				}
			}
			var cmd tea.Cmd
			m.emailTextInput, cmd = m.emailTextInput.Update(msg)
			return m, cmd
		case loginStepConfirmation:
			if msg.Type == tea.KeyEnter {
				m.isBusy = true
				return m, tea.Batch(m.spinner.Tick, m.confirm)
			}
			var cmd tea.Cmd
			m.confirmationTextInput, cmd = m.confirmationTextInput.Update(msg)
			return m, cmd
		}
	case tea.WindowSizeMsg:
		m.emailTextInput.Width = msg.Width - 2
	}
	cmds := []tea.Cmd{}
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m loginModel) View() string {
	content := titleStyle.Render("Welcome to Puffery") + "\n\n"

	if m.currentStep >= loginStepHasAccount {
		content += m.hasAccountQuestion.View() + "\n"
	}
	if m.currentStep >= loginStepEmail {
		content += m.emailTextInput.View() + "\n"
	}
	if m.currentStep == loginStepConfirmation {
		content += "\n"
		content += "We sent you a link to log in by email.\n"
		content += m.confirmationTextInput.View() + "\n"
	}

	if m.isBusy {
		content += m.spinner.View() + "\n"
	}

	return content
}

type loginModelSuccessfulMsg struct{}
type loginModelFailedMsg struct{ err error }
type loginModelNeedsConfirmationMsg struct{}

func (m loginModel) register() tea.Msg {
	_, err := Api.Register(m.emailTextInput.Value())
	if err != nil {
		return loginModelFailedMsg{err}
	} else if m.emailTextInput.Value() == "" {
		return loginModelNeedsConfirmationMsg{}
	} else {
		return loginModelNeedsConfirmationMsg{}
	}
}

func (m loginModel) login() tea.Msg {
	_, err := Api.Login(m.emailTextInput.Value())
	if err != nil {
		return loginModelFailedMsg{err}
	} else {
		return loginModelNeedsConfirmationMsg{}
	}
}

func (m loginModel) confirm() tea.Msg {
	if m.hasAccountQuestion.answer {
		return m.confirmLogin()
	} else {
		return m.confirmRegister()
	}
}

func (m loginModel) confirmLogin() tea.Msg {
	_, err := Api.ConfirmLogin(m.confirmationTextInput.Value())
	if err != nil {
		return loginModelFailedMsg{err}
	} else {
		return loginModelSuccessfulMsg{}
	}
}

func (m loginModel) confirmRegister() tea.Msg {
	_, err := Api.ConfirmEmail(m.confirmationTextInput.Value())
	if err != nil {
		return loginModelFailedMsg{err}
	} else {
		return loginModelSuccessfulMsg{}
	}
}
