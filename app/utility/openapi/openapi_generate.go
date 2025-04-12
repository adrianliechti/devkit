package openapi

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/go-cli"
)

var generateCommand = &cli.Command{
	Name:  "generate",
	Usage: "generate openapi client/server",

	Category: utility.Category,

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		path := cli.MustFile("Select Swagger/OpenAPI schema", []string{".json", ".yaml"})

		langs, err := listLanguages(ctx, client)

		if err != nil {
			return err
		}

		_, lang, err := cli.Select("Select language", langs)

		if err != nil {
			return err
		}

		return runGenerate(ctx, client, path, lang)
	},
}

func listLanguages(ctx context.Context, client engine.Client) ([]string, error) {
	container := engine.Container{
		Image: "swaggerapi/swagger-codegen-cli-v3",
		Args: []string{
			"langs",
		},
	}

	var data bytes.Buffer

	if err := client.Run(ctx, container, engine.RunOptions{Stdout: &data}); err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`Available languages: \[(.*)\]`)
	matches := re.FindStringSubmatch(data.String())

	if len(matches) < 2 {
		return nil, errors.New("failed to get available languages")
	}

	return strings.Split(matches[1], ", "), nil
}

func runGenerate(ctx context.Context, client engine.Client, path, language string) error {
	path, err := filepath.Abs(path)

	if err != nil {
		return err
	}

	dir, file := filepath.Split(path)

	output := filepath.Join(dir, "generated_"+language)

	if _, err := os.Stat(output); err == nil {
		if ok, err := cli.Confirm("Overwrite "+output+"?", false); err != nil || !ok {
			return nil
		}
	}

	os.MkdirAll(output, 0755)

	args := []string{
		"generate",

		"-l", language,
		"-i", "/src/" + file,
		"-o", "/output",
	}

	container := engine.Container{
		Image: "swaggerapi/swagger-codegen-cli-v3",
		Args:  args,

		Mounts: []engine.ContainerMount{
			{
				Path:     "/src",
				HostPath: dir,
			},

			{
				Path:     "/output",
				HostPath: output,
			},
		},
	}

	return client.Run(ctx, container, engine.RunOptions{})
}
