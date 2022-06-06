package catalog

import (
	"context"
	"errors"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/engine"
)

const (
	KindKey = "local.devkit.kind"
)

func SelectContainer(ctx context.Context, client engine.Client, kind string, all bool) (*engine.Container, error) {
	containers, err := client.List(ctx, engine.ListOptions{
		All: all,

		LabelSelector: map[string]string{
			KindKey: kind,
		},
	})

	if err != nil {
		return nil, err
	}

	var items []string

	if err != nil {
		return nil, err
	}

	for _, i := range containers {
		items = append(items, i.Name)
	}

	if len(items) == 0 {
		return nil, errors.New("no container(s) found")
	}

	i, _, err := cli.Select("select container", items)

	if err != nil {
		return nil, err
	}

	container := containers[i]
	return &container, nil
}

func MustContainer(ctx context.Context, client engine.Client, kind string, all bool) engine.Container {
	container, err := SelectContainer(ctx, client, kind, all)

	if err != nil {
		cli.Fatal(err)
	}

	return *container
}

func printMapTable(v map[string]string) {
	rows := [][]string{}

	for k, v := range v {
		rows = append(rows, []string{k, v})
	}

	cli.Table([]string{"Key", "Value"}, rows)
}
