package frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/spotifydown/api"
	"github.com/spudtrooper/spotifydown/handlers"
)

func ListenAndServe(ctx context.Context, client *api.Client, port int, host string) error {
	mux := http.NewServeMux()
	handler.Init(mux)
	if err := handler.AddHandlers(ctx, mux, handlers.CreateHandlers(client),
		handler.AddHandlersPrefix("api"),
		handler.AddHandlersKey("spotofydown"),
		handler.AddHandlersIndexTitle("unofficial spotofydown API"),
		handler.AddHandlersFooterHTML(`Details: <a target="_" href="//github.com/spudtrooper/spotofydown">github.com/spudtrooper/spotofydown</a>`),
		handler.AddHandlersSourceLinks(true),
		handler.AddHandlersSerializedSourceLocations(handlers.SourceLocations),
	); err != nil {
		return err
	}
	mux.Handle("/", http.RedirectHandler("/api", http.StatusSeeOther))

	log.Printf("listening on http://%s:%d", host, port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return err
	}

	return nil
}
