package cluster

import (
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/helm"
	"github.com/adrianliechti/devkit/pkg/kind"
	"github.com/adrianliechti/devkit/pkg/kubectl"

	"github.com/adrianliechti/devkit/app/cluster/extension/dashboard"
	"github.com/adrianliechti/devkit/app/cluster/extension/observability"
)

func SetupCommand() *cli.Command {
	return &cli.Command{
		Name:  "setup",
		Usage: "setup as current cluster",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "cluster name",
			},
		},

		Action: func(c *cli.Context) error {
			name := c.String("name")

			if _, _, err := docker.Tool(c.Context); err != nil {
				return err
			}

			if _, _, err := kind.Tool(c.Context); err != nil {
				return err
			}

			if _, _, err := helm.Tool(c.Context); err != nil {
				return err
			}

			if _, _, err := kubectl.Tool(c.Context); err != nil {
				return err
			}

			// if err := kind.Create(c.Context, name); err != nil {
			// 	return err
			// }

			for _, image := range append(dashboard.Images, observability.Images...) {
				docker.Pull(c.Context, image)
				kind.LoadImage(c.Context, name, image)
			}

			namespace := "loop"
			kubeconfig := ""

			if err := dashboard.Install(c.Context, kubeconfig, namespace); err != nil {
				return err
			}

			if err := observability.Install(c.Context, kubeconfig, namespace); err != nil {
				return err
			}

			return nil
		},
	}
}
