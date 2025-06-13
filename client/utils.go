package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

const (
	authURL = "https://members-ng.iracing.com/auth"
	baseURL = "https://members-ng.iracing.com/data/"
)

func (c *Client) authClient() error {
	type reqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	rb, err := json.Marshal(reqBody{Email: c.email, Password: c.password})
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Post(authURL, "application/json", bytes.NewReader(rb))

	if resp.StatusCode != 200 {
		return fmt.Errorf("could not authenticate, status code: %d", resp.StatusCode)
	}

	if err != nil {
		return err
	}
	return nil
}

func followLink[T any](client *Client, url string) (T, error) {
	var t T
	type linkResp struct {
		Link string `json:"link"`
	}

	resp, err := client.httpClient.Get(url)
	// we are probably not authed, so we try to auth and then perform the req again
	if resp.StatusCode != 200 {
		err := client.authClient()
		if err != nil {
			return t, err
		}
		resp, err = client.httpClient.Get(url)
	}
	if err != nil {
		return t, fmt.Errorf("could not get car classes: %s", err)
	}

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return t, fmt.Errorf("could not read the body: %s", err)
	}

	var lr linkResp
	if err := json.Unmarshal(bb, &lr); err != nil {
		return t, fmt.Errorf("could not unmarshal request: %s", err)
	}

	resp, err = client.httpClient.Get(lr.Link)
	if err != nil {
		return t, fmt.Errorf("could not make last request: %s", err)
	}
	bb, err = io.ReadAll(resp.Body)
	if err != nil {
		return t, fmt.Errorf("could not read data about cars: %s", err)
	}
	if err := json.Unmarshal(bb, &t); err != nil {
		return t, fmt.Errorf("error unmarsharling data: %s", err)
	}
	return t, nil
}

// followLink2 exists because sometimes iRacing provides different response
// for a link to follow to get the data back
func followLink2[T any](client *Client, url string) (T, error) {
	var t T
	type linkResp struct {
		Link string `json:"data_url"`
	}

	resp, err := client.httpClient.Get(url)
	// we are probably not authed, so we try to auth and then perform the req again
	if resp.StatusCode != 200 {
		err := client.authClient()
		if err != nil {
			return t, err
		}
		resp, err = client.httpClient.Get(url)
	}
	if err != nil {
		return t, fmt.Errorf("could not get car classes: %s", err)
	}

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return t, fmt.Errorf("could not read the body: %s", err)
	}

	var lr linkResp
	if err := json.Unmarshal(bb, &lr); err != nil {
		return t, fmt.Errorf("could not unmarshal request: %s", err)
	}

	resp, err = client.httpClient.Get(lr.Link)
	if err != nil {
		return t, fmt.Errorf("could not make last request: %s", err)
	}
	bb, err = io.ReadAll(resp.Body)
	if err != nil {
		return t, fmt.Errorf("could not read data about cars: %s", err)
	}
	if err := json.Unmarshal(bb, &t); err != nil {
		return t, fmt.Errorf("error unmarsharling data: %s", err)
	}
	return t, nil
}

func createURL(path string) string {
	return baseURL + path
}
