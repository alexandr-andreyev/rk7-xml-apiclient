package rk7client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	CashServerIP   string
	CashServerPort int
	Username       string
	Password       string
	HTTPClient     *http.Client
}

func NewClient(CashServerIP string, Port int, UserName, Password string) *Client {
	return &Client{
		CashServerIP:   CashServerIP,
		CashServerPort: Port,
		Username:       UserName,
		Password:       Password,
		HTTPClient: &http.Client{
			Timeout:   time.Minute,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
	}
}

func (c *Client) newRequest(method string, body interface{}) (*http.Request, error) {
	url := fmt.Sprintf("https://%s:%d/rk7api/v0/xmlinterface.xml", c.CashServerIP, c.CashServerPort)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/xml")
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

func (с *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := с.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		body, err := io.ReadAll(resp.Body)

		fmt.Println(string(body))
		if err != nil {
			return nil, err
		}
		return resp, err
	}
	if resp.StatusCode == 401 {
		s := fmt.Sprintf("Unauthorized, status: %d", resp.StatusCode)
		err = errors.New(s)
		return nil, err
	}
	s := fmt.Sprintf("Unknown error. Status: %d", resp.StatusCode)
	err = errors.New(s)
	return nil, err
}
