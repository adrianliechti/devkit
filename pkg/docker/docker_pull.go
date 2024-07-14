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

func Pull(ctx context.Context, image, platform string, options PullOptions) error {
	tool, _, err := Info(ctx)

	if err != nil {
		return err
	}

	pull := exec.CommandContext(ctx, tool, pullArgs(image, platform, options)...)
	pull.Stdout = options.Stdout
	pull.Stderr = options.Stderr

	return pull.Run()
}

func pullArgs(image, platform string, options PullOptions) []string {
	args := []string{
		"pull",
	}

	if platform != "" {
		args = append(args, "--platform", platform)
	}

	args = append(args, image)

	return args
}
