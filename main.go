package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/adrianliechti/devkit/app/database/db2"
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
	"github.com/adrianliechti/devkit/app/platform/jenkins"
	"github.com/adrianliechti/devkit/app/platform/sonarqube"
	"github.com/adrianliechti/devkit/app/storage/minio"
	"github.com/adrianliechti/devkit/app/storage/vault"
	"github.com/adrianliechti/devkit/app/template"
	"github.com/adrianliechti/devkit/app/utility/cloc"
	"github.com/adrianliechti/devkit/app/utility/code"
	"github.com/adrianliechti/devkit/app/utility/git"
	"github.com/adrianliechti/devkit/app/utility/image"
	"github.com/adrianliechti/devkit/app/utility/proxy"
	"github.com/adrianliechti/devkit/app/utility/server"
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

		Commands: []*cli.Command{
			mariadb.Command,
			postgres.Command,
			mongodb.Command,
			mssql.Command,
			db2.Command,

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

			jenkins.Command,
			sonarqube.Command,

			git.Command,
			image.Command,

			template.Command,

			cloc.Command,

			code.Command,
			server.Command,
			proxy.Command,
		},
	}
}
