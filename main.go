package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/adrianliechti/devkit/app/cloc"
	"github.com/adrianliechti/devkit/app/cluster"
	"github.com/adrianliechti/devkit/app/elasticsearch"
	"github.com/adrianliechti/devkit/app/etcd"
	"github.com/adrianliechti/devkit/app/git"
	"github.com/adrianliechti/devkit/app/image"
	"github.com/adrianliechti/devkit/app/influxdb"
	"github.com/adrianliechti/devkit/app/kafka"
	"github.com/adrianliechti/devkit/app/mariadb"
	"github.com/adrianliechti/devkit/app/minio"
	"github.com/adrianliechti/devkit/app/mongodb"
	"github.com/adrianliechti/devkit/app/mssql"
	"github.com/adrianliechti/devkit/app/nats"
	"github.com/adrianliechti/devkit/app/postgres"
	"github.com/adrianliechti/devkit/app/proxy"
	"github.com/adrianliechti/devkit/app/redis"
	"github.com/adrianliechti/devkit/app/vault"
	"github.com/adrianliechti/devkit/app/webserver"
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
		// Name:    "loop",
		Version: version,
		// Usage:   "DevOps Loop",

		HideHelpCommand: true,

		// Flags: []cli.Flag{
		// 	app.KubeconfigFlag,
		// },

		Commands: []*cli.Command{
			cluster.Command,

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
