package mosquitto

import (
	"fmt"

	"github.com/adrianliechti/devkit/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var (
	_ catalog.Manager       = &Manager{}
	_ catalog.Decorator     = &Manager{}
	_ catalog.ShellProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "mosquitto"
}

func (m *Manager) Category() catalog.Category {
	return catalog.MessagingCategory
}

func (m *Manager) DisplayName() string {
	return "Mosquitto"
}

func (m *Manager) Description() string {
	return "Eclipse Mosquitto is an open source message broker that implements the MQTT protocol versions 5.0, 3.1.1 and 3.1"
}

const (
	DefaultShell = "/bin/ash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "eclipse-mosquitto:2"

	return engine.Container{
		Image: image,

		Ports: []*engine.ContainerPort{
			{
				Port:  1883,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/mosquitto/config",
			},
			{
				Path: "/mosquitto/data",
			},
			{
				Path: "/mosquitto/log",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	var uri string

	for _, p := range instance.Ports {
		if p.HostPort == nil || p.Port != 1883 {
			continue
		}

		uri = fmt.Sprintf("tcp://localhost:%d", *p.HostPort)
	}

	return map[string]string{
		"URI": uri,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}
