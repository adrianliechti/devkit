package nats

import (
	"fmt"

	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"

	"github.com/sethvargo/go-password/password"
)

var (
	_ catalog.Manager       = &Manager{}
	_ catalog.Decorator     = &Manager{}
	_ catalog.ShellProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "nats"
}

func (m *Manager) Category() catalog.Category {
	return catalog.MessagingCategory
}

func (m *Manager) DisplayName() string {
	return "NATS"
}

func (m *Manager) Description() string {
	return "NATS is a connective technology that powers modern distributed systems."
}

const (
	DefaultShell = "/bin/ash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "nats:2-alpine"

	username := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},

		Args: []string{
			"-js",
			"-sd", "/data",
			"--name", "default",
			"--cluster_name", "default",
			"--user", username,
			"--pass", password,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  4222,
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
	username := instance.Env["USERNAME"]
	password := instance.Env["PASSWORD"]

	var uri string

	for _, p := range instance.Ports {
		if p.HostPort == nil || p.Port != 4222 {
			continue
		}

		uri = fmt.Sprintf("nats://%s:%s@localhost:%d", username, password, *p.HostPort)
	}

	return map[string]string{
		"Username": username,
		"Password": password,
		"URI":      uri,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}
