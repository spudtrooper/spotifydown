package api

import (
	"encoding/json"
	"fmt"

	goutillog "github.com/spudtrooper/goutil/log"
	"github.com/spudtrooper/goutil/request"
)

//go:generate genopts --function Base verbose

type ApiLogger interface {
	goutillog.Logger
}

// Client is a client for spotifydown.com
type Client struct {
	logger goutillog.Logger
}

// NewClientFromFlags creates a new client from command line flags
func NewClientFromFlags() (*Client, error) {
	client := NewClient()
	return client, nil
}

//go:generate genopts --function NewClient "logger:ApiLogger"
// NewClient creates a new client directly from the API Key
func NewClient(optss ...NewClientOption) *Client {
	opts := MakeNewClientOptions(optss...)

	return &Client{
		logger: goutillog.Must(opts.Logger()),
	}
}

func (c *Client) get(path string, payload interface{}) error {
	uri := fmt.Sprintf("https://api.spotifydown.com/%s", path)

	headers := map[string]string{
		"authority":          `api.spotifydown.com`,
		"accept":             `*/*`,
		"accept-language":    `en-US,en;q=0.9`,
		"cache-control":      `no-cache`,
		"origin":             `https://spotifydown.com`,
		"pragma":             `no-cache`,
		"referer":            `https://spotifydown.com/`,
		"sec-ch-ua":          `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`,
		"sec-ch-ua-mobile":   `?0`,
		"sec-ch-ua-platform": `"macOS"`,
		"sec-fetch-dest":     `empty`,
		"sec-fetch-mode":     `cors`,
		"sec-fetch-site":     `same-site`,
		"user-agent":         `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36`,
	}

	res, err := request.Get(uri, nil, request.RequestExtraHeaders(headers))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(res.Data, payload); err != nil {
		return err
	}
	return nil
}
