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

type operationFailedMsg struct {
	err error
}

type viewMessagesOfChannelMsg struct {
	channel puffery.Channel
}

type viewMessagesOfAllChannelsMsg struct{}

type didLoginMsg struct {
	puffery.TokenResponse
}
