package jupyter

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "jupyter/datascience-notebook"

	token := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"RESTARTABLE": "yes",

			"JUPYTER_TOKEN":      token,
			"JUPYTER_ENABLE_LAB": "yes",
		},

		Ports: []*container.ContainerPort{
			{
				Port:     8888,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/home/jovyan/work",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	token := container.Env["JUPYTER_TOKEN"]

	return map[string]string{
		"Token": token,
	}
}
