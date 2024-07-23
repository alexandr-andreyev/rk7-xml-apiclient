package rk7client

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) newRequest(method string, body interface{}) (*http.Request, error) {
	url := fmt.Sprintf("https://%s:%d/rk7api/v0/xmlinterface.xml", c.CashServerIP, c.CashServerPort)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := xml.NewEncoder(buf).Encode(body)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/xml")
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// TODO add error if cash server down
func (с *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := с.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)

	// fmt.Println(string(body))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		err = xml.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			s := fmt.Sprintf("xmkl decode error: %s", err)
			err = errors.New(s)
			return nil, err
		}
		return resp, nil
	}
	if resp.StatusCode == 401 {
		return nil, errUnauthorized
	}
	s := fmt.Sprintf("Unknown error. Status: %d", resp.StatusCode)
	err = errors.New(s)
	return nil, err
}
