package api

import (
	"os"
	"testing"

	"context"
	handler_local "github.com/CoachApplication/handler-local"
	"time"
)

func makeLocalSettings() handler_local.Settings {
	workingDir, _ := os.Getwd()
	return handler_local.Settings{
		ExecPath: workingDir,
		Paths:    *handler_local.NewSettingScopePaths(),
	}
}

func TestMakeLocalApp(t *testing.T) {
	dur, _ := time.ParseDuration("5s")
	ctx, _ := context.WithTimeout(context.Background(), dur)
	settings := makeLocalSettings()

	if _, err := MakeLocalApp(ctx, settings); err != nil {
		t.Error("Error occured making an app object", err.Error())
	}

}
