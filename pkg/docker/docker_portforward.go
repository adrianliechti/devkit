package docker

import (
	"context"
	"encoding/json"
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

	type inspectType struct {
		NetworkSettings struct {
			IPAddress string `json:"IPAddress"`
		} `json:"NetworkSettings"`
	}

	inspect := exec.CommandContext(ctx, tool, "inspect", container)

	data, err := inspect.Output()

	if err != nil {
		return err
	}

	var info []inspectType

	if err := json.Unmarshal(data, &info); err != nil {
		return err
	}

	if len(info) != 1 {
		return errors.New("invalid container")
	}

	ip := info[0].NetworkSettings.IPAddress

	if ip == "" {
		return errors.New("invalid container ip")
	}

	args := []string{
		"run", "-i", "--rm",
	}

	args = append(args, "--publish", fmt.Sprintf("127.0.0.1:%d:%d", source, source))

	args = append(args, "alpine/socat")

	args = append(args, fmt.Sprintf("TCP4-LISTEN:%d,fork,reuseaddr", source), fmt.Sprintf("TCP4:%s:%d", ip, target))

	socat := exec.CommandContext(ctx, tool, args...)
	socat.Stdout = os.Stdout
	socat.Stderr = os.Stderr

	return socat.Run()

}
