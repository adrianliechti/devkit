package openapi

import (
	"context"
	"os"
	"path/filepath"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var mergeCommand = &cli.Command{
	Name:  "merge",
	Usage: "merge openapi schema",

	Category: utility.Category,

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		path := cli.MustFile("Select Swagger/OpenAPI schema", []string{".json", ".yaml"})

		return runMerge(ctx, client, path)
	},
}

func runMerge(ctx context.Context, client engine.Client, path string) error {
	path, err := filepath.Abs(path)

	if err != nil {
		return err
	}

	tmpdir, err := os.MkdirTemp("", "openapi-")

	if err != nil {
		return err
	}

	defer os.RemoveAll(tmpdir)

	dir, file := filepath.Split(path)

	output := filepath.Join(dir, "merged.yaml")

	if _, err := os.Stat(output); err == nil {
		if ok, err := cli.Confirm("Overwrite "+output+"?", false); err != nil || !ok {
			return nil
		}
	}

	args := []string{
		"generate",

		"-l", "openapi-yaml",
		"-i", "/src/" + file,
		"-o", "/output",
		"-DoutputFile=openapi_merged.yaml",

		"--flatten-inline-schem",
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
				HostPath: tmpdir,
			},
		},
	}

	if err := client.Run(ctx, container, engine.RunOptions{}); err != nil {
		return err
	}

	data, err := os.ReadFile(tmpdir + "/openapi_merged.yaml")

	if err != nil {
		return err
	}

	if err := os.WriteFile(output, data, 0644); err != nil {
		return err
	}

	return nil
}
