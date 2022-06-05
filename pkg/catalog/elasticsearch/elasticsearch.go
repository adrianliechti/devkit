package elasticsearch

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/adrianliechti/devkit/pkg/to"
	"github.com/sethvargo/go-password/password"
)

var (
	_ catalog.Manager       = &Manager{}
	_ catalog.Decorator     = &Manager{}
	_ catalog.ShellProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "elasticsearch"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "Elasticsearch"
}

func (m *Manager) Description() string {
	return "Elasticsearch is a search engine based on the Lucene library."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (container.Container, error) {
	image := "docker.elastic.co/elasticsearch/elasticsearch:8.2.0"

	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"node.name": "es",

			"cluster.name":   "default",
			"discovery.type": "single-node",

			"xpack.security.enabled": "true",

			"ELASTIC_PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     9200,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/usr/share/elasticsearch/data",
			},
		},

		PlatformContext: &container.PlatformContext{
			MaxNoFiles: to.IntPtr(65535),
		},
	}, nil
}

func (m *Manager) Shell(instance container.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	username := "elastic"
	password := instance.Env["ELASTIC_PASSWORD"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}, nil
}
