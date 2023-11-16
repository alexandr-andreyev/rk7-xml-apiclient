package rk7client

import (
	"crypto/tls"
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
