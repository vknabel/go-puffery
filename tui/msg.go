package tui

import (
	puffery "github.com/vknabel/go-puffery"
)

type didLoadChannelsMsg struct {
	channels []puffery.Channel
}

type didLoadMessagesMsg struct {
	messages []puffery.Message
}

type viewMessagesOfChannelMsg struct {
	channel puffery.Channel
}

type viewMessagesOfAllChannelsMsg struct{}
