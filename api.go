package puffery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	Root  string
	Token string
}

func (a *Api) get(requestPath string) ([]byte, error) {
	requestUrl := a.Root + "/" + requestPath
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+a.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(bodyBytes))
		return nil, err
	}
	return bodyBytes, nil
}

func (a *Api) delete(requestPath string) ([]byte, error) {
	requestUrl := a.Root + "/" + requestPath
	req, err := http.NewRequest(http.MethodDelete, requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+a.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(bodyBytes))
		return nil, err
	}
	return bodyBytes, nil
}

func (a *Api) post(requestPath string, body []byte) ([]byte, error) {
	requestUrl := a.Root + "/" + requestPath
	req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+a.Token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(bodyBytes))
		return nil, err
	}
	return bodyBytes, nil
}

func (a *Api) put(requestPath string, body []byte) ([]byte, error) {
	requestUrl := a.Root + "/" + requestPath
	req, err := http.NewRequest(http.MethodPut, requestUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+a.Token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(bodyBytes))
		return nil, err
	}
	return bodyBytes, nil
}

func (a *Api) Hello() error {
	_, err := a.get("/hello")
	return err
}

func (a *Api) Register(email string) (TokenResponse, error) {
	body, err := json.Marshal(CreateUserRequest{Email: email})
	if err != nil {
		return TokenResponse{}, err
	}
	bodyBytes, err := a.post("/api/v1/register", body)
	if err != nil {
		return TokenResponse{}, err
	}
	var tokenResponse TokenResponse
	err = json.Unmarshal(bodyBytes, &tokenResponse)
	return tokenResponse, err
}

func (a *Api) Login(email string) (LoginAttemptResponse, error) {
	body, err := json.Marshal(LoginUserRequest{Email: email})
	if err != nil {
		return LoginAttemptResponse{}, err
	}
	bodyBytes, err := a.post("/api/v1/login", body)
	if err != nil {
		return LoginAttemptResponse{}, err
	}
	var loginAttemptResponse LoginAttemptResponse
	err = json.Unmarshal(bodyBytes, &loginAttemptResponse)
	return loginAttemptResponse, err
}

func (a *Api) ConfirmLogin(confirmationId string) (TokenResponse, error) {
	bodyBytes, err := a.post("/api/v1/confirmations/login/"+confirmationId, nil)
	if err != nil {
		return TokenResponse{}, err
	}
	var tokenResponse TokenResponse
	err = json.Unmarshal(bodyBytes, &tokenResponse)
	if err != nil {
		fmt.Println(string(bodyBytes))
		return TokenResponse{}, err
	}
	return tokenResponse, nil
}

func (a *Api) ConfirmEmail(confirmationId string) (ConfirmedEmailResponse, error) {
	bodyBytes, err := a.post("/api/v1/confirmations/email/"+confirmationId, nil)
	if err != nil {
		return ConfirmedEmailResponse{}, err
	}
	var confirmedEmailResponse ConfirmedEmailResponse
	err = json.Unmarshal(bodyBytes, &confirmedEmailResponse)
	return confirmedEmailResponse, err
}

func (a *Api) Profile() (User, error) {
	bodyBytes, err := a.get("/api/v1/profile")
	if err != nil {
		return User{}, err
	}
	var user UserResponse
	err = json.Unmarshal(bodyBytes, &user)
	return user, err
}

func (a *Api) UpdateProfile(req UpdateProfileRequest) (User, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return User{}, err
	}
	bodyBytes, err := a.put("/api/v1/profile", body)
	if err != nil {
		return User{}, err
	}
	var user UserResponse
	err = json.Unmarshal(bodyBytes, &user)
	return user, err
}

func (a *Api) CreateDevice(req CreateDeviceRequest) (DeviceResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return DeviceResponse{}, err
	}
	bodyBytes, err := a.post("/api/v1/devices", body)
	if err != nil {
		return DeviceResponse{}, err
	}
	var device DeviceResponse
	err = json.Unmarshal(bodyBytes, &device)
	return device, err
}

func (a *Api) UpdateDevice(id string, req UpdateDeviceRequest) (DeviceResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return DeviceResponse{}, err
	}
	bodyBytes, err := a.put("/api/v1/devices/"+id, body)
	if err != nil {
		return DeviceResponse{}, err
	}
	var device DeviceResponse
	err = json.Unmarshal(bodyBytes, &device)
	return device, err
}

