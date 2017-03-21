package main

import (
	"context"
	handler_local "github.com/james-nesbitt/coach-local"
	"os"
	"os/user"
	"path"
	"testing"
)

func makeLocalApi() handler_local.Settings {
	workingDir, _ = os.Getwd()
	ctx := context.Background()
	return MakeLocalAPISettings(workingDir, ctx)
}

func TestDiscoverUser(t *testing.T) {
	app := makeLocalApi()

	DiscoverUser(&app)

	if current, err := user.Current(); err == nil {
		if app.User.Uid != current.Uid {
			t.Error("User discovery did not use the current user")
		}
		if app.User.Uid != current.Uid {
			t.Error("User discovery did not use the current user")
		}
	}
}

func TestDiscoverProject(t *testing.T) {
	currentPath, _ := os.Getwd()
	rootPath := path.Dir(currentPath) // we know that root is one dir up
	projectPath := path.Join(rootPath, COACH_PROJECT_CONF_FOLDER)

	app := makeLocalApi()

	DiscoverProject(&app)

	if app.ExecPath != currentPath {
		t.Error("Project discovery uses incorrect current path")
	}
	if app.ProjectDoesntExist {
		t.Error("Project discovery indicates that no app exists, when one should have been found")
	}
	if getPath, err := app.Paths.Get("project"); err != nil {
		t.Error("Project discovery produced no path for the project")
	} else if getPath != projectPath {
		t.Error("Project discovery has the wrong Project path")
	}
}

func TestDiscoverEnvironmentPath(t *testing.T) {
	environment := "test"

	currentPath, _ := os.Getwd()
	rootPath := path.Dir(currentPath) // we know that root is one dir up
	envPath := path.Join(rootPath, COACH_PROJECT_CONF_FOLDER, COACH_ENVIRONMENT_CONF_SUBPATH, environment)

	app := makeLocalApi()

	app.ProjectRootPath = rootPath
	DiscoverEnvironment(&app, environment)

	if getPath, err := app.Paths.Get(environment); err != nil {
		t.Error("Environment discovery did not assign the env path as a config path")
	} else if getPath != envPath {
		t.Error("Environment discovery returned the wrong path for the test environment", envPath, getPath)
	}
}
