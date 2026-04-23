// Package rk7client предоставляет клиент для работы с XML API кассового сервера R-Keeper 7.
//
// Пример использования:
//
//	client := rk7client.NewClient("192.168.0.91", 8086, "http", "http")
//
//	info, err := client.GetSystemInfo()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(info.CommandResult.SystemInfo.ShiftDate)
package rk7client

import (
	"crypto/tls"
	"net/http"
	"time"
)

// Client — HTTP-клиент для взаимодействия с XML API R-Keeper 7.
type Client struct {
	CashServerIP   string
	CashServerPort int
	Username       string
	Password       string
	HTTPClient     *http.Client
}

// NewClient создаёт клиент для кассового сервера по адресу ip:port с указанными учётными данными.
// TLS-сертификат сервера не проверяется (самоподписанный).
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
