package app

import (
	"context"
	"strconv"

	"github.com/LukeHagar/plexgo/models/operations"
	"github.com/davecgh/go-spew/spew"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
)

type PlexFiles struct {
	minion.WorkerDefaults[*PlexFiles]
	ID string `bson:"id" json:"id"`
}

func (j *PlexFiles) Kind() string { return "plex_files" }
func (j *PlexFiles) Work(ctx context.Context, job *minion.Job[*PlexFiles]) error {
	a := ContextApp(ctx)
	l := a.Workers.Log.Named("plex_files")
	id := job.Args.ID

	lib, err := a.DB.LibraryGet(id)
	if err != nil {
		return fae.Wrap(err, "getting library")
	}

	section := -1

	list, err := a.Plex.Library.GetLibraries(ctx)
	if err != nil {
		return fae.Wrap(err, "getting libraries")
	}
	for _, plib := range list.GetObject().MediaContainer.Directory {
		for _, loc := range plib.Location {
			if *loc.Path == lib.Path {
				l.Debugf("found section %s", *plib.Key)
				num, err := strconv.Atoi(*plib.Key)
				if err != nil {
					l.Warn("section not a number")
					continue
				}
				section = num
			}
		}
	}

	if section == -1 {
		l.Warn("section not found")
		return nil
	}

	resp, err := a.Plex.Library.GetLibraryItems(ctx, section, operations.TagAll)
	if err != nil {
		return fae.Wrap(err, "getting library items")
	}
	for _, item := range resp.GetObject().MediaContainer.Metadata {
		l.Debugf("item %s", *item.Title)
		rk, err := strconv.ParseFloat(*item.RatingKey, 64)
		if err != nil {
			l.Warn("rating key not a number")
			continue
		}
		resp, err := a.Plex.Library.GetMetadataChildren(ctx, rk)
		if err != nil {
			return fae.Wrap(err, "getting library items")
		}
		for _, child := range resp.GetObject().MediaContainer.Metadata {
			l.Debugf("child %s", *child.Title)
			rk2, err := strconv.ParseFloat(*child.RatingKey, 64)
			if err != nil {
				l.Warn("rating key not a number")
				continue
			}
			resp, err := a.Plex.Library.GetMetadataChildren(ctx, rk2)
			if err != nil {
				return fae.Wrap(err, "getting library items")
			}
			for _, child2 := range resp.GetObject().MediaContainer.Metadata {
				l.Debugf("child2 %s", spew.Sdump(child2))
				rk3, err := strconv.ParseFloat(*child2.RatingKey, 64)
				if err != nil {
					l.Warn("rating key not a number")
					continue
				}
				resp, err := a.Plex.Library.GetMetadata(ctx, rk3)
				if err != nil {
					return fae.Wrap(err, "getting library items")
				}
				l.Debugf("metadata %s", spew.Sdump(resp.GetObject().MediaContainer.Metadata))
			}
		}
	}

	return nil
}
