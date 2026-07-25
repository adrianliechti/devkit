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

// MustPortOrRandom returns the port given on the command line, or a free port
// (preferring the given one) when the flag was not set.
func MustPortOrRandom(ctx context.Context, cmd *cli.Command, name string, preference int) int {
	if port := int(cmd.Int(PortFlagName(name))); port > 0 {
		return port
	}

	port, err := system.FreePort(preference)

	if err != nil {
		cli.Fatal(err)
	}

	return port
}
