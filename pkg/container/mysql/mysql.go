package mysql

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "docker.io/library/mysql:8-oracle"

	database := "db"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"MYSQL_DATABASE":      database,
			"MYSQL_ROOT_PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     3306,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/var/lib/mysql",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	database := container.Env["MYSQL_DATABASE"]
	username := "root"
	password := container.Env["MYSQL_ROOT_PASSWORD"]

	return map[string]string{
		"Database": database,
		"Username": username,
		"Password": password,
	}
}

func ClientCmd() (string, []string) {
	return DefaultShell, []string{
		"-c",
		"mysql --user=root --password=${MYSQL_ROOT_PASSWORD} ${MYSQL_DATABASE}",
	}
}
