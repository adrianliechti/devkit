package docker

import (
	"context"
	"io"
	"os"
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

	if options.Stdout == nil {
		options.Stdout = os.Stdout
	}

	if options.Stderr == nil {
		options.Stderr = os.Stderr
	}

	cmd := exec.CommandContext(ctx, tool, "pull", image)
	cmd.Stdout = options.Stdout
	cmd.Stderr = options.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
