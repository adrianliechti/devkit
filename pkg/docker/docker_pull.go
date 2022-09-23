package docker

import (
	"context"
	"io"
	"os"
	"os/exec"
)

type PullOptions struct {
	Platform string

	Stdout io.Writer
	Stderr io.Writer
}

func Pull(ctx context.Context, image string, options PullOptions) error {
	tool, _, err := Info(ctx)

	if err != nil {
		return err
	}

	pull := exec.CommandContext(ctx, tool, pullArgs(image, options)...)
	pull.Stdout = options.Stdout
	pull.Stderr = options.Stderr

	return pull.Run()
}

func PullInteractive(ctx context.Context, image string, options PullOptions) error {

	if options.Stdout == nil {
		options.Stdout = os.Stdout
	}

	if options.Stderr == nil {
		options.Stderr = os.Stderr
	}

	return Pull(ctx, image, options)
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
