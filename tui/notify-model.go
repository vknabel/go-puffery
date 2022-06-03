package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	puffery "github.com/vknabel/go-puffery"
	"github.com/vknabel/go-puffery/multitext"
	"github.com/vknabel/go-puffery/nav"
)

type notifyModel struct {
	channel puffery.Channel

	loadingSpinner spinner.Model
	isLoading      bool

	titleInput textinput.Model
	bodyInput  multitext.Model
}

func initialNotifyModel(channel puffery.Channel) notifyModel {
	title := textinput.NewModel()
	title.Placeholder = "Title"
	title.Focus()

	bodyLine := textinput.New()
	bodyLine.Placeholder = "Body"
	body := multitext.New(bodyLine)

	loading := spinner.New()
	loading.Spinner = spinner.Dot

	return notifyModel{
		channel:        channel,
		titleInput:     title,
		bodyInput:      body,
		loadingSpinner: loading,
	}
}

func (m notifyModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m notifyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case operationFailedMsg:
		m.isLoading = false
	case multitext.TrailingNewlines:
		if msg.Count == 1 {
			break
		}
		m.bodyInput.Blur()
		if m.bodyInput.Value() == "" || m.titleInput.Value() == "" {
			m.titleInput.Focus()
			return m, textinput.Blink
		}
		m.isLoading = true
		return m, tea.Batch(m.loadingSpinner.Tick, m.SendNotification())
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab, tea.KeyEnter:
			switch {
			case m.titleInput.Focused():
				m.titleInput.Blur()
				m.bodyInput.Focus()
				return m, textinput.Blink
			case m.bodyInput.Focused():
				if msg.Type == tea.KeyEnter {
					break
				}
				m.bodyInput.Blur()
				if m.bodyInput.Value() == "" || m.titleInput.Value() == "" {
					m.titleInput.Focus()
					return m, textinput.Blink
				}
				return m, m.SendNotification()
			}
		}
	}
	cmds := []tea.Cmd{}
	var cmd tea.Cmd
	m.titleInput, cmd = m.titleInput.Update(msg)
	cmds = append(cmds, cmd)
	m.bodyInput, cmd = m.bodyInput.Update(msg)
	cmds = append(cmds, cmd)
	m.loadingSpinner, cmd = m.loadingSpinner.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m notifyModel) View() string {
	title := lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("230")).
		Padding(0, 1)

	var loadingText string
	if m.isLoading {
		loadingText = m.loadingSpinner.View() + " Sending..."
	} else {
		loadingText = ""
	}

	return fmt.Sprintf(
		"  %s\n\n  %s\n\n  %s\n%s\n",
		title.Render(m.channel.Title),
		m.titleInput.View(),
		m.bodyInput.View(),
		loadingText,
	)
}

func (m *notifyModel) SendNotification() tea.Cmd {
	m.isLoading = true
	return func() tea.Msg {
		_, err := Api.CreateMessage(m.channel, puffery.CreateMessageRequest{
			Title: m.titleInput.Value(),
			Body:  strings.Trim(m.bodyInput.Value(), " \n"),
		})
		if err != nil {
			return operationFailedMsg{err}
		}
		return nav.PagePopMsg{}
	}
}
