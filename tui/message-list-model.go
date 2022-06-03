package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	list "github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	puffery "github.com/vknabel/go-puffery"
	"github.com/vknabel/go-puffery/nav"
)

type messageListModel struct {
	channel         *puffery.Channel
	messageListView list.Model
}

func initialMessageModel(channel *puffery.Channel) messageListModel {
	m := messageListModel{
		channel:         channel,
		messageListView: list.NewModel(nil, list.NewDefaultDelegate(), 0, 0),
	}
	if channel != nil {
		m.messageListView.Title = channel.Title
	} else {
		m.messageListView.Title = "All messages"
	}

	m.messageListView.SetSpinner(spinner.Dot)
	m.messageListView.DisableQuitKeybindings()
	m.messageListView.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("+", "n"),
				key.WithHelp("+/n", "new"),
			),
			key.NewBinding(
				key.WithKeys("esc"),
				key.WithHelp("esc", "back"),
			),
		}
	}
	m.messageListView.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("esc"),
				key.WithHelp("esc", "back"),
			),
			key.NewBinding(
				key.WithKeys("+", "n"),
				key.WithHelp("+/n", "new"),
			),
			key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reload"),
			),
		}
	}
	return m
}

type messageItem struct {
	message puffery.Message
}

func (i messageItem) Title() string       { return i.message.Title }
func (i messageItem) FilterValue() string { return i.message.Title + " " + i.message.Body }
func (i messageItem) Description() string {
	return strings.Trim(strings.ReplaceAll(i.message.Body, "\n", " "), " ")
}

func (m messageListModel) Init() tea.Cmd {
	var cmd tea.Cmd
	if m.channel != nil {
		cmd = m.LoadMessagesOfChannel(*m.channel)
	} else {
		cmd = m.LoadMessagesOfAllChannels()
	}
	return tea.Batch(
		m.messageListView.StartSpinner(),
		cmd,
	)
}

func (m messageListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.messageListView.SetSize(msg.Width-h, msg.Height-v)
	case didLoadMessagesMsg:
		messageItems := make([]list.Item, len(msg.messages))
		for i, message := range msg.messages {
			messageItems[i] = messageItem{message}
		}
		m.messageListView.SetItems(messageItems)
		m.messageListView.StopSpinner()
	case viewMessagesOfChannelMsg:
		return m, tea.Batch(
			m.messageListView.StartSpinner(),
			m.LoadMessagesOfChannel(msg.channel),
		)
	case viewMessagesOfAllChannelsMsg:
		return m, tea.Batch(
			m.messageListView.StartSpinner(),
			m.LoadMessagesOfAllChannels(),
		)
	case tea.KeyMsg:
		switch msg.String() {
		case "+", "n":
			if m.channel != nil {
				notify := initialNotifyModel(*m.channel)
				return m, nav.Push(notify)
			}
		}
	case nav.PageRestoreMsg:
		if m.channel != nil {
			return m, m.LoadMessagesOfChannel(*m.channel)
		} else {
			return m, m.LoadMessagesOfAllChannels()
		}
	}

	var cmd tea.Cmd
	m.messageListView, cmd = m.messageListView.Update(msg)
	return m, cmd
}

func (m messageListModel) View() string {
	return m.messageListView.View()
}

func (m *messageListModel) LoadMessagesOfChannel(channel puffery.Channel) tea.Cmd {
	m.messageListView.Title = channel.Title
	m.channel = &channel
	return func() tea.Msg {
		messages, err := Api.MessagesOfChannel(channel)
		if err != nil {
			return nav.PagePushMsg{Page: initialErrorModel(err)}
		}
		return didLoadMessagesMsg{messages}
	}
}

func (m *messageListModel) LoadMessagesOfAllChannels() tea.Cmd {
	m.messageListView.Title = "All messages"
	m.channel = nil
	return func() tea.Msg {
		messages, err := Api.MessagesOfAllChannels()
		if err != nil {
			return nav.PagePushMsg{Page: initialErrorModel(err)}
		}
		return didLoadMessagesMsg{messages}
	}
}
