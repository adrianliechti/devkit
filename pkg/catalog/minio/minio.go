package minio

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
	return "minio"
}

func (m *Manager) Category() catalog.Category {
	return catalog.StorageCategory
}

func (m *Manager) DisplayName() string {
	return "MinIO S3 Object Storage"
}

func (m *Manager) Description() string {
	return "MinIO is a High Performance Object Storage and is API compatible with Amazon S3 cloud storage service."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "minio/minio"

	username := "root"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"MINIO_ROOT_USER":     username,
			"MINIO_ROOT_PASSWORD": password,
		},

		Args: []string{
			"server", "/data", "--console-address", ":9001",
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  9000,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	username := instance.Env["MINIO_ROOT_USER"]
	password := instance.Env["MINIO_ROOT_PASSWORD"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  9001,
		Proto: engine.ProtocolTCP,
	}, nil
}
