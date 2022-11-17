package redis

import (
	"fmt"

	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/sethvargo/go-password/password"
)

var (
	_ catalog.Manager        = &Manager{}
	_ catalog.Decorator      = &Manager{}
	_ catalog.ShellProvider  = &Manager{}
	_ catalog.ClientProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "redis"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "Redis Cache"
}

func (m *Manager) Description() string {
	return "Redis is an open source, in-memory data structure store, used as a database, cache, and message broker."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "redis:7-bullseye"

	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"REDIS_PASSWORD": password,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  6379,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	password := instance.Env["REDIS_PASSWORD"]

	var uri string

	for _, p := range instance.Ports {
		if p.HostPort == nil || p.Port != 6379 {
			continue
		}

		uri = fmt.Sprintf("redis://:%s@localhost:%d", password, *p.HostPort)
	}

	return map[string]string{
		"Password": password,
		"URI":      uri,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) Client(instance engine.Container) (string, []string, error) {
	return "", []string{
		DefaultShell,
		"-c",
		"redis-cli",
	}, nil
}
