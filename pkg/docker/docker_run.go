package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"

	"github.com/adrianliechti/devkit/pkg/engine"
)

type RunOptions struct {
	Name string

	Platform string

	Temporary  bool
	Privileged bool

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	Attach      bool
	TTY         bool
	Interactive bool

	Dir  string
	User string

	Env     map[string]string
	Ports   []engine.ContainerPort
	Volumes []engine.ContainerMount
}

func Run(ctx context.Context, image string, options RunOptions, args ...string) error {
	tool, _, err := Info(ctx)

	if err != nil {
		return err
	}

	run := exec.CommandContext(ctx, tool, runArgs(image, options, args...)...)
	run.Stdin = options.Stdin
	run.Stdout = options.Stdout
	run.Stderr = options.Stderr

	return run.Run()
}

func RunInteractive(ctx context.Context, image string, options RunOptions, args ...string) error {
	if options.Stdin == nil {
		options.Stdin = os.Stdin
	}

	if options.Stdout == nil {
		options.Stdout = os.Stdout
	}

	if options.Stderr == nil {
		options.Stderr = os.Stderr
	}

	options.Temporary = true

	options.TTY = true
	options.Attach = true
	options.Interactive = true

	return Run(ctx, image, options, args...)
}

func runArgs(image string, options RunOptions, arg ...string) []string {
	args := []string{
		"run",
	}

	if options.Name != "" {
		args = append(args, "--name", options.Name)
	}

	if options.User != "" {
		args = append(args, "--user", options.User)
	}

	if options.Platform != "" {
		args = append(args, "--platform", options.Platform)
	}

	if options.Temporary {
		args = append(args, "--rm")
	}

	if options.Privileged {
		args = append(args, "--privileged")
	}

	if !options.Attach {
		args = append(args, "--detach")
	}

	if options.Interactive {
		args = append(args, "--interactive")
	}

	if options.TTY {
		args = append(args, "--tty")
	}

	if options.Dir != "" {
		args = append(args, "--workdir", options.Dir)
	}

	for key, value := range options.Env {
		args = append(args, "--env", key+"="+value)
	}

	for _, p := range options.Ports {
		port := strconv.Itoa(p.Port)
		proto := p.Proto

		if proto == "" {
			proto = engine.ProtocolTCP
		}

		hostIP := "127.0.0.1"
		hostPort := ""

		if p.HostIP != "" {
			hostIP = p.HostIP
		}

		if p.HostPort != nil {
			hostPort = strconv.Itoa(*p.HostPort)
		}

		args = append(args, "--publish", fmt.Sprintf("%s:%s:%s/%s", hostIP, hostPort, port, proto))
	}

	for _, v := range options.Volumes {
		if v.Volume != "" {
			args = append(args, "--volume", fmt.Sprintf("%s:%s", v.Volume, v.Path))
		}

		if v.HostPath != "" {
			args = append(args, "--volume", fmt.Sprintf("%s:%s", v.HostPath, v.Path))
		}
	}

	args = append(args, image)
	args = append(args, arg...)

	return args
}
