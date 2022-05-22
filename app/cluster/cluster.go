package cluster

import (
	"context"
	"errors"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/kind"
)

const (
	Kubernetes = "cluster"
	Category   = "CLUSTER"
)

var Command = &cli.Command{
	Name:  Kubernetes,
	Usage: "local Kubernetes cluster",

	Category: Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		ListCommand(),

		CreateCommand(),
		DeleteCommand(),

		DashboardCommand(),
		GrafanaCommand(),
	},
}

func SelectCluster(ctx context.Context) (string, error) {
	list, err := kind.List(ctx)

	var items []string

	if err != nil {
		return "", err
	}

	for _, c := range list {
		items = append(items, c)
	}

	if len(items) == 0 {
		return "", errors.New("no instances found")
	}

	i, _, err := cli.Select("Select instance", items)

	if err != nil {
		return "", err
	}

	return list[i], nil
}

func MustCluster(ctx context.Context) string {
	cluster, err := SelectCluster(ctx)

	if err != nil {
		cli.Fatal(err)
	}

	return cluster
}
