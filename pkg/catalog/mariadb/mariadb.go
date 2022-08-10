package mariadb

import (
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
	return "mariadb"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "MariaDB Database Server"
}

func (m *Manager) Description() string {
	return "MariaDB is a community-developed, commercially supported fork of the MySQL relational database management system, intended to remain free and open-source software."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "mariadb:10-focal"

	database := "db"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"MARIADB_DATABASE":      database,
			"MARIADB_ROOT_PASSWORD": password,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  3306,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/var/lib/mysql",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	database := instance.Env["MARIADB_DATABASE"]
	username := "root"
	password := instance.Env["MARIADB_ROOT_PASSWORD"]

	return map[string]string{
		"Database": database,
		"Username": username,
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) Client(instance engine.Container) (string, []string, error) {
	return "", []string{
		DefaultShell,
		"-c",
		"mysql --user=root --password=${MARIADB_ROOT_PASSWORD} ${MARIADB_DATABASE}",
	}, nil
}
