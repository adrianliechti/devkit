package elasticsearch

import (
	"github.com/adrianliechti/devkit/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"

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

func (m *Manager) New() (engine.Container, error) {
	image := "elasticsearch:8.14.3"

	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		MaxFiles: 65535,

		Env: map[string]string{
			"node.name": "es",

			"cluster.name":   "default",
			"discovery.type": "single-node",

			"xpack.security.enabled": "true",

			"ELASTIC_PASSWORD": password,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  9200,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/usr/share/elasticsearch/data",
			},
		},
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	username := "elastic"
	password := instance.Env["ELASTIC_PASSWORD"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}, nil
}
