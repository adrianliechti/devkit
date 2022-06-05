package postgres

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/container"
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
	return "postgres"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "PostgreSQL Database Server"
}

func (m *Manager) Description() string {
	return "PostgreSQL is a powerful, open source object-relational database system with over 30 years of active development that has earned it a strong reputation for reliability, feature robustness, and performance."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (container.Container, error) {
	image := "postgres:14-bullseye"

	database := "postgres"
	username := "postgres"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"POSTGRES_DB":       database,
			"POSTGRES_USER":     username,
			"POSTGRES_PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     5432,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/var/lib/postgresql/data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	database := instance.Env["POSTGRES_DB"]
	username := instance.Env["POSTGRES_USER"]
	password := instance.Env["POSTGRES_PASSWORD"]

	return map[string]string{
		"Database": database,
		"Username": username,
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance container.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) Client(instance container.Container) (string, []string, error) {
	return DefaultShell, []string{
		"-c",
		"psql --username ${POSTGRES_USER} --dbname ${POSTGRES_DB}",
	}, nil
}