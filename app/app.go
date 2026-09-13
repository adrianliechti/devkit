package app

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/devkit/pkg/engine/moby"
	"github.com/adrianliechti/go-cli"
)

func MustClient(ctx context.Context, cmd *cli.Command) engine.Client {
	client, err := moby.New()

	if err != nil {
		cli.Fatal(err)
	}

	return client
}
