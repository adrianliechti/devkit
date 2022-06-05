package catalog

import (
	"context"
	"errors"
	"fmt"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/to"
)

const (
	KindKey = "local.devkit.kind"
)

func Container(ctx context.Context, name string) (*container.Container, error) {
	info, err := docker.Info(ctx, name)

	if err != nil {
		return nil, err
	}

	return convertContainer(info), nil
}

func SelectContainer(ctx context.Context, kind string) (*container.Container, error) {
	list, err := docker.List(ctx, docker.ListOptions{
		All: true,

		Filter: []string{
			"label=" + KindKey + "=" + kind,
		},
	})

	var items []string

	if err != nil {
		return nil, err
	}

	for _, c := range list {
		name := c.Names[0]
		items = append(items, name)
	}

	if len(items) == 0 {
		return nil, errors.New("no instances found")
	}

	i, _, err := cli.Select("select instance", items)

	if err != nil {
		return nil, err
	}

	return Container(ctx, list[i].ID)
}

func MustContainer(ctx context.Context, kind string) container.Container {
	container, err := SelectContainer(ctx, kind)

	if err != nil {
		cli.Fatal(err)
	}

	return *container
}

func convertContainer(info *docker.ContainerInfo) *container.Container {
	if info == nil {
		return nil
	}

	container := &container.Container{
		Name: info.Name,

		//Labels: info.Labels,

		Image:   info.Image,
		Command: info.Cmd,
		Args:    info.Args,

		WorkingDir: info.Dir,

		Env: info.Env,
	}

	return container
}

func convertPullOptions(container container.Container) docker.PullOptions {
	options := docker.PullOptions{}

	if container.PlatformContext != nil {
		options.Platform = container.PlatformContext.Platform
	}

	return options
}

func convertRunOptions(container container.Container) docker.RunOptions {
	options := docker.RunOptions{
		Name:   container.Name,
		Labels: container.Labels,

		Dir: container.WorkingDir,

		Env: container.Env,
	}

	if container.PlatformContext != nil {
		options.Platform = container.PlatformContext.Platform

		if container.PlatformContext.MaxNoProcs != nil {
			options.MaxNoProcs = *container.PlatformContext.MaxNoProcs
		}

		if container.PlatformContext.MaxNoFiles != nil {
			options.MaxNoFiles = *container.PlatformContext.MaxNoFiles
		}
	}

	if container.SecurityContext != nil {
		options.Privileged = to.Bool(container.SecurityContext.Privileged)

		if container.SecurityContext.RunAsUser != nil {
			user := fmt.Sprintf("%d", *container.SecurityContext.RunAsUser)

			if container.SecurityContext.RunAsGroup != nil {
				user += fmt.Sprintf(":%d", *container.SecurityContext.RunAsGroup)
			}

			options.User = user
		}
	}

	ports := map[int]int{}

	for _, p := range container.Ports {
		if p.HostPort == nil {
			println("no host port")
			continue
		}

		hostPort := *p.HostPort

		ports[hostPort] = p.Port
	}

	options.Ports = ports

	volumes := map[string]string{}

	for _, v := range container.VolumeMounts {
		if v.HostPath == "" {
			continue
		}

		volumes[v.HostPath] = v.Path
	}

	options.Volumes = volumes

	return options
}

func printMapTable(v map[string]string) {
	rows := [][]string{}

	for k, v := range v {
		rows = append(rows, []string{k, v})
	}

	cli.Table([]string{"Key", "Value"}, rows)
}
