package vault

import (
	"fmt"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"

	"github.com/sethvargo/go-password/password"
)

func createCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create instance",

		Flags: []cli.Flag{
			app.PortFlag(""),
		},

		Action: func(c *cli.Context) error {
			ctx := c.Context
			image := "vault:latest"

			port := app.MustPortOrRandom(c, "", 8200)

			token := password.MustGenerate(10, 4, 0, false, false)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: Vault,
				},

				Env: map[string]string{
					"VAULT_DEV_ROOT_TOKEN_ID": token,
				},

				Ports: map[int]int{
					port: 8200,
				},
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"URL", fmt.Sprintf("http://localhost:%d", port)},
				{"Token", token},
			})

			return nil
		},
	}
}
