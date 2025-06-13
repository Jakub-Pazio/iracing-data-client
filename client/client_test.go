package client

import "testing"

func TestNewClient(t *testing.T) {
	_, err := NewClient("mail@iracing.com", "somepassword")

	if err != nil {
		t.Error("error is not nil but should be ")
	}
}
