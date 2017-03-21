package main

import (
	"os"
	"os/user"
	"path"

	handler_local "github.com/james-nesbitt/coach-local"
)

const (
	COACH_PROJECT_CONF_FOLDER      = ".coach"       // If the project has existing setitngs, they will be in this subfolder, somewhere up the file tree.
	COACH_USER_CONF_SUBPATH        = "coach"        // If the user has user-scope config, they will be in this subfolder
	COACH_ENVIRONMENT_CONF_SUBPATH = "environments" // a subpath in .coach for per environment settings
)

/**
 * Some utility functions for discovering scope paths
 *
 * Scope is usually defined as being in three layers:
 *   - core : stuff that is hardcoded into the API/CLI
 *   - user : configuration for the user across all projects
 *   - project : stuff in the project root configuration path
 *
 *   - environment : specific project environment settings,
 *        usually kept in a sub-path of the project settings.
 *        and keyed using a CLI global flag.
 *
 * @NOTE the difference OSes have different approaches for
 *  determining user settings, usually just based on different
 *  paths for the configuration.
 */

/**
 * DiscoverUser who is the current user and enable some user paths
 *
 * @NOTE we have to play some games for different OSes here
 *
 * dependending on OS, determine if the user has any settings
 * if so, add a conf path for them.
 */
func DiscoverUser(settings *handler_local.Settings) error {

	if current, err := user.Current(); err == nil {
		settings.User = *current
	} else {
		settings.User = user.User{
			Username: "anonymous",
			Name:     "Anonymous",
		}
	}

	if settings.User.HomeDir == "" {
		/**
		 * There is an issue in some envs (CoreOSX) where
		 * the golang user library does not set a current user
		 * object properly, so we fall back to checking
		 * an ENV variable.
		 */
		settings.User.HomeDir = os.Getenv("HOME")
	}
	if settings.User.HomeDir != "" {
		homeDir := settings.User.HomeDir

		// This is a common, but not very good user config path for *Nix OSes
		homeConfDir := path.Join(homeDir, "."+COACH_PROJECT_CONF_FOLDER) // if in the home folder, add a "."

		if _, err := os.Stat(path.Join(homeDir, "Library")); err == nil {
			// OSX
			homeConfDir = path.Join(homeDir, "Library", COACH_USER_CONF_SUBPATH)
		} else if _, err := os.Stat(path.Join(homeDir, ".config")); err == nil {
			// Good *Nix/BSD
			homeConfDir = path.Join(homeDir, ".config", COACH_USER_CONF_SUBPATH)
		}

		/**
		 * @TODO does anybody care about any other OS? (logi?)
		 */

		/**
		 * Set up some frequently used paths
		 */
		settings.Paths.Set("user", homeConfDir)
	}

	return nil
}

/**
 * Discover project paths
 *
 * Recursively navigate up the file path until we discover a folder that
 * has the key configuration subfolder in it.  That path is marked as the
 * application root, and the subfolder is marked as a conf path
 */
func DiscoverProject(settings *handler_local.Settings) error {
	workingDir := settings.ExecPath

	homeDir := ""
	if &settings.User != nil {
		homeDir = settings.User.HomeDir
	}

	projectRootDirectory := workingDir
	_, err := os.Stat(path.Join(projectRootDirectory, COACH_PROJECT_CONF_FOLDER))
RootSearch:
	for err != nil {
		projectRootDirectory = path.Dir(projectRootDirectory)
		if projectRootDirectory == homeDir || projectRootDirectory == "." || projectRootDirectory == "/" {
			// Could not find a project folder
			projectRootDirectory = workingDir
			settings.ProjectDoesntExist = true
			break RootSearch
		}
		_, err = os.Stat(path.Join(projectRootDirectory, COACH_PROJECT_CONF_FOLDER))
	}

	/**
	 * Set up some frequently used paths
	 */
	settings.ProjectRootPath = projectRootDirectory
	settings.Paths.Set("project", path.Join(projectRootDirectory, COACH_PROJECT_CONF_FOLDER))

	return err
}

/**
 * Discover environment path for a specific environment
 *
 */
func DiscoverEnvironment(settings *handler_local.Settings, environment string) error {

	/**
	 * @TODO actually check to see if the path exists, so that we can warn if it doesn't?
	 */

	// add the environment sub path of the main project conf directory, as a conf path
	settings.Paths.Set(environment, path.Join(settings.ProjectRootPath, COACH_PROJECT_CONF_FOLDER, COACH_ENVIRONMENT_CONF_SUBPATH, environment))

	return nil
}
