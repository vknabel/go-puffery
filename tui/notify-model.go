package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	puffery "github.com/vknabel/go-puffery"
	"github.com/vknabel/go-puffery/nav"
)

type notifyModel struct {
	height         int
	channel        puffery.Channel
	loadingSpinner spinner.Model
	isLoading      bool
	titleInput     textinput.Model
	bodyInput      textarea.Model
	help           help.Model
	keys           notifyKeyMap
}

type notifyKeyMap struct {
	jump key.Binding
	send key.Binding
	back key.Binding
}

func (m notifyModel) ShortHelp() []key.Binding {
	if m.bodyInput.Focused() {
		return []key.Binding{
			m.keys.jump,
			m.keys.send,
			m.keys.back,
		}
	} else {
		return []key.Binding{
			m.keys.jump,
			m.keys.back,
		}
	}
}

func (n notifyModel) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		n.ShortHelp(),
	}
}

func initialNotifyModel(channel puffery.Channel) notifyModel {
	title := textinput.NewModel()
	title.Placeholder = "Title"
	title.Focus()

	body := textarea.New()
	body.Placeholder = "Body"
	body.ShowLineNumbers = false
	body.BlurredStyle.Placeholder = body.BlurredStyle.Placeholder.
		Foreground(title.PlaceholderStyle.GetForeground())
	body.FocusedStyle.Placeholder = body.FocusedStyle.Placeholder.
		Foreground(title.PlaceholderStyle.GetForeground())

	loading := spinner.New()
	loading.Spinner = spinner.Dot

	return notifyModel{
		channel:        channel,
		titleInput:     title,
		bodyInput:      body,
		loadingSpinner: loading,
		help:           help.New(),
		keys: notifyKeyMap{
			jump: key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "jump")),
			send: key.NewBinding(key.WithKeys(), key.WithHelp("⏎⏎", "send")),
			back: backKeyBinding,
		},
	}
}

func (m notifyModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m notifyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// case multitext.TrailingNewlines:
	// 	if msg.Count == 1 {
	// 		break
	// 	}
	// 	m.bodyInput.Blur()
	// 	if m.bodyInput.Value() == "" || m.titleInput.Value() == "" {
	// 		m.titleInput.Focus()
	// 		return m, textinput.Blink
	// 	}
	// 	m.isLoading = true
	// 	return m, tea.Batch(m.loadingSpinner.Tick, m.SendNotification())
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.bodyInput.SetWidth(msg.Width - 4)
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
					value := m.bodyInput.Value()
					if len(value) > 0 && value[len(value)-1] == '\n' {
						m.bodyInput.Blur()
						if m.bodyInput.Value() == "" || m.titleInput.Value() == "" {
							m.titleInput.Focus()
							return m, textinput.Blink
						}
						m.isLoading = true
						return m, tea.Batch(m.loadingSpinner.Tick, m.SendNotification())
					}
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
	view := "  " + titleStyle.Render("notify "+m.channel.Title)
	view += "\n\n"
	view += "  " + m.titleInput.View()
	view += "\n\n"
	view += lipgloss.NewStyle().MarginLeft(2).Render(m.bodyInput.View())

	if m.isLoading {
		view += "  " + m.loadingSpinner.View() + " Sending..."
	}

	availableHeight := m.height
	availableHeight -= lipgloss.Height(view)
	view += lipgloss.NewStyle().MarginLeft(2).Render(
		lipgloss.PlaceVertical(
			availableHeight,
			lipgloss.Bottom,
			m.help.View(m),
		),
	)
	return view
}

func (m *notifyModel) SendNotification() tea.Cmd {
	m.isLoading = true
	return func() tea.Msg {
		_, err := Api.CreateMessage(m.channel, puffery.CreateMessageRequest{
			Title: m.titleInput.Value(),
			Body:  strings.Trim(m.bodyInput.Value(), " \n"),
		})
		if err != nil {
			return nav.PagePushMsg{Page: initialErrorModel(err)}
		}
		return nav.PagePopMsg{}
	}
}
