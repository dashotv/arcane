package app

import (
	"context"

	"github.com/LukeHagar/plexgo"

	"github.com/dashotv/fae"
)

func init() {
	initializers = append(initializers, setupPlex)
	starters = append(starters, startPlex)
}

func setupPlex(a *Application) error {
	a.Plex = plexgo.New(
		plexgo.WithSecurity(a.Config.PlexToken),
		plexgo.WithServerURL(a.Config.PlexURL),
		plexgo.WithXPlexClientIdentifier(a.Config.PlexIdentifier),
	)

	return nil
}

func startPlex(ctx context.Context, a *Application) error {
	res, err := a.Plex.Server.GetServerCapabilities(ctx)
	if err != nil {
		a.Log.Errorf("failed to get server capabilities: %v", err)
		return fae.Wrap(err, "failed to get server capabilities")
	}
	if res.Object == nil {
		a.Log.Error("failed to get server capabilities: response object is nil")
		return fae.Wrap(err, "failed to get server capabilities: response object is nil")
	}

	return nil
}
