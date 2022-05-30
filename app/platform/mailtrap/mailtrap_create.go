package mailtrap

import (
	"fmt"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/sethvargo/go-password/password"
)

func CreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create instance",

		Flags: []cli.Flag{
			app.PortFlag(""),
		},

		Action: func(c *cli.Context) error {
			ctx := c.Context
			image := "adrianliechti/loop-mailtrap"

			port := app.MustPortOrRandom(c, "", 2525)

			username := "admin"
			password := password.MustGenerate(10, 4, 0, false, false)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: MailTrap,
				},

				Platform: "linux/amd64",

				Env: map[string]string{
					"USERNAME": username,
					"PASSWORD": password,
				},

				Ports: map[int]int{
					port: 25,
				},
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"Host", fmt.Sprintf("localhost:%d", port)},
				{"Username", username},
				{"Password", password},
			})

			return nil
		},
	}
}
