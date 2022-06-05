package sonarqube

import (
	"runtime"

	"github.com/adrianliechti/devkit/pkg/container"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "sonarqube:9-community"

	if runtime.GOARCH == "arm64" {
		image = "mwizner/sonarqube:9.4.0-community"
	}

	// MaxNoProcs: 8192,
	// MaxNoFiles: 131072,

	return container.Container{
		Image: image,

		Env: map[string]string{
			"SONAR_ES_BOOTSTRAP_CHECKS_DISABLE": "true",
			"SONAR_SEARCH_JAVAADDITIONALOPTS":   "-Dbootstrap.system_call_filter=false",
		},

		Ports: []*container.ContainerPort{
			{
				Port:     9000,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/opt/sonarqube/data",
			},
			{
				Path: "/opt/sonarqube/logs",
			},
			{
				Path: "/opt/sonarqube/extensions",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	username := "admin"
	password := "admin"

	return map[string]string{
		"Username": username,
		"Password": password,
	}
}
