package app

import (
	"context"
	"os"
	"path/filepath"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
)

type LibraryFilesAll struct {
	minion.WorkerDefaults[*LibraryFilesAll]
}

func (j *LibraryFilesAll) Kind() string { return "library_files_all" }
func (j *LibraryFilesAll) Work(ctx context.Context, job *minion.Job[*LibraryFilesAll]) error {
	//args := job.Args
	a := ContextApp(ctx)
	list, err := a.DB.Library.Query().Run()
	if err != nil {
		return fae.Wrap(err, "querying")
	}

	for _, lib := range list {
		if lib.Name == "ecchi" {
			if err := a.Workers.Enqueue(&LibraryFiles{ID: lib.ID.Hex()}); err != nil {
				return fae.Wrap(err, "enqueue")
			}
		}
	}

	return nil
}

type LibraryFiles struct {
	minion.WorkerDefaults[*LibraryFiles]
	ID string `bson:"id" json:"id"`
}

func (j *LibraryFiles) Kind() string { return "library_files" }
func (j *LibraryFiles) Work(ctx context.Context, job *minion.Job[*LibraryFiles]) error {
	a := ContextApp(ctx)
	id := job.Args.ID

	lib, err := a.DB.LibraryGet(id)
	if err != nil {
		return fae.Wrap(err, "getting library")
	}

	l := a.Log.Named("library_files").With("dir", lib.Path)
	l.Debug("walking")

	err = filepath.Walk(lib.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		l.Debugf("path: %s", path)

		if info.IsDir() {
			return nil
		}

		f, err := a.DB.FileCreateOrUpdate(path)
		if err != nil {
			return fae.Wrap(err, "creating or updating file")
		}
		f.Size = info.Size()
		f.ModifiedAt = info.ModTime().Unix()

		if err := a.DB.File.Save(f); err != nil {
			return fae.Wrap(err, "saving file")
		}

		return nil
	})
	if err != nil {
		return fae.Wrap(err, "walking")
	}

	return nil
}
