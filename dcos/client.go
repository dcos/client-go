package dcos

import "net/http"

type Client struct {
	HttpClient *http.Client

	baseURL string
}

func NewClient(httpClient *http.Client) (*Client, error) {
	return &Client{HttpClient: httpClient}, nil
}
