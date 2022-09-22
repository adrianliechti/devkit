package ravendb

import (
	"fmt"

	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var (
	_ catalog.Manager       = &Manager{}
	_ catalog.Decorator     = &Manager{}
	_ catalog.ShellProvider = &Manager{}
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
	// database := instance.Env["POSTGRES_DB"]
	// username := instance.Env["POSTGRES_USER"]
	// password := instance.Env["POSTGRES_PASSWORD"]

	var uri string

	for _, p := range instance.Ports {
		if p.HostPort == nil || p.Port != 8080 {
			continue
		}

		//uri = fmt.Sprintf("postgresql://%s:%s@localhost:%d/%s?sslmode=disable", username, password, *p.HostPort, database)
		uri = fmt.Sprintf("http://localhost:%d/", *p.HostPort)
	}

	return map[string]string{
		// 	"Database": database,
		// 	"Username": username,
		// 	"Password": password,
		"URI": uri,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}
