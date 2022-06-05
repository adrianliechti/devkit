package postgres

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
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

		Ports: []container.ContainerPort{
			{
				Port:     5432,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []container.VolumeMount{
			{
				Path: "/var/lib/postgresql/data",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	database := container.Env["POSTGRES_DB"]
	username := container.Env["POSTGRES_USER"]
	password := container.Env["POSTGRES_PASSWORD"]

	return map[string]string{
		"Database": database,
		"Username": username,
		"Password": password,
	}
}

func ClientCmd() (string, []string) {
	return DefaultShell, []string{
		"-c",
		"psql --username ${POSTGRES_USER} --dbname ${POSTGRES_DB}",
	}
}
