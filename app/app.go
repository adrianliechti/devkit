package app

import (
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

func Client(c *cli.Context) (engine.Client, error) {
	return moby.New()
}

func MustClient(c *cli.Context) engine.Client {
	client, err := Client(c)

	if err != nil {
		cli.Fatal(err)
	}

	return client
}
