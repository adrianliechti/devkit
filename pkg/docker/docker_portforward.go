package docker

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func PortForward(ctx context.Context, container string, source, target int) error {
	tool, _, err := Tool(ctx)

	if err != nil {
		return err
	}

	info, err := Info(ctx, container)

	if err != nil {
		return err
	}

	if info.IPAddress == "" {
		return errors.New("invalid container ip")
	}

	args := []string{
		"run", "-i", "--rm",
		"-p", fmt.Sprintf("127.0.0.1:%d:%d", source, source),

		"alpine/socat",
		fmt.Sprintf("TCP4-LISTEN:%d,fork,reuseaddr", source),
		fmt.Sprintf("TCP4:%s:%d", info.IPAddress, target),
	}

	socat := exec.CommandContext(ctx, tool, args...)
	socat.Stdout = os.Stdout
	socat.Stderr = os.Stderr

	return socat.Run()
}