func (a *Api) CreateChannel(req CreateChannelRequest) (Channel, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return SubscribedChannelResponse{}, err
	}
	bodyBytes, err := a.post("/api/v1/channels", body)
	if err != nil {
		return SubscribedChannelResponse{}, err
	}
	var channel SubscribedChannelResponse
	err = json.Unmarshal(bodyBytes, &channel)
	return channel, err
}

func (a *Api) GetChannel(id string) (Channel, error) {
	bodyBytes, err := a.get("/api/v1/channels/" + id)
	if err != nil {
		return SubscribedChannelResponse{}, err
	}
	var channel SubscribedChannelResponse
	err = json.Unmarshal(bodyBytes, &channel)
	return channel, err
}

func (a *Api) UpdateChannel(id string, req UpdateSubscriptionRequest) (Channel, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return SubscribedChannelResponse{}, err
	}
	bodyBytes, err := a.put("/api/v1/channels/"+id, body)
	if err != nil {
		return SubscribedChannelResponse{}, err
	}
	var channel SubscribedChannelResponse
	err = json.Unmarshal(bodyBytes, &channel)
	return channel, err
}

func (a *Api) UnsubscribeChannel(id string) (SubscribedChannelDeletedResponse, error) {
	bodyBytes, err := a.delete("/api/v1/channels/" + id)
	if err != nil {
		return SubscribedChannelDeletedResponse{}, err
	}
	var channel SubscribedChannelDeletedResponse
	err = json.Unmarshal(bodyBytes, &channel)
	return channel, err
}

func (a *Api) ChannelStats(id string) (SubscribedChannelDeletedResponse, error) {
	bodyBytes, err := a.delete("/api/v1/channels/" + id + "/stats")
	if err != nil {
		return SubscribedChannelDeletedResponse{}, err
	}
	var channel SubscribedChannelDeletedResponse
	err = json.Unmarshal(bodyBytes, &channel)
	return channel, err
}

func (a *Api) Channels() ([]Channel, error) {
	bodyBytes, err := a.get("/api/v1/channels")
	if err != nil {
		return nil, err
	}
	var channels []Channel
	err = json.Unmarshal(bodyBytes, &channels)
	return channels, err
}

func (a *Api) OwnChannels() ([]Channel, error) {
	bodyBytes, err := a.get("/api/v1/channels/own")
	if err != nil {
		return nil, err
	}
	var channels []Channel
	err = json.Unmarshal(bodyBytes, &channels)
	return channels, err
}

func (a *Api) SharedChannels() ([]Channel, error) {
	bodyBytes, err := a.get("/api/v1/channels/shared")
	if err != nil {
		return nil, err
	}
	var channels []Channel
	err = json.Unmarshal(bodyBytes, &channels)
	return channels, err
}

func (a *Api) PublicNotify(notifyKey string, req CreateMessageRequest) (NotifyMessageResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return NotifyMessageResponse{}, err
	}
	bodyBytes, err := a.post("/api/v1/notify/"+notifyKey, body)
	if err != nil {
		return NotifyMessageResponse{}, err
	}
	var message NotifyMessageResponse
	err = json.Unmarshal(bodyBytes, &message)
	return message, err
}

func (a *Api) MessagesOfAllChannels() ([]Message, error) {
	bodyBytes, err := a.get("/api/v1/channels/messages")
	if err != nil {
		return nil, err
	}
	var messages []Message
	err = json.Unmarshal(bodyBytes, &messages)
	return messages, err
}

func (a *Api) MessagesOfChannel(channel Channel) ([]Message, error) {
	bodyBytes, err := a.get("/api/v1/channels/" + channel.Id + "/messages")
	if err != nil {
		return nil, err
	}
	var messages []Message
	err = json.Unmarshal(bodyBytes, &messages)
	return messages, err
}

func (a *Api) CreateMessage(channel Channel, req CreateMessageRequest) (Message, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return MessageResponse{}, err
	}
	bodyBytes, err := a.post("/api/v1/channels/"+channel.Id+"/messages", body)
	if err != nil {
		return MessageResponse{}, err
	}
	var message MessageResponse
	err = json.Unmarshal(bodyBytes, &message)
	return message, err
}

func (a *Api) SubscribeToChannel(req CreateSubscriptionRequest) (Channel, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return SubscribedChannelResponse{}, err
	}
	bodyBytes, err := a.post("/api/v1/channels/subscribe", body)
	if err != nil {
		return SubscribedChannelResponse{}, err
	}
	var channel SubscribedChannelResponse
	err = json.Unmarshal(bodyBytes, &channel)
	return channel, err
}
