package api

import (
	"context"
	coach_api "github.com/CoachApplication/api"
	coach_base "github.com/CoachApplication/base"
	handler_local "github.com/CoachApplication/handler-local"
)

func MakeLocalApp(ctx context.Context, settings handler_local.Settings) (coach_api.Application, error) {

	if settings.ProjectDoesntExist {
		app := coach_base.NewApplication(nil)
		// start off with a local config handler [bare necessity for further configuration]
		app.AddBuilder(handler_local.NewBuilder(ctx, settings))
		// Activate just the project handler, which can be used to generate a new project
		app.Activate("local.standard", []string{"project"}, nil)

		return app.Application(), nil
	} else {
		app := coach_base.NewApplication(nil)

		// start off with a local config handler [bare necessity for further configuration]
		app.AddBuilder(handler_local.NewBuilder(ctx, settings))
		app.Activate("local.standard", []string{"config"}, nil)

		return app.Application(), nil
	}
}
