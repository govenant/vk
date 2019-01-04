package main

import (
	"log"
	"net/http"
	"os"

	"github.com/govenant/vk"
)

func UsernameToID(client vk.Client, username string) (int, error) {
	type Response struct {
		vk.ErrorResponse

		Response []struct {
			ID int `json:"id"`
		} `json:"response"`
	}

	var response = &Response{}

	err := client.Request("users.get", vk.Payload{"user_ids": username}, response)

	if err != nil {
		return 0, err
	}

	if response.Error != nil {
		return 0, response.Error
	}

	return response.Response[0].ID, nil
}

// Main function.
func main() {
	var OnResponse = func(response *http.Response) error {
		log.Print(response.Status)

		return nil
	}

	var client = vk.Client{
		AccessToken: os.Getenv("VK_ACCESS_TOKEN"),
		Language:    "en",
		Version:     "5.92",
		OnResponse:  OnResponse, // optional (logging, etc)
	}

	id, _ := UsernameToID(client, "durov")
	log.Println(id) // 1
}
