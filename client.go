package pingidentityapi

import (
	"net/http"
	"encoding/json"
	"gopkg.in/resty.v1"
)

type Client struct {
	baseURL 		string
	username 		string
	password 		string
	*resty.Client
}

type Configuration struct {
	baseURL			string
	username		string
	password		string 
	transport		http.RoundTripper
}


func NewClient(config *Configuration) *Client {
	client := resty.New()
	if config.transport != nil {
		client.SetTransport(config.transport)
	}
	client.SetHeader("X-Xsrf-Header", "PingAccess")
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Content-Type", "application/json")
	client.SetBasicAuth(config.username, config.password)
	client.SetRESTMode()
	return &Client {
		baseURL: config.baseURL,
		Client: client,
	}
}

func (c *Client) Get(path string) (map[string]interface{}, error) {
	resp, err := c.R().Get(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	json.Unmarshal(resp.Body(), &m)
	return m, err
}

func (c *Client) Post(path string, body map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.R().SetBody(body).Post(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	json.Unmarshal(resp.Body(), &m)
	return m, err
}

func (c *Client) Put(path string, body map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.R().SetBody(body).Put(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	json.Unmarshal(resp.Body(), &m)
	return m, err
}

func (c *Client) Delete(path string) (map[string]interface{}, error) {
	resp, err := c.R().Delete(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	json.Unmarshal(resp.Body(), &m)
	return m, err
}