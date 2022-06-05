package db2

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
	return "db2"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "IBM DB2"
}

func (m *Manager) Description() string {
	return "Db2 is a family of data management products, including database servers, developed by IBM."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (container.Container, error) {
	image := "ibmcom/db2:11.5.7.0a"

	database := "db"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"LICENSE": "accept",

			"DBNAME": database,

			"DB2INST1_PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     50000,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/database",
			},
		},
	}, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	databaseInstance := "db2inst1"
	database := instance.Env["DBNAME"]
	password := instance.Env["DB2INST1_PASSWORD"]

	return map[string]string{
		"Instance": databaseInstance,
		"Database": database,
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance container.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) Client(instance container.Container) (string, []string, error) {
	return DefaultShell, []string{
		"-c",
		"su - db2inst1",
	}, nil
}
