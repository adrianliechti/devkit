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

func SelectContainer(ctx context.Context, client engine.Engine, kind string, all bool) (*engine.Container, error) {
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

func MustContainer(ctx context.Context, client engine.Engine, kind string, all bool) engine.Container {
	container, err := SelectContainer(ctx, client, kind, all)

	if err != nil {
		cli.Fatal(err)
	}

	return *container
}

// func convertContainer(info *docker.ContainerInfo) *container.Container {
// 	if info == nil {
// 		return nil
// 	}

// 	container := &container.Container{
// 		Name: info.Name,

// 		//Labels: info.Labels,

// 		Image:   info.Image,
// 		Command: info.Cmd,
// 		Args:    info.Args,

// 		WorkingDir: info.Dir,

// 		Env: info.Env,
// 	}

// 	return container
// }

// func convertPullOptions(container container.Container) docker.PullOptions {
// 	options := docker.PullOptions{}

// 	if container.PlatformContext != nil {
// 		options.Platform = container.PlatformContext.Platform
// 	}

// 	return options
// }

func convertRunOptions(container engine.Container) docker.RunOptions {
	options := docker.RunOptions{
		Name:   container.Name,
		Labels: container.Labels,

		Privileged: container.Privileged,

		Dir:  container.Dir,
		User: strings.Join([]string{container.RunAsUser, container.RunAsGroup}, ":"),

		Env:     container.Env,
		Ports:   []docker.ContainerPort{},
		Volumes: []docker.ContainerMount{},
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
		options.Ports = append(options.Ports, docker.ContainerPort{
			Port:     p.Port,
			Protocol: docker.Protocol(p.Proto),

			HostIP:   p.HostIP,
			HostPort: p.HostPort,
		})
	}

	for _, v := range container.Mounts {
		if v.HostPath != "" {
			options.Volumes = append(options.Volumes, docker.ContainerMount{
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
