// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"
	_ "embed"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/spotifydown/api"
)

//go:generate minimalcli gsl --input handlers.go --uri_root "github.com/spudtrooper/spotifydown/blob/main/handlers" --output handlers.go.json
//go:embed handlers.go.json
var SourceLocations []byte

func CreateHandlers(client *api.Client) []handler.Handler {
	b := handler.NewHandlerBuilder()

	b.NewHandler("Metadata",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.MetadataParams)
			return client.Metadata(p.Options()...)
		},
		api.MetadataParams{},
		handler.NewHandlerExtraRequiredFields([]string{"track"}),
	)

	b.NewHandler("GetID",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetIDParams)
			return client.GetID(p.Options()...)
		},
		api.GetIDParams{},
		handler.NewHandlerExtraRequiredFields([]string{"track"}),
	)

	b.NewHandler("Download",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.DownloadParams)
			return client.Download(p.Options()...)
		},
		api.DownloadParams{},
		handler.NewHandlerExtraRequiredFields([]string{"id"}),
	)

	b.NewHandler("Progress",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.ProgressParams)
			return client.Progress(p.Options()...)
		},
		api.ProgressParams{},
		handler.NewHandlerExtraRequiredFields([]string{"task_id"}),
	)

	b.NewHandler("Convert",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.ConvertParams)
			return client.Convert(p.Options()...)
		},
		api.ConvertParams{},
		handler.NewHandlerExtraRequiredFields([]string{"track"}),
	)

	return b.Build()
}
