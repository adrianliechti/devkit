package mongodb

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "mongo:5-focal"

	database := "db"
	username := "root"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"MONGO_INITDB_DATABASE":      database,
			"MONGO_INITDB_ROOT_USERNAME": username,
			"MONGO_INITDB_ROOT_PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     27017,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/data/db",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	database := container.Env["MONGO_INITDB_DATABASE"]
	username := container.Env["MONGO_INITDB_ROOT_USERNAME"]
	password := container.Env["MONGO_INITDB_ROOT_PASSWORD"]

	return map[string]string{
		"Database": database,
		"Username": username,
		"Password": password,
	}
}

func ClientCmd() (string, []string) {
	return DefaultShell, []string{
		"-c",
		"mongo --quiet --norc --username ${MONGO_INITDB_ROOT_USERNAME} --password ${MONGO_INITDB_ROOT_PASSWORD}",
	}
}
