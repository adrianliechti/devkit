package sonarqube

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/platform"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	SonarQube = "sonarqube"
)

var Command = &cli.Command{
	Name:  SonarQube,
	Usage: "local SonarQube server",

	Category: platform.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(SonarQube),

		CreateCommand(),
		common.DeleteCommand(SonarQube),

		common.LogsCommand(SonarQube),
		common.ShellCommand(SonarQube, "/bin/bash"),
	},
}
