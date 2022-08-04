package anka_api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	controllerAddress *url.URL
	apiKey            string
	httpClient        *http.Client
}

const StatusPath = "/api/v1/status"
const VMPath = "/api/v1/vm"

func NewClient(addr string, certs TLSCerts, apiKey string) (*Client, error) {
	if err := setUpTLS(certs); err != nil {
		return nil, err
	}
	uri, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	return &Client{
		controllerAddress: uri,
		apiKey:            apiKey,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}, nil
}

func (c *Client) newRequest(method, path string, query string, body interface{}) (*http.Request, error) {
	c.controllerAddress.Path = path
	c.controllerAddress.RawQuery = query
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, c.controllerAddress.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c Client) Status() (*StatusBodyResponse, error) {
	req, err := c.newRequest("GET", StatusPath, "", nil)
	if err != nil {
		return nil, err
	}
	var status StatusBodyResponse
	_, err = c.do(req, &status)
	if err != nil {
		return nil, err
	}
	return &status, err
}

func (c Client) StartVM() (error, *DefaultResponse) {
	return nil, &DefaultResponse{}
}

type Metadata map[string]interface{}

type StartVMOptions struct {
	Tag                    string
	Version                int
	Name                   string
	ExternalId             string
	count                  int
	NodeId                 string
	StartupScript          string
	StartupScriptCondition int
	ScriptMonitoring       bool
	ScriptTimeout          int
	ScriptFailHandler      int
	NodeTemplate           string
	GroupId                string
	Priority               int
	USBDevice              string
	VCPU                   int
	VRAM                   int
	Metadata               Metadata
	MacAddress             string
	VlanTag                string
}

type UpdateVMOptions struct {
	Name       string
	ExternalId string
	Metadata   Metadata
}
