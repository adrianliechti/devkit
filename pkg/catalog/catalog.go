package catalog

import (
	"github.com/adrianliechti/devkit/pkg/container"
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

	New() (container.Container, error)
	Info(container.Container) (map[string]string, error)
}

type Decorator interface {
	Manager
	DisplayName() string
	Description() string
}

type ShellProvider interface {
	Manager
	Shell(container.Container) (string, error)
}

type ClientProvider interface {
	Manager
	Client(container.Container) (string, []string, error)
}

type ConsoleProvider interface {
	Manager
	ConsolePort(container.Container) (*container.ContainerPort, error)
}
