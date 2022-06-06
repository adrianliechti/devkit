package catalog

import (
	"context"
	"errors"
	"strings"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
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

func convertRunOptions(container engine.Container) docker.RunOptions {
	options := docker.RunOptions{
		Name:   container.Name,
		Labels: container.Labels,

		Privileged: container.Privileged,

		Dir:  container.Dir,
		User: strings.Join([]string{container.RunAsUser, container.RunAsGroup}, ":"),

		Env:     container.Env,
		Ports:   []engine.ContainerPort{},
		Volumes: []engine.ContainerMount{},
	}

	// TODO
	// if container.PlatformContext != nil {
	// 	options.Platform = container.PlatformContext.Platform

	// 	if container.PlatformContext.MaxNoProcs != nil {
	// 		options.MaxNoProcs = *container.PlatformContext.MaxNoProcs
	// 	}

	// 	if container.PlatformContext.MaxNoFiles != nil {
	// 		options.MaxNoFiles = *container.PlatformContext.MaxNoFiles
	// 	}
	// }

	for _, p := range container.Ports {
		options.Ports = append(options.Ports, engine.ContainerPort{
			Port:  p.Port,
			Proto: p.Proto,

			HostIP:   p.HostIP,
			HostPort: p.HostPort,
		})
	}

	for _, v := range container.Mounts {
		if v.Volume != "" {
			options.Volumes = append(options.Volumes, engine.ContainerMount{
				Path:   v.Path,
				Volume: v.Volume,
			})
		}

		if v.HostPath != "" {
			options.Volumes = append(options.Volumes, engine.ContainerMount{
				Path:     v.Path,
				HostPath: v.HostPath,
			})
		}
	}

	return options
}

func printMapTable(v map[string]string) {
	rows := [][]string{}

	for k, v := range v {
		rows = append(rows, []string{k, v})
	}

	cli.Table([]string{"Key", "Value"}, rows)
}
