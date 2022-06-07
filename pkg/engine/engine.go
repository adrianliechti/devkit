package engine

import (
	"context"
	"io"
	"net"
)

type Client interface {
	List(ctx context.Context, options ListOptions) ([]Container, error)

	Pull(ctx context.Context, image string, options PullOptions) error

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
	Platform string

	Stdout io.Writer
	Stderr io.Writer
}

type CreateOptions struct {
	Platform string

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

	Image string

	Privileged bool
<<<<<<< HEAD
	RunAsUser  string
	RunAsGroup string

=======

	RunAsUser  string
	RunAsGroup string

	MaxFiles     int64
	MaxProcesses int64

>>>>>>> d02e966ce156371bbaffbed0fbfcaf8cd9a711bd
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
