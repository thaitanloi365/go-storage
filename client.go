package storage

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

// Client instance
type Client struct {
	req          *gorequest.SuperAgent
	timeout      time.Duration
	retryTimes   int
	retryTimeout time.Duration
	endpoint     string
}

// ClientConfig config
type ClientConfig struct {
	Endpoint     string
	Timeout      time.Duration
	RetryTimes   int
	RetryTimeout time.Duration
	Debug        bool
}

// NewClient init client
func NewClient(config *ClientConfig) *Client {
	var timeout = time.Second * 30
	var retryTimes = 3
	var retryTimeout = time.Second
	if config != nil {
		if config.Timeout > 0 {
			timeout = config.Timeout
		}

		if config.RetryTimes > 0 {
			retryTimes = config.RetryTimes
		}

		if config.RetryTimeout > 0 {
			retryTimeout = config.RetryTimeout
		}
	}

	var client = &Client{
		req:          gorequest.New().SetDebug(config.Debug),
		timeout:      timeout,
		retryTimes:   retryTimes,
		retryTimeout: retryTimeout,
		endpoint:     config.Endpoint,
	}

	return client
}

func (c *Client) buildURL(path string) string {
	return fmt.Sprintf("%s/%s", c.endpoint, path)
}

// UploadParams params
type UploadParams struct {
	APIPath  string
	FileName string
	Content  bytes.Buffer
}

// Upload upload file
func (c *Client) Upload(params UploadParams) {
	resp, _, errs := c.req.
		Post(c.buildURL(params.APIPath)).
		Retry(c.retryTimes, c.retryTimeout, http.StatusBadRequest, http.StatusInternalServerError).
		Timeout(c.timeout).
		End()
	if len(errs) > 0 {

	}
	fmt.Println("resp", resp)
}
