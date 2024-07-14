package engine

import (
	"context"
	"io"
	"net"
)

type Client interface {
	List(ctx context.Context, options ListOptions) ([]Container, error)

	Pull(ctx context.Context, image, platform string, options PullOptions) error

	Create(ctx context.Context, spec Container, options CreateOptions) (string, error)
	Delete(ctx context.Context, container string, options DeleteOptions) error

	Inspect(ctx context.Context, container string) (Container, error)

	Logs(ctx context.Context, container string, options LogsOptions) error
}

type ListOptions struct {
	All bool

	LabelSelector map[string]string
}

type PullOptions struct {
	Stdout io.Writer
	Stderr io.Writer
}

type CreateOptions struct {
	Stdout io.Writer
	Stderr io.Writer
}

type DeleteOptions struct {
}

type LogsOptions struct {
	Follow bool

	Stdout io.Writer
	Stderr io.Writer
}

type Container struct {
	ID   string
	Name string

	Labels map[string]string

	Image    string
	Platform string

	Privileged bool

	RunAsUser  string
	RunAsGroup string

	MaxFiles     int64
	MaxProcesses int64

	Env map[string]string
	Dir string

	Command []string
	Args    []string

	Hostname  string
	IPAddress net.IP

	Ports  []*ContainerPort
	Mounts []*ContainerMount
}

type Protocol string

const (
	ProtocolTCP Protocol = "tcp"
	ProtocolUDP Protocol = "udp"
)

type ContainerPort struct {
	Name string

	Port  int
	Proto Protocol

	HostIP   string
	HostPort *int
}

type ContainerMount struct {
	Path string

	Volume   string
	HostPath string
}
