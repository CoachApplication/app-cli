package main

import (
	"context"

	handler_local "github.com/james-nesbitt/coach-local"
)

// Construct a LocalAPI by checking some paths for the current user.
func MakeLocalAPISettings(workingDir string, ctx context.Context) handler_local.Settings {
	// create an initial empty settings
	return handler_local.Settings{
		ExecPath: workingDir,
		Paths:    *handler_local.NewSettingScopePaths(),
	}
}
