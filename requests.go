package puffery

type CreateUserRequest struct {
	Email *string `json:"email"`
}

type UpdateProfileRequest struct {
	Email string `json:"email"`
}

type LoginUserRequest struct {
	Email string `json:"email"`
}

type CreateChannelRequest struct {
	Title    string `json:"title"`
	IsSilent bool   `json:"isSilent"`
}

type CreateSubscriptionRequest struct {
	ReceiveOrNotifyKey string `json:"receiveOrNotifyKey"`
	IsSilent           bool   `json:"isSilent"`
}

type UpdateSubscriptionRequest struct {
	IsSilent bool `json:"isSilent"`
}

type CreateMessageRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Color string `json:"color"`
}

type CreateDeviceRequest struct {
	Token        string `json:"token"`
	IsProduction bool   `json:"isProduction"`
}

type UpdateDeviceRequest struct {
	IsProduction bool `json:"isProduction"`
}
