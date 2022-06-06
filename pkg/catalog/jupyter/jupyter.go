package jupyter

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
	return "jupyter"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "Jupyter"
}

func (m *Manager) Description() string {
	return "Jupyter is a project and community whose goal is to develop open-source software, open-standards, and services for interactive computing across dozens of programming languages."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "jupyter/datascience-notebook"

	token := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"RESTARTABLE": "yes",

			"JUPYTER_TOKEN":      token,
			"JUPYTER_ENABLE_LAB": "yes",
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  8888,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/home/jovyan/work",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	token := instance.Env["JUPYTER_TOKEN"]

	return map[string]string{
		"Token": token,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  8888,
		Proto: engine.ProtocolTCP,
	}, nil
}
