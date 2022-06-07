package app

import (
	"strings"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/system"
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

func Port(c *cli.Context, name string) int {
	return c.Int(PortFlagName(name))
}

func MustPort(c *cli.Context, name string) int {
	port := Port(c, name)

	if port <= 0 {
		cli.Fatal(PortFlagName(name) + " missing")
	}

	return port
}

func PortOrRandom(c *cli.Context, name string, preference int) (int, error) {
	port := Port(c, name)

	if port > 0 {
		return port, nil
	}

	return system.FreePort(preference)
}

func MustPortOrRandom(c *cli.Context, name string, preference int) int {
	port, err := PortOrRandom(c, name, preference)

	if err != nil {
		cli.Fatal(err)
	}

	return port
}

func RandomPort(c *cli.Context, preference int) (int, error) {
	return system.FreePort(preference)
}

func MustRandomPort(c *cli.Context, preference int) int {
	port, err := RandomPort(c, preference)

	if err != nil {
		cli.Fatal(err)
	}

	return port
}
