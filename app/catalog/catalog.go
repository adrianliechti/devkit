package catalog

import (
	"context"
	"errors"
	"fmt"

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

func printContainerInfo(container engine.Container, info map[string]string) {
	rowsPorts := [][]string{}

	for _, p := range container.Ports {
		if container.IPAddress == nil || p.HostPort == nil {
			continue
		}

		rowsPorts = append(rowsPorts, []string{fmt.Sprintf("localhost:%d", *p.HostPort), fmt.Sprintf("%s://%s:%d", p.Proto, container.IPAddress, p.Port)})
	}

	if len(rowsPorts) > 0 {
		cli.Table([]string{"Mapping", "Target"}, rowsPorts)
	}

	rowsInfo := [][]string{}

	rowsInfo = append(rowsInfo, []string{"Name", container.Name})

	for k, v := range info {
		rowsInfo = append(rowsInfo, []string{k, v})
	}

	if len(rowsInfo) > 0 {
		if len(rowsPorts) > 0 {
			cli.Info()
		}

		cli.Table([]string{"Description", "Value"}, rowsInfo)
	}
}
