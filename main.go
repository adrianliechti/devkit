package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/app/template"
	"github.com/adrianliechti/devkit/app/utility/cloc"
	"github.com/adrianliechti/devkit/app/utility/code"
	"github.com/adrianliechti/devkit/app/utility/git"
	"github.com/adrianliechti/devkit/app/utility/image"
	"github.com/adrianliechti/devkit/app/utility/proxy"
	"github.com/adrianliechti/devkit/app/utility/server"

	"github.com/adrianliechti/devkit/catalog/activemq"
	"github.com/adrianliechti/devkit/catalog/artifactory"
	"github.com/adrianliechti/devkit/catalog/azurite"
	"github.com/adrianliechti/devkit/catalog/cassandra"
	"github.com/adrianliechti/devkit/catalog/db2"
	"github.com/adrianliechti/devkit/catalog/elasticsearch"
	"github.com/adrianliechti/devkit/catalog/etcd"
	"github.com/adrianliechti/devkit/catalog/immudb"
	"github.com/adrianliechti/devkit/catalog/influxdb"
	"github.com/adrianliechti/devkit/catalog/jcr"
	"github.com/adrianliechti/devkit/catalog/jenkins"
	"github.com/adrianliechti/devkit/catalog/jupyter"
	"github.com/adrianliechti/devkit/catalog/kafka"
	"github.com/adrianliechti/devkit/catalog/keycloak"
	"github.com/adrianliechti/devkit/catalog/mailpit"
	"github.com/adrianliechti/devkit/catalog/mariadb"
	"github.com/adrianliechti/devkit/catalog/minio"
	"github.com/adrianliechti/devkit/catalog/mongodb"
	"github.com/adrianliechti/devkit/catalog/mosquitto"
	"github.com/adrianliechti/devkit/catalog/mssql"
	"github.com/adrianliechti/devkit/catalog/mysql"
	"github.com/adrianliechti/devkit/catalog/nats"
	"github.com/adrianliechti/devkit/catalog/nexus"
	"github.com/adrianliechti/devkit/catalog/oracle"
	"github.com/adrianliechti/devkit/catalog/postgres"
	"github.com/adrianliechti/devkit/catalog/rabbitmq"
	"github.com/adrianliechti/devkit/catalog/redis"
	"github.com/adrianliechti/devkit/catalog/sonarqube"
	"github.com/adrianliechti/devkit/catalog/vault"

	"github.com/adrianliechti/devkit/pkg/cli"
)

var version string

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer stop()

	app := initApp()

	if err := app.Run(ctx, os.Args); err != nil {
		cli.Fatal(err)
	}
}

func initApp() cli.Command {
	return cli.Command{
		Usage: "DevKit",

		Suggest: true,
		Version: version,

		HideHelpCommand: true,

		Commands: []*cli.Command{
			catalog.Command(&activemq.Manager{}),
			catalog.Command(&artifactory.Manager{}),
			catalog.Command(&azurite.Manager{}),
			catalog.Command(&cassandra.Manager{}),
			catalog.Command(&db2.Manager{}),
			catalog.Command(&elasticsearch.Manager{}),
			catalog.Command(&etcd.Manager{}),
			catalog.Command(&immudb.Manager{}),
			catalog.Command(&influxdb.Manager{}),
			catalog.Command(&jcr.Manager{}),
			catalog.Command(&jenkins.Manager{}),
			catalog.Command(&jupyter.Manager{}),
			catalog.Command(&kafka.Manager{}),
			catalog.Command(&keycloak.Manager{}),
			catalog.Command(&mailpit.Manager{}),
			catalog.Command(&mariadb.Manager{}),
			catalog.Command(&minio.Manager{}),
			catalog.Command(&mongodb.Manager{}),
			catalog.Command(&mosquitto.Manager{}),
			catalog.Command(&mssql.Manager{}),
			catalog.Command(&mysql.Manager{}),
			catalog.Command(&nats.Manager{}),
			catalog.Command(&nexus.Manager{}),
			catalog.Command(&oracle.Manager{}),
			catalog.Command(&postgres.Manager{}),
			catalog.Command(&rabbitmq.Manager{}),
			catalog.Command(&redis.Manager{}),
			catalog.Command(&sonarqube.Manager{}),
			catalog.Command(&vault.Manager{}),

			template.Command,

			git.Command,
			image.Command,

			cloc.Command,
			code.Command,
			proxy.Command,
			server.Command,
		},
	}
}
