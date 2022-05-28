package docker

import (
	"context"
	"io"
	"os/exec"
)

type PullOptions struct {
	Stdout io.Writer
	Stderr io.Writer
}

func Pull(ctx context.Context, image string, options PullOptions) error {
	tool, _, err := Tool(ctx)

	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, tool, "pull", image)
	cmd.Stdout = options.Stdout
	cmd.Stderr = options.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
