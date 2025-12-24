package compose

import (
	"context"

	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/compose-spec/compose-go/v2/types"
)

func LoadCompose(ctx context.Context, name string, configFiles []string) (*types.Project, error) {
	opts, err := cli.NewProjectOptions(configFiles, cli.WithOsEnv, cli.WithDotEnv, cli.WithName(name))
	if err != nil {
		return nil, err
	}

	proj, err := opts.LoadProject(ctx)
	if err != nil {
		return nil, err
	}

	return proj, nil
}
