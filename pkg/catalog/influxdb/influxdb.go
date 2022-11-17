package influxdb

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/sethvargo/go-password/password"
)

var (
	_ catalog.Manager         = &Manager{}
	_ catalog.Decorator       = &Manager{}
	_ catalog.ShellProvider   = &Manager{}
	_ catalog.ConsoleProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "influxdb"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "InfluxDB"
}

func (m *Manager) Description() string {
	return "InfluxDB is an open-source time series database."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "influxdb:2.5"

	org := "default"
	bucket := "default"

	token := password.MustGenerate(10, 4, 0, false, false)

	username := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"DOCKER_INFLUXDB_INIT_MODE": "setup",

			"DOCKER_INFLUXDB_INIT_ORG":    org,
			"DOCKER_INFLUXDB_INIT_BUCKET": bucket,

			"DOCKER_INFLUXDB_INIT_USERNAME": username,
			"DOCKER_INFLUXDB_INIT_PASSWORD": password,

			"DOCKER_INFLUXDB_INIT_ADMIN_TOKEN": token,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  8086,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/var/lib/influxdb2",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	org := instance.Env["DOCKER_INFLUXDB_INIT_ORG"]
	bucket := instance.Env["DOCKER_INFLUXDB_INIT_BUCKET"]

	token := instance.Env["DOCKER_INFLUXDB_INIT_ADMIN_TOKEN"]

	username := instance.Env["DOCKER_INFLUXDB_INIT_USERNAME"]
	password := instance.Env["DOCKER_INFLUXDB_INIT_PASSWORD"]

	return map[string]string{
		"Org":    org,
		"Bucket": bucket,

		"Token": token,

		"Username": username,
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  8086,
		Proto: engine.ProtocolTCP,
	}, nil
}
