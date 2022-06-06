package docker

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func PortForward(ctx context.Context, container string, source, target int) error {
	tool, _, err := Tool(ctx)

	if err != nil {
		return err
	}

	inspect, err := exec.CommandContext(ctx, tool, "inspect", "--format", "{{ .NetworkSettings.IPAddress }}", container).Output()

	if err != nil {
		return err
	}

	address := strings.TrimRight(string(inspect), "\n")

	if address == "" {
		return errors.New("invalid container ip")
	}

	args := []string{
		"run", "-i", "--rm",
		"-p", fmt.Sprintf("127.0.0.1:%d:%d", source, source),

		"alpine/socat",
		fmt.Sprintf("TCP4-LISTEN:%d,fork,reuseaddr", source),
		fmt.Sprintf("TCP4:%s:%d", address, target),
	}

	socat := exec.CommandContext(ctx, tool, args...)
	socat.Stdout = os.Stdout
	socat.Stderr = os.Stderr

	return socat.Run()
}
