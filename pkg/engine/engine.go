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

	Run(ctx context.Context, spec Container, options RunOptions) error
	Exec(ctx context.Context, containerID string, command []string, options ExecOptions) error

	PortForward(ctx context.Context, containerID, address string, ports map[int]int, readyChan chan struct{}) error
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
}

type RunOptions struct {
	Stdin  io.Reader
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

type ExecOptions struct {
	Privileged bool

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	Dir  string
	User string

	Env map[string]string
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

	Ports  []ContainerPort
	Mounts []ContainerMount
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
	HostPort int
}

type ContainerMount struct {
	Path string

	Volume   string
	HostPath string
}
