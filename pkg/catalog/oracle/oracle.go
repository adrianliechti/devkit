package oracle

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

var (
	_ catalog.Manager        = &Manager{}
	_ catalog.Decorator      = &Manager{}
	_ catalog.ShellProvider  = &Manager{}
	_ catalog.ClientProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "oracle"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "Oracle Database Server"
}

func (m *Manager) Description() string {
	return "Oracle Database is a multi-model database management system produced and marketed by Oracle Corporation"
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (container.Container, error) {
	image := "gvenzl/oracle-xe:21"

	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"ORACLE_PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     1521,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/opt/oracle/oradata",
			},
		},
	}, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	password := instance.Env["ORACLE_PASSWORD"]

	return map[string]string{
		"Service":  "XEPDB1",
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance container.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) Client(instance container.Container) (string, []string, error) {
	return DefaultShell, []string{
		"-c",
		"sqlplus system/${ORACLE_PASSWORD}@XEPDB1",
	}, nil
}
