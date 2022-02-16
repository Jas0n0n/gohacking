package api

const BaseURL = "https://127.0.0.1:3443/api/v1/me"

type Client struct {
	apikey string
}

func New(apikey string) *Client  {
	return &Client{apikey: apikey}
}
