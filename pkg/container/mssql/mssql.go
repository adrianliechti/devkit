package mssql

import (
	"runtime"

	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
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

		Ports: []container.ContainerPort{
			{
				Port:     1433,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []container.VolumeMount{
			{
				Path: "/var/opt/mssql",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	username := "sa"
	password := container.Env["SA_PASSWORD"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}
}

func ClientCmd() (string, []string) {
	return DefaultShell, []string{
		"-c",
		"/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P ${SA_PASSWORD}",
	}
}
