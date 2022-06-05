package container

type Container struct {
	Name string

	Labels map[string]string

	Image   string
	Command []string
	Args    []string

	WorkingDir string

	Env map[string]string

	Ports []*ContainerPort

	VolumeMounts []*VolumeMount

	PlatformContext *PlatformContext
	SecurityContext *SecurityContext
}

type ContainerPort struct {
	Name string

	Port     int
	Protocol Protocol

	HostIP   string
	HostPort *int
}

type SecurityContext struct {
	Privileged *bool
	RunAsUser  *int64
	RunAsGroup *int64
}

type PlatformContext struct {
	Platform string

	MaxNoProcs *int
	MaxNoFiles *int
}

type Protocol string

const (
	ProtocolTCP Protocol = "TCP"
	ProtocolUDP Protocol = "UDP"
)

type VolumeMount struct {
	Name string

	Path string

	HostPath string
}
