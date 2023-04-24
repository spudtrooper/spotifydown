package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/spotifydown/api"
	"github.com/spudtrooper/spotifydown/frontend"
)

var (
	portForTesting = flag.Int("port_for_testing", 0, "port to listen on")
	host           = flag.String("host", "localhost", "host name")
)

func main() {
	flag.Parse()
	var port int
	if *portForTesting != 0 {
		port = *portForTesting
	} else {
		p, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatalf("invalid port: %v", err)
		}
		port = p
	}
	if port == 0 {
		log.Fatalf("port is required")
	}

	ctx := context.Background()
	client := api.NewClient()
	check.Err(frontend.ListenAndServe(ctx, client, port, *host))
}
