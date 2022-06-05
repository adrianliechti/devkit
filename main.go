package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/app/catalog/cassandra"
	"github.com/adrianliechti/devkit/app/catalog/db2"
	"github.com/adrianliechti/devkit/app/catalog/elasticsearch"
	"github.com/adrianliechti/devkit/app/catalog/etcd"
	"github.com/adrianliechti/devkit/app/catalog/grafana"
	"github.com/adrianliechti/devkit/app/catalog/immudb"
	"github.com/adrianliechti/devkit/app/catalog/influxdb"
	"github.com/adrianliechti/devkit/app/catalog/jenkins"
	"github.com/adrianliechti/devkit/app/catalog/jupyter"
	"github.com/adrianliechti/devkit/app/catalog/kafka"
	"github.com/adrianliechti/devkit/app/catalog/mailtrap"
	"github.com/adrianliechti/devkit/app/catalog/mariadb"
	"github.com/adrianliechti/devkit/app/catalog/minio"
	"github.com/adrianliechti/devkit/app/catalog/mongodb"
	"github.com/adrianliechti/devkit/app/catalog/mssql"
	"github.com/adrianliechti/devkit/app/catalog/mysql"
	"github.com/adrianliechti/devkit/app/catalog/nats"
	"github.com/adrianliechti/devkit/app/catalog/rabbitmq"
	"github.com/adrianliechti/devkit/app/catalog/redis"
	"github.com/adrianliechti/devkit/app/catalog/sonarqube"
	"github.com/adrianliechti/devkit/app/catalog/vault"
	"github.com/adrianliechti/devkit/app/template"
	"github.com/adrianliechti/devkit/app/utility/cloc"
	"github.com/adrianliechti/devkit/app/utility/code"
	"github.com/adrianliechti/devkit/app/utility/git"
	"github.com/adrianliechti/devkit/app/utility/image"
	"github.com/adrianliechti/devkit/app/utility/proxy"
	"github.com/adrianliechti/devkit/app/utility/server"
	"github.com/adrianliechti/devkit/pkg/catalog/postgres"
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
			mysql.Command,
			mariadb.Command,
			//postgres.Command,
			catalog.Command(&postgres.Manager{}),
			mongodb.Command,
			mssql.Command,
			cassandra.Command,
			db2.Command,

			etcd.Command,
			redis.Command,
			immudb.Command,
			influxdb.Command,
			elasticsearch.Command,

			nats.Command,
			kafka.Command,
			rabbitmq.Command,

			jenkins.Command,
			sonarqube.Command,
			grafana.Command,
			jupyter.Command,
			mailtrap.Command,

			minio.Command,
			vault.Command,

			template.Command,

			git.Command,
			image.Command,

			cloc.Command,

			code.Command,
			server.Command,
			proxy.Command,
		},
	}
}
