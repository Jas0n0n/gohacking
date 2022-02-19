package shodan

const BaseURL = "https://api.shodan.io"

type Client struct {
	apikey string
}

func New(apikey string) *Client {
	return &Client{apikey: apikey}
}

