package db2

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
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
	}
}

func Info(container *container.Container) map[string]string {
	database := container.Env["DBNAME"]
	instance := "db2inst1"
	password := container.Env["DB2INST1_PASSWORD"]

	return map[string]string{
		"Instance": instance,
		"Database": database,
		"Password": password,
	}
}

func ClientCmd() (string, []string) {
	return DefaultShell, []string{
		"-c",
		"su - db2inst1",
	}
}
