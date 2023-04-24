package api

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spudtrooper/goutil/io"
	goutillog "github.com/spudtrooper/goutil/log"
	"github.com/spudtrooper/goutil/request"
)

var (
	userCreds   = flag.String("uber_user_creds", ".user_creds.json", "file with user credentials")
	noUserCreds = flag.Bool("uber_no_user_creds", false, "Don't use user creds event if it exists")
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

// NewClientFromFile creates a new client from a JSON file like `user_creds-example.json`
func NewClientFromFile(credsFile string) (*Client, error) {
	creds, err := readCreds(credsFile)
	log.Printf("creds: %+v", creds)
	if err != nil {
		return nil, err
	}
	return &Client{}, nil
}

type Creds struct {
}

func ReadCredsFromFlags() (Creds, error) {
	return readCreds(*userCreds)
}

func WriteCredsFromFlags(creds Creds) error {
	b, err := json.Marshal(&creds)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(*userCreds, b, 0755); err != nil {
		return err
	}
	log.Printf("wrote to %s", *userCreds)
	return nil
}

func readCreds(credsFile string) (creds Creds, ret error) {
	if !io.FileExists(credsFile) {
		return
	}
	credsBytes, err := ioutil.ReadFile(credsFile)
	if err != nil {
		ret = err
		return
	}
	if err := json.Unmarshal(credsBytes, &creds); err != nil {
		ret = err
		return
	}
	return
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
