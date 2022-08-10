package catalog

import (
	"github.com/adrianliechti/devkit/pkg/engine"
)

type Category string

const (
	DatabaseCategory  Category = "Database"
	MessagingCategory Category = "Messaging"
	PlatformCategory  Category = "Platform"
	StorageCategory   Category = "Storage"
)

type Manager interface {
	Name() string
	Category() Category

	New() (engine.Container, error)
	Info(engine.Container) (map[string]string, error)
}

type Decorator interface {
	Manager
	DisplayName() string
	Description() string
}

type ShellProvider interface {
	Manager
	Shell(engine.Container) (string, error)
}

type ClientProvider interface {
	Manager
	Client(engine.Container) (string, []string, error)
}

type ClientContainerProvider interface {
	Manager
	Client(engine.Container) (string, string, []string, error)
}

type ConsoleProvider interface {
	Manager
	ConsolePort(engine.Container) (*engine.ContainerPort, error)
}
