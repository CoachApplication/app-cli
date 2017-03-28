package api

import (
	"os"
	"testing"

	handler_local "github.com/CoachApplication/handler-local"
)

func makeLocalSettings() handler_local.Settings {
	workingDir, _ := os.Getwd()
	return handler_local.Settings{
		ExecPath: workingDir,
		Paths:    *handler_local.NewSettingScopePaths(),
	}
}

func TestMakeLocalApp(t *testing.T) {
	settings := makeLocalSettings()

	if _, err := MakeLocalApp(settings); err != nil {
		t.Error("Error occured making an app object", err.Error())
	}

}
