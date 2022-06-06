package engine

import (
	"context"
	"io"
)

type PullOptions struct {
	Platform string

	Stdout io.Writer
	Stderr io.Writer
}

type RemoveOptions struct {
}

type LogsOptions struct {
	Follow bool

	Stdout io.Writer
	Stderr io.Writer
}

type Engine interface {
	Pull(ctx context.Context, image string, options PullOptions) error

	Remove(ctx context.Context, container string, options RemoveOptions) error

	Logs(ctx context.Context, container string, options LogsOptions) error
}
