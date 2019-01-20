package goipapi

import (
	"bytes"
	"fmt"
	"net/http"
)

const apiURL = "http://ip-api.com/"

type Client struct {
	Format string
	Fields string
}

func NewClient(format string) *Client {
	return &Client{
		Format: format,
		Fields: "status,message,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,isp,org,as,reverse,mobile,proxy,query",
	}
}

//LookupIP some ....
func (c *Client) LookupIP(ip string) (jsonString string, err error) {
	url := fmt.Sprintf("%s/%s/%s?fields=%s", apiURL, c.Format, ip, c.Fields)

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

func (c *Client) SetFields(fields string) *Client {
	c.Fields = fields
	return c
}
