package goipapi

import (
	"bytes"
	"fmt"
	"net/http"
)

const apiURL = "http://ip-api.com/"

type Client struct {
	Format string
}

func NewClient(format string) *Client {
	return &Client{Format: format}
}

//LookupIP some ....
func (c *Client) LookupIP(ip string) (jsonString string, err error) {
	url := fmt.Sprintf("%s/%s/%s", apiURL, c.Format, ip)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	buff := new(bytes.Buffer)
	buff.ReadFrom(resp.Body)
	jsonString = buff.String()

	return
}
