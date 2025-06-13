package client

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

type Client struct {
	httpClient http.Client
	email      string
	password   string
}

func NewClient(email string, password string) (Client, error) {
	h := sha256.New()
	h.Write([]byte(password + email))
	hashedPassword := (base64.StdEncoding.EncodeToString((h.Sum(nil))))

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return Client{}, fmt.Errorf("could not create cookie jar: %s", err)
	}
	return Client{
		httpClient: http.Client{
			Jar: jar,
		},
		email:    email,
		password: hashedPassword,
	}, nil
}

func (c *Client) GetAllCars() []int {
	return []int{42, 1337}
}
