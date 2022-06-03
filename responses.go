package puffery

type User = UserResponse
type UserResponse struct {
	ID          string `json:"id"`
	IsConfirmed bool   `json:"isConfirmed"`
	Email       string `json:"email"`
}

type TokenResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type LoginAttemptResponse struct{}
type ConfirmedEmailResponse struct{}

type Channel = SubscribedChannelResponse
type SubscribedChannelResponse struct {
	Id             string  `json:"id"`
	Title          string  `json:"title"`
	IsSilent       bool    `json:"isSilent"`
	ReceiveOnlyKey string  `json:"receiveOnlyKey"`
	NotifyKey      *string `json:"notifyKey"`
}

type SubscribedChannelStatisticsResponse struct {
	Notifiers int `json:"notifiers"`
	Receivers int `json:"receivers"`
	Messages  int `json:"messages"`
}

type SubscribedChannelDeletedResponse struct{}

type Message = MessageResponse
type MessageResponse struct {
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	ColorName string `json:"colorName"`
	ID        string `json:"id"`
	Title     string `json:"title"`
	Channel   string `json:"channel"`
}

type NotifyMessageResponse struct {
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	Color     string `json:"color"`
	ID        string `json:"id"`
	Title     string `json:"title"`
}

type DeviceResponse struct {
	ID           string `json:"id"`
	Token        string `json:"token"`
	IsProduction bool   `json:"isProduction"`
}
