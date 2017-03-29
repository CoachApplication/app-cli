package main

import (
	"context"
	"os"
	"testing"

	cli_api "github.com/CoachApplication/app-cli/api"
)

func Test_Main(t *testing.T) {
	ctx := context.Background()

	workingDir, _ = os.Getwd()
	settings := MakeLocalAPISettings(workingDir, ctx)

	// Discover the current User (paths for the user like ~ and ~/.config/wundertools)
	DiscoverUser(&settings)

	// Discover for the project
	DiscoverProject(&settings)

	// if we have an environment set, then discover it.
	if environment != "" {
		DiscoverEnvironment(&settings, environment)
	}

	/**
	 * The settings object has been created, now create an API,
	 * and use it to get a list of operations, which we will
	 * convert to CLI commands.
	 */

	// Build a local Application implementation from the settings
	local, _ := cli_api.MakeLocalApp(ctx, settings)

	// Get a list of operations from the API
	localOps := local.Operations()

	t.Error(localOps)
}
