package cluster

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/kind"
	"github.com/adrianliechti/devkit/pkg/kubectl"
)

func DashboardCommand() *cli.Command {
	return &cli.Command{
		Name:  "dashboard",
		Usage: "open instance Dashboard",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "cluster name",
			},
			&cli.StringFlag{
				Name:  app.PortFlag.Name,
				Usage: "local dashboard port",
			},
		},

		Action: func(c *cli.Context) error {
			port := app.MustPortOrRandom(c, 9090)
			name := c.String("name")

			if name == "" {
				name = MustCluster(c.Context)
			}

			dir, err := ioutil.TempDir("", "kind")

			if err != nil {
				return err
			}

			defer os.RemoveAll(dir)
			kubeconfig := path.Join(dir, "kubeconfig")

			if err := kind.Kubeconfig(c.Context, name, kubeconfig); err != nil {
				return err
			}

			time.AfterFunc(3*time.Second, func() {
				url := fmt.Sprintf("http://127.0.0.1:%d", port)
				cli.OpenURL(url)
			})

			if err := kubectl.Invoke(c.Context, kubeconfig, "port-forward", "-n", "loop", "service/dashboard", fmt.Sprintf("%d:80", port)); err != nil {
				return err
			}

			return nil
		},
	}
}
