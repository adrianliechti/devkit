package sonarqube

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/sonarqube"
)

const (
	SonarQube = "sonarqube"
)

var Command = &cli.Command{
	Name:  SonarQube,
	Usage: "local SonarQube server",

	Category: app.PlatformCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(SonarQube),

		catalog.CreateCommand(SonarQube, sonarqube.New, sonarqube.Info),
		catalog.DeleteCommand(SonarQube),

		catalog.InfoCommand(SonarQube, sonarqube.Info),
		catalog.LogsCommand(SonarQube),

		catalog.ShellCommand(SonarQube, sonarqube.DefaultShell),
	},
}
