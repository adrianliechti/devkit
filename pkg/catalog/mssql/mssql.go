package mssql

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
	return "mssql"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "Microsoft SQL Server"
}

func (m *Manager) Description() string {
	return "Microsoft SQL Server is a relational database management system developed by Microsoft."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "mcr.microsoft.com/azure-sql-edge:1.0.6"

	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Privileged: true,

		Env: map[string]string{
			"ACCEPT_EULA": "Y",
			"SA_PASSWORD": password,

			"MSSQL_PID":         "Developer",
			"MSSQL_SA_PASSWORD": password,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  1433,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/var/opt/mssql",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	username := "sa"
	password := instance.Env["SA_PASSWORD"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) Client(instance engine.Container) (string, []string, error) {
	image := "mcr.microsoft.com/mssql-tools:latest"

	username := "sa"
	password := instance.Env["SA_PASSWORD"]

	return image, []string{
		DefaultShell,
		"-c",
		fmt.Sprintf("/opt/mssql-tools/bin/sqlcmd -S %s,1433 -U %s -P %s", instance.IPAddress, username, password),
	}, nil
}
