package cli

import (
	"context"

	"github.com/spudtrooper/minimalcli/handler"
	flag "github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/spotifydown/api"
	"github.com/spudtrooper/spotifydown/handlers"
)

func Main(ctx context.Context) error {
	flag.String("track", "", "track")
	flag.String("id", "", "id")
	flag.String("verbose", "", "verbose")
	flag.String("task_id", "", "task_id")

	client, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}

	return handler.RunCLI(ctx, handlers.CreateHandlers(client)...)
}
