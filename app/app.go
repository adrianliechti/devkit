package app

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/devkit/pkg/engine/moby"
)

const (
	DatabaseCategory  = "DATABASE"
	MessagingCategory = "MESSAGING"
	PlatformCategory  = "PLATFORM"
	StorageCategory   = "STORAGE"
	UtilityCategory   = "UTILILITY"
	TemplateCategory  = "TEMPLATE"
)

func Client(ctx context.Context, cmd *cli.Command) (engine.Client, error) {
	return moby.New()
}

func MustClient(ctx context.Context, cmd *cli.Command) engine.Client {
	client, err := Client(ctx, cmd)

	if err != nil {
		cli.Fatal(err)
	}

	return client
}
