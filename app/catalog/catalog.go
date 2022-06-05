package catalog

import (
	"context"
	"errors"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

const (
	KindKey = "local.devkit.kind"
)

func SelectContainer(ctx context.Context, kind string) (string, error) {
	list, err := docker.List(ctx, docker.ListOptions{
		All: true,

		Filter: []string{
			"label=" + KindKey + "=" + kind,
		},
	})

	var items []string

	if err != nil {
		return "", err
	}

	for _, c := range list {
		name := c.Names[0]
		items = append(items, name)
	}

	if len(items) == 0 {
		return "", errors.New("no instances found")
	}

	i, _, err := cli.Select("Select instance", items)

	if err != nil {
		return "", err
	}

	return list[i].ID, nil
}

func MustContainer(ctx context.Context, kind string) string {
	container, err := SelectContainer(ctx, kind)

	if err != nil {
		cli.Fatal(err)
	}

	return container
}
