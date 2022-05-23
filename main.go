package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/adrianliechti/devkit/app/database/elasticsearch"
	"github.com/adrianliechti/devkit/app/database/etcd"
	"github.com/adrianliechti/devkit/app/database/influxdb"
	"github.com/adrianliechti/devkit/app/database/mariadb"
	"github.com/adrianliechti/devkit/app/database/mongodb"
	"github.com/adrianliechti/devkit/app/database/mssql"
	"github.com/adrianliechti/devkit/app/database/postgres"
	"github.com/adrianliechti/devkit/app/database/redis"
	"github.com/adrianliechti/devkit/app/messaging/kafka"
	"github.com/adrianliechti/devkit/app/messaging/nats"
	"github.com/adrianliechti/devkit/app/storage/minio"
	"github.com/adrianliechti/devkit/app/storage/vault"
	"github.com/adrianliechti/devkit/app/template"
	"github.com/adrianliechti/devkit/app/utility/cloc"
	"github.com/adrianliechti/devkit/app/utility/code"
	"github.com/adrianliechti/devkit/app/utility/git"
	"github.com/adrianliechti/devkit/app/utility/image"
	"github.com/adrianliechti/devkit/app/utility/proxy"
	"github.com/adrianliechti/devkit/app/utility/webserver"
	"github.com/adrianliechti/devkit/pkg/cli"
)

var version string

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGPIPE)
	defer stop()

	app := initApp()

	if err := app.RunContext(ctx, os.Args); err != nil {
		cli.Fatal(err)
	}
}

func initApp() cli.App {
	return cli.App{
		Version: version,

		HideHelpCommand: true,

		// Flags: []cli.Flag{
		// 	app.KubeconfigFlag,
		// },

		Commands: []*cli.Command{
			mariadb.Command,
			postgres.Command,
			mongodb.Command,
			mssql.Command,

			etcd.Command,
			redis.Command,
			influxdb.Command,
			elasticsearch.Command,

			minio.Command,
			vault.Command,

			nats.Command,
			kafka.Command,
			// rabbitmqCommand,

			// registryCommand,
			// mailtrapCommand,

			// codeCommand,
			// grafanaCommand,
			// jupyterCommand,

			git.Command,
			image.Command,

			template.Command,

			code.Command,
			cloc.Command,
			proxy.Command,
			webserver.Command,
		},
		// 	// Cluster
		// 	cluster.Command,

		// 	application.Command,
		// 	config.Command,
		// 	connect.Command,
		// 	catapult.Command,
		// 	dashboard.Command,

		// 	// Development
		// 	{
		// 		Name:  "local",
		// 		Usage: "local development instances",

		// 		Category: app.CategoryDevelopment,

		// 		HideHelpCommand: true,

		// 		Subcommands: []*cli.Command{
		//
		// 		},
		// 	},
		// 	remote.Command,
		// 	expose.Command,

		// 	// Utilities
		// 	git.Command,
		// 	tool.Command,
		// 	image.Command,
		// 	template.Command,
		// },
	}
}
