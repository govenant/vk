package vk

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	AndroidClientID     = "2274003"              // Android app client_id
	AndroidClientSecret = "hHbZxrka2uZ6jB1inYsH" // Android app client_secret

	WindowsClientID     = "3697615"              // Windows app client_id
	WindowsClientSecret = "AlVXZFMUqyrnABp8ncuU" // Windows app client_secret

	IOSClientID     = "3140623"              // iOS app client_id
	IOSClientSecret = "VeWdmVclDCtn6ihuP1nt" // iOS app client_secret

	AccessTokenURL = "https://oauth.vk.com/access_token"
)

func GetServiceToken(ClientID, ClientSecret string) (string, error) {
	var url, _ = url.Parse(AccessTokenURL)
	var query = url.Query()

	query.Set("client_id", ClientID)
	query.Set("client_secret", ClientSecret)
	query.Set("grant_type", "client_credentials")
	query.Set("v", Version)

	url.RawQuery = query.Encode()

	response, err := http.Get(url.String())

	if err != nil {
		return "", err
	}

	var result = &AccessTokenResponse{}

	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", result.Error
	}

	return result.AccessToken, nil
}
