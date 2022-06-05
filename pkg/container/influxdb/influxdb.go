package influxdb

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "influxdb:2.2"

	org := "default"
	bucket := "default"

	token := password.MustGenerate(10, 4, 0, false, false)

	username := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"DOCKER_INFLUXDB_INIT_MODE": "setup",

			"DOCKER_INFLUXDB_INIT_ORG":    org,
			"DOCKER_INFLUXDB_INIT_BUCKET": bucket,

			"DOCKER_INFLUXDB_INIT_USERNAME": username,
			"DOCKER_INFLUXDB_INIT_PASSWORD": password,

			"DOCKER_INFLUXDB_INIT_ADMIN_TOKEN": token,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     8086,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/var/lib/influxdb2",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	org := container.Env["DOCKER_INFLUXDB_INIT_ORG"]
	bucket := container.Env["DOCKER_INFLUXDB_INIT_BUCKET"]

	token := container.Env["DOCKER_INFLUXDB_INIT_ADMIN_TOKEN"]

	username := container.Env["DOCKER_INFLUXDB_INIT_USERNAME"]
	password := container.Env["DOCKER_INFLUXDB_INIT_PASSWORD"]

	return map[string]string{
		"Org":    org,
		"Bucket": bucket,

		"Token": token,

		"Username": username,
		"Password": password,
	}
}

func ConsolePort() container.ContainerPort {
	return container.ContainerPort{
		Port:     8086,
		Protocol: container.ProtocolTCP,
	}
}
