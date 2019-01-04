package vk

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const Version = "5.92"

type Payload map[string]string
type OnResponseFunc func(*http.Response) error

type Client struct {
	AccessToken string
	Version     string
	Language    string
	OnResponse  OnResponseFunc
}

func (client *Client) Request(method string, parameters Payload, storage interface{}) error {
	var payload = url.Values{}

	for key, value := range parameters {
		payload.Set(key, value)
	}

	if client.AccessToken != "" {
		payload.Set("access_token", client.AccessToken)
	}

	if client.Language != "" {
		payload.Set("lang", client.Language)
	}

	if client.Version != "" {
		payload.Set("v", client.Version)
	} else {
		payload.Set("v", Version)
	}

	response, err := http.PostForm(
		"https://api.vk.com/method/"+method,
		payload,
	)

	if err != nil {
		return err
	}

	if client.OnResponse != nil {
		if err := client.OnResponse(response); err != nil {
			return err
		}
	}

	if storage != nil {
		defer response.Body.Close()

		return json.NewDecoder(response.Body).Decode(&storage)
	}

	return nil
}
