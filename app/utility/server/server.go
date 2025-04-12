package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/go-cli"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Command = &cli.Command{
	Name:  "server",
	Usage: "start Web server",

	Category: utility.Category,

	Flags: []cli.Flag{
		app.PortFlag(""),

		&cli.BoolFlag{
			Name:  "spa",
			Usage: "enable SPA redirect",
		},

		&cli.StringFlag{
			Name:  "index",
			Usage: "index file name",
			Value: "index.html",
		},
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		port := app.MustPortOrRandom(ctx, cmd, "", 3000)

		spa := cmd.Bool("spa")
		index := cmd.String("index")

		if index == "" {
			index = "index.html"
		}

		return startWebServer(ctx, port, index, spa)
	},
}

func startWebServer(ctx context.Context, port int, index string, spa bool) error {
	root, err := os.Getwd()

	if err != nil {
		return err
	}

	if port == 0 {
		port = 3000
	}

	e := echo.New()
	e.HidePort = true
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri}\n",
	}))

	e.Use(middleware.CORS())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Index: index,
		HTML5: spa,

		Browse:     true,
		Filesystem: http.Dir(root),
	}))

	go func() {
		<-ctx.Done()
		e.Close()
	}()

	cli.Infof("Server started on http://127.0.0.1:%d", port)
	cli.Info()

	if err := e.Start(fmt.Sprintf("127.0.0.1:%d", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
