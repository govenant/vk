package vk

import "testing"

func TestGetServiceToken(test *testing.T) {
	if token, err := GetServiceToken(WindowsClientID, WindowsClientSecret); err != nil {
		test.Fatal(err)
	} else if token == "" {
		test.Fatal("Empty token!")
	}
}
