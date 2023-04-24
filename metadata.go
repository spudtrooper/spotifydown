package main

import (
	"flag"
	"log"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/json"
	"github.com/spudtrooper/goutil/request"
)

func main() {
	flag.Parse()
	// Options
	printData := true
	printCookies := true
	printPayload := true

	// Data
	uri := request.MakeURL("https://api.spotifydown.com/metadata/track/50M7nY1oQuNHecs0ahWAtI")
	cookie := [][2]string{}
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

		"cookie": request.CreateCookie(cookie),
	}

	// Make the request
	var payload interface{}
	var res *request.Response
	var err error

	res, err = request.Get(uri, &payload, request.RequestExtraHeaders(headers))

	if printData {
		log.Printf("data: %s", string(res.Data))
	}
	if printCookies {
		log.Printf("cookies: %v", res.Cookies)
	}
	if printPayload {
		log.Printf("payload: %s", json.MustColorMarshal(payload))
	}
	check.Err(err)
}
