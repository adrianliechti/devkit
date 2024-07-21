package template

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/engine"
)

type template string

const (
	Category = "TEMPLATES"
)

var (
	TemplateAngular template = "angular"
	TemplateASPNET  template = "aspnet"
	TemplateGolang  template = "golang"
	TemplateNginx   template = "nginx"
	TemplatePython  template = "python"
	TemplateReact   template = "react"
	TemplateSpring  template = "spring"
)

var Command = &cli.Command{
	Name:  "template",
	Usage: "create new applications from template",

	HideHelpCommand: true,

	Category: Category,

	Commands: []*cli.Command{
		reactCommand,
		angularCommand,
		golangCommand,
		pythonCommand,
		springCommand,
		aspnetCommand,
		nginxCommand,
	},
}

func Name(ctx context.Context, cmd *cli.Command, placeholder string) string {
	name := cmd.String("name")

	if name == "" {
		name, _ = cli.Prompt("App Name", placeholder)
	}

	return name
}

func MustName(ctx context.Context, cmd *cli.Command, placeholder string) string {
	name := Name(ctx, cmd, placeholder)

	if name == "" {
		cli.Fatal(errors.New("missing name"))
	}

	return name
}

func Group(ctx context.Context, cmd *cli.Command, placeholder string) string {
	group := cmd.String("group")

	if group == "" {
		group, _ = cli.Prompt("App Group", placeholder)
	}

	return group
}

func MustGroup(ctx context.Context, cmd *cli.Command, placeholder string) string {
	group := Group(ctx, cmd, placeholder)

	if group == "" {
		cli.Fatal(errors.New("missing group"))
	}

	return group
}

func runTemplate(ctx context.Context, client engine.Client, path string, template template, options templateOptions) error {
	if options.Name == "" {
		options.Name = "demo"
	}

	if options.Version == "" {
		options.Version = "1.0.0"
	}

	path, err := app.EmptyDir(path, options.Name)

	if err != nil {
		return err
	}

	image := fmt.Sprintf("ghcr.io/adrianliechti/loop-template:%s", template)

	cli.MustRun("Pulling Image...", func() error {
		return client.Pull(ctx, image, "", engine.PullOptions{})
	})

	runOptions := docker.RunOptions{
		Env: options.env(),

		Volumes: []engine.ContainerMount{
			{
				Path:     "/src",
				HostPath: path,
			},
		},
	}

	return docker.RunInteractive(ctx, image, runOptions)
}

type templateOptions struct {
	Group   string
	Name    string
	Version string

	Host string

	EnableIngress     bool
	EnablePersistence bool
}

func (o *templateOptions) env() map[string]string {
	name := o.Name
	version := o.Version

	group := strings.ToLower(o.Group)
	artifact := strings.ToLower(name)

	chart := strings.ToLower(name)
	chartVersion := strings.ToLower(version)

	image := strings.ToLower(name)
	imageTag := strings.ToLower(version)

	host := strings.ToLower(o.Host)

	ingress := strconv.FormatBool(o.EnableIngress)
	persistent := strconv.FormatBool(o.EnablePersistence)

	result := map[string]string{
		"APP_NAME": name,

		"APP_GROUP":    group,
		"APP_ARTIFACT": artifact,

		"CHART_NAME":    chart,
		"CHART_VERSION": chartVersion,

		"IMAGE_REPOSITORY": image,
		"IMAGE_TAG":        imageTag,

		"APP_HOSTNAME": host,

		"APP_INGRESS":    ingress,
		"APP_PERSISTENT": persistent,
	}

	return result
}
