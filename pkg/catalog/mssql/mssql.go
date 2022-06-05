package mssql

import (
	"runtime"

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

func (m *Manager) New() (container.Container, error) {
	image := "mcr.microsoft.com/mssql/server:2019-latest"

	if runtime.GOARCH == "arm64" {
		image = "mcr.microsoft.com/azure-sql-edge"
	}

	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"ACCEPT_EULA": "Y",
			"SA_PASSWORD": password,

			"MSSQL_PID":         "Developer",
			"MSSQL_SA_PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     1433,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/var/opt/mssql",
			},
		},
	}, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	username := "sa"
	password := instance.Env["SA_PASSWORD"]

	return map[string]string{
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
		"/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P ${SA_PASSWORD}",
	}, nil
}
