package app

import (
	"context"
	"strings"

	"github.com/adrianliechti/devkit/pkg/system"
	"github.com/adrianliechti/go-cli"
)

func PortFlagName(name string) string {
	if name == "" || strings.ToLower(name) == "port" {
		return "port"
	}

	return strings.ToLower(name + "-port")
}

func PortFlag(name string) *cli.IntFlag {
	flagUsage := "port"

	if strings.ToLower(name) == "port" {
		name = ""
	}

	if name != "" {
		flagUsage = name + " port"
	}

	return &cli.IntFlag{
		Name:  PortFlagName(name),
		Usage: flagUsage,
	}
}

func Port(ctx context.Context, cmd *cli.Command, name string) int {
	return int(cmd.Int(PortFlagName(name)))
}

func MustPort(ctx context.Context, cmd *cli.Command, name string) int {
	port := Port(ctx, cmd, name)

	if port <= 0 {
		cli.Fatal(PortFlagName(name) + " missing")
	}

	return port
}

func PortOrRandom(ctx context.Context, cmd *cli.Command, name string, preference int) (int, error) {
	port := Port(ctx, cmd, name)

	if port > 0 {
		return port, nil
	}

	return system.FreePort(preference)
}

func MustPortOrRandom(ctx context.Context, cmd *cli.Command, name string, preference int) int {
	port, err := PortOrRandom(ctx, cmd, name, preference)

	if err != nil {
		cli.Fatal(err)
	}

	return port
}

func RandomPort(ctx context.Context, cmd *cli.Command, preference int) (int, error) {
	return system.FreePort(preference)
}

func MustRandomPort(ctx context.Context, cmd *cli.Command, preference int) int {
	port, err := RandomPort(ctx, cmd, preference)

	if err != nil {
		cli.Fatal(err)
	}

	return port
}
