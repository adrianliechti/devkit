package proxy

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/proxy"
)

var Command = &cli.Command{
	Name:  "proxy",
	Usage: "start temporary Proxy server",

	Category: utility.Category,

	Flags: []cli.Flag{
		app.PortFlag(""),

		&cli.StringFlag{
			Name:  "username",
			Usage: "proxy username",
		},

		&cli.StringFlag{
			Name:  "password",
			Usage: "proxy password",
		},
	},

	Action: func(c *cli.Context) error {
		port := app.MustPortOrRandom(c, "", 3128)

		username := c.String("username")
		password := c.String("password")

		return runProxy(c.Context, port, username, password)
	},
}

func runProxy(ctx context.Context, port int, username, password string) error {
	if port == 0 {
		port = 3128
	}

	config := proxy.Config{
		Username: username,
		Password: password,
	}

	server := &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%d", port),

		Handler: proxy.New(config),

		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	go func() {
		<-ctx.Done()

		server.Shutdown(context.Background())
	}()

	cli.Infof("Starting proxy at port %d", port)

	return server.ListenAndServe()
}
