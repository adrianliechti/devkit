package postgres

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

func (m *Manager) New() (engine.Container, error) {
	image := "postgres:14-bullseye"

	database := "db"
	username := "postgres"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"POSTGRES_DB":       database,
			"POSTGRES_USER":     username,
			"POSTGRES_PASSWORD": password,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  5432,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/var/lib/postgresql/data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	database := instance.Env["POSTGRES_DB"]
	username := instance.Env["POSTGRES_USER"]
	password := instance.Env["POSTGRES_PASSWORD"]

	var uri string

	for _, p := range instance.Ports {
		if p.HostPort == nil || p.Port != 5432 {
			continue
		}

		uri = fmt.Sprintf("postgresql://%s:%s@localhost:%d/%s?sslmode=disable", username, password, *p.HostPort, database)
	}

	return map[string]string{
		"Database": database,
		"Username": username,
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
		"psql --username ${POSTGRES_USER} --dbname ${POSTGRES_DB}",
	}, nil
}
