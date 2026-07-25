package git

import (
	"context"
	"errors"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/Masterminds/semver"
)

var (
	minimalVersion = semver.MustParse("2.0.0")

	errNotFound = errors.New("git not found. see https://git-scm.com/download")
	errOutdated = errors.New("git is outdated. see https://git-scm.com/download")
)

// Info returns the path and version of the local git installation.
func Info(ctx context.Context) (string, *semver.Version, error) {
	name := "git"

	if runtime.GOOS == "windows" {
		name = "git.exe"
	}

	path, err := exec.LookPath(name)

	if err != nil {
		return "", nil, errNotFound
	}

	v, err := version(ctx, path)

	if err != nil {
		return path, nil, err
	}

	if v.LessThan(minimalVersion) {
		return path, v, errOutdated
	}

	return path, v, nil
}

func version(ctx context.Context, path string) (*semver.Version, error) {
	cmd := exec.CommandContext(ctx, path, "--version")
	data, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	version := strings.TrimSpace(string(data))

	r, _ := regexp.Compile(`^.* version (\d+\.\d+(\.\d+)?).*`)

	if matches := r.FindStringSubmatch(version); len(matches) == 3 {
		return semver.NewVersion(matches[1])
	}

	return semver.NewVersion("0.0.0")
}
