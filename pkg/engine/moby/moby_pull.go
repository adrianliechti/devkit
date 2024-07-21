package moby

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"strings"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/cpuguy83/dockercfg"
	"github.com/distribution/reference"
	"github.com/docker/docker/api/types/image"
)

func (m *Moby) Pull(ctx context.Context, reference, platform string, options engine.PullOptions) error {
	if options.Stdout == nil {
		options.Stdout = io.Discard
	}

	if options.Stderr == nil {
		options.Stderr = io.Discard
	}

	out, err := m.client.ImagePull(ctx, reference, image.PullOptions{
		Platform:     platform,
		RegistryAuth: pullCredentials(reference),
	})

	if err != nil {
		return err
	}

	defer out.Close()

	if _, err := io.Copy(options.Stdout, out); err != nil {
		return err
	}

	return nil
}

func pullCredentials(image string) string {
	// echo "https://index.docker.io/v1/"|docker-credential-desktop get
	// {"ServerURL":"https://index.docker.io/v1/","Username":"xxxxx","Secret":"xxxxx"}

	ref, err := reference.ParseNormalizedNamed(image)

	if err != nil {
		return ""
	}

	parts := strings.Split(ref.Name(), "/")

	if len(parts) == 0 {
		return ""
	}

	host := dockercfg.ResolveRegistryHost(parts[0])

	username, password, err := dockercfg.GetRegistryCredentials(host)

	if err != nil {
		return ""
	}

	type AuthConfig struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}

	data, err := json.Marshal(AuthConfig{
		Username: username,
		Password: password,
	})

	if err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(data)
}
