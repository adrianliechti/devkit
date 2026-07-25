package image

import (
	"context"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/go-cli"
)

var Command = &cli.Command{
	Name:  "image",
	Usage: "docker/oci image tools",

	HideHelpCommand: true,

	Category: utility.Category,

	Commands: []*cli.Command{
		toolCommand("lint", "lint Dockerfile using dockle", func(image string) engine.Container {
			return engine.Container{
				Image: "goodwithtech/dockle:v0.4.15",

				Env: map[string]string{
					"DOCKER_CONTENT_TRUST": "1",
				},

				Args:   []string{image},
				Mounts: []engine.ContainerMount{dockerSocket},
			}
		}),

		toolCommand("scan", "scan image vulnerabilities using trivy", func(image string) engine.Container {
			return engine.Container{
				Image: "aquasec/trivy:0.59.0",

				Args: []string{"--quiet", "image", image},

				Mounts: []engine.ContainerMount{
					{Path: "/root/.cache/", Volume: "trivy-cache"},
				},
			}
		}),

		toolCommand("inspect", "inspect image using whaler", func(image string) engine.Container {
			return engine.Container{
				Image:    "pegleg/whaler",
				Platform: "linux/amd64",

				Args:   []string{image},
				Mounts: []engine.ContainerMount{dockerSocket},
			}
		}),

		toolCommand("bom", "show image bill of material using syft", func(image string) engine.Container {
			return engine.Container{
				Image: "anchore/syft",

				Args:   []string{"-o", "syft-table", image},
				Mounts: []engine.ContainerMount{dockerSocket},
			}
		}),

		toolCommand("browse", "browse image using dive (interactive)", func(image string) engine.Container {
			return engine.Container{
				Image:    "wagoodman/dive:v0.12",
				Platform: "linux/amd64",

				Args:   []string{image},
				Mounts: []engine.ContainerMount{dockerSocket},
			}
		}),
	},
}

// The tools below inspect images through the engine, not the registry.
var dockerSocket = engine.ContainerMount{
	Path:     "/var/run/docker.sock",
	HostPath: "/var/run/docker.sock",
}

// toolCommand runs a containerized tool against the image given via --image.
func toolCommand(name, usage string, spec func(image string) engine.Container) *cli.Command {
	return &cli.Command{
		Name:  name,
		Usage: usage,

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "image",
				Usage:    "docker image",
				Required: true,
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := app.MustClient(ctx, cmd)
			image := cmd.String("image")

			// Best effort: locally built images cannot be pulled, but can be analyzed.
			cli.MustRun("Pulling Image...", func() error {
				client.Pull(ctx, image, "", engine.PullOptions{})
				return nil
			})

			return client.Run(ctx, spec(image), engine.RunOptions{})
		},
	}
}
