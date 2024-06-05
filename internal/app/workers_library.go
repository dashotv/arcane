package app

import (
	"context"
	"os"
	"path/filepath"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
)

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

	err = filepath.Walk(lib.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := a.DB.FileCreateOrUpdate(path)
		if err != nil {
			return fae.Wrap(err, "creating or updating file")
		}
		f.Size = info.Size()
		f.ModifiedAt = info.ModTime().Unix()

		return nil
	})
	if err != nil {
		return fae.Wrap(err, "walking")
	}

	return nil
}
