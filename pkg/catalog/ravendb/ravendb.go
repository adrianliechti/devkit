package ravendb

import (
	"fmt"

	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var (
	_ catalog.Manager         = &Manager{}
	_ catalog.Decorator       = &Manager{}
	_ catalog.ShellProvider   = &Manager{}
	_ catalog.ClientProvider  = &Manager{}
	_ catalog.ConsoleProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "ravendb"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "RavenDB NoSQL Database"
}

func (m *Manager) Description() string {
	return "Easy, Fast, Reliable with Multi-Document ACID Transactions"
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "ravendb/ravendb:latest-lts"

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"RAVEN_Setup_Mode": "None",

			"RAVEN_License_Eula_Accepted": "true",

			"RAVEN_Security_UnsecuredAccessAllowed": "PublicNetwork",
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  8080,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/opt/RavenDB/config",
			},
			{
				Path: "/opt/RavenDB/Server/RavenData",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	var uri string

	for _, p := range instance.Ports {
		if p.HostPort == nil || p.Port != 8080 {
			continue
		}

		uri = fmt.Sprintf("http://localhost:%d/", *p.HostPort)
	}

	return map[string]string{
		"URI": uri,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) Client(instance engine.Container) (string, []string, error) {
	return "", []string{
		DefaultShell,
		"-c",
		"./rvn admin-channel",
	}, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  8080,
		Proto: engine.ProtocolTCP,
	}, nil
}
