package vk

import (
	"net/http"
	"testing"
)

func GetToken() string {
	token, err := GetServiceToken(IOSClientID, IOSClientSecret)

	if err != nil {
		panic(err)
	}

	return token
}

func AppBalance(client Client) (int, error) {
	type Response struct {
		ErrorResponse

		Response int
	}

	var response = &Response{}

	err := client.Request("secure.getAppBalance", Payload{"client_secret": IOSClientSecret}, response)

	if err != nil {
		return 0, err
	}

	if response.Error != nil {
		return 0, response.Error
	}

	return 0, nil
}

func TestAPI(test *testing.T) {
	var OnResponse = func(response *http.Response) error {
		return nil
	}

	var client = Client{
		AccessToken: GetToken(),
		Language:    "en",
		Version:     Version,
		OnResponse:  OnResponse,
	}

	if balance, err := AppBalance(client); err != nil {
		test.Fatal(err)
	} else if balance != 0 {
		test.Errorf("Incorrect ID!")
		test.Errorf("Expected: 0")
		test.Errorf("Got: %d", balance)
		return
	}
}

func TestAPIError(test *testing.T) {
	var OnResponse = func(response *http.Response) error {
		return &Error{Code: 113}
	}

	var client = Client{
		AccessToken: GetToken(),
		OnResponse:  OnResponse,
	}

	if _, err := AppBalance(client); err == nil {
		test.Fatalf("Error aren't occured!")
	} else if err, ok := err.(*Error); !ok {
		test.Fatalf("Incorrect error type!")
	} else if err.Code != 113 {
		test.Errorf("Incorrect error code!\n")
		test.Errorf("Expected: 113\n")
		test.Errorf("Got: %d\n", err.Code)
	}
}
