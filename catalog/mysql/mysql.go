package mysql

import (
	"github.com/adrianliechti/devkit/catalog"
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
	return "mysql"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "MySQL Database Server"
}

func (m *Manager) Description() string {
	return "MySQL is an open-source relational database management system. "
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "docker.io/library/mysql:9-oracle"

	database := "db"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"MYSQL_DATABASE":      database,
			"MYSQL_ROOT_PASSWORD": password,
		},

		Ports: []engine.ContainerPort{
			{
				Port:  3306,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []engine.ContainerMount{
			{
				Path: "/var/lib/mysql",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	database := instance.Env["MYSQL_DATABASE"]
	username := "root"
	password := instance.Env["MYSQL_ROOT_PASSWORD"]

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
		"mysql --user=root --password=${MYSQL_ROOT_PASSWORD} ${MYSQL_DATABASE}",
	}, nil
}
