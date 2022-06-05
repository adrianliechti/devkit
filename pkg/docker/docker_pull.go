package docker

import (
	"context"
	"io"
	"os/exec"
)

type PullOptions struct {
	Platform string

	Stdout io.Writer
	Stderr io.Writer
}

func Pull(ctx context.Context, image string, options PullOptions) error {
	tool, _, err := Tool(ctx)

	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, tool, pullArgs(image, options)...)
	cmd.Stdout = options.Stdout
	cmd.Stderr = options.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func pullArgs(image string, options PullOptions) []string {
	args := []string{
		"pull",
	}

	if options.Platform != "" {
		args = append(args, "--platform", options.Platform)
	}

	args = append(args, image)

	return args
}
