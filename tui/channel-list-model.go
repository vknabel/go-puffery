package tui

import (
	"github.com/charmbracelet/bubbles/key"
	list "github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	puffery "github.com/vknabel/go-puffery"
	"github.com/vknabel/go-puffery/nav"
)

type channelListModel struct {
	channelListView list.Model
}

func initialChannelListModel() channelListModel {
	channelListViewDelegate := NewStyledDefaultListDelegate()
	m := channelListModel{
		channelListView: list.New(nil, channelListViewDelegate, 0, 0),
	}
	m.channelListView.Title = "Channels"
	m.channelListView.Styles.Title = titleStyle

	m.channelListView.View()
	m.channelListView.SetSpinner(spinner.Dot)
	m.channelListView.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("⏎", "details"),
			),
		}
	}
	m.channelListView.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("⏎", "details"),
			),
		}
	}
	return m
}

type channelItem struct {
	channel *puffery.Channel
}

func (i channelItem) Title() string {
	if i.channel != nil {
		return i.channel.Title
	}
	return "All channels"
}
func (i channelItem) Description() string {
	if i.channel == nil {
		return "messages of all channels"
	}
	description := "RO: " + i.channel.ReceiveOnlyKey
	if i.channel != nil && i.channel.NotifyKey != nil {
		description += " / N: " + *i.channel.NotifyKey
	}
	return description
}
func (i channelItem) FilterValue() string {
	return i.Description()
}

func (m channelListModel) Init() tea.Cmd {
	return func() tea.Msg {
		channels, err := Api.Channels()
		if err != nil {
			return nav.PagePushMsg{Page: initialErrorModel(err)}
		}
		return didLoadChannelsMsg{channels}
	}
}

func (m channelListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			if selected, ok := m.channelListView.SelectedItem().(channelItem); ok {
				return m, nav.Push(initialMessageModel(selected.channel))
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.channelListView.SetSize(msg.Width-h, msg.Height-v)
	case didLoadChannelsMsg:
		channelItems := make([]list.Item, len(msg.channels)+1)
		channelItems[0] = channelItem{nil}
		for i, c := range msg.channels {
			current := c
			channelItems[i+1] = channelItem{&current}
		}
		m.channelListView.SetItems(channelItems)
	}

	var cmd tea.Cmd
	m.channelListView, cmd = m.channelListView.Update(msg)
	return m, cmd
}

func (m channelListModel) View() string {
	return m.channelListView.View()
}
