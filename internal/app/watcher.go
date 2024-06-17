package app

// import (
// 	"context"
// 	"io/fs"
// 	"log"
// 	"path/filepath"
//
// 	"github.com/fsnotify/fsnotify"
//
// 	"github.com/dashotv/fae"
// )
//
// func init() {
// 	initializers = append(initializers, setupWatcher)
// 	starters = append(starters, startWatcher)
// }
//
// func setupWatcher(a *Application) error {
// 	a.Watcher = &Watcher{
// 		Directories: a.Config.WatcherDirectories,
// 	}
//
// 	return nil
// }
//
// func startWatcher(ctx context.Context, a *Application) error {
// 	if err := a.Watcher.Watch(ctx); err != nil {
// 		a.Log.Errorf("failed to watch directories: %v", err)
// 		return fae.Wrap(err, "failed to watch directories")
// 	}
//
// 	return nil
// }
//
// type Watcher struct {
// 	Directories []string
// }
//
// func (w *Watcher) Watch(ctx context.Context) error {
// 	a := ContextApp(ctx)
// 	if a == nil {
// 		return fae.New("app context not found")
// 	}
//
// 	l := a.Log.Named("watcher")
//
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()
//
// 	for _, dir := range w.Directories {
// 		dir := dir
// 		go func() {
// 			defer TickTock("watcher")
// 			err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
// 				if err != nil {
// 					return err
// 				}
// 				if d.IsDir() {
// 					return watcher.Add(path)
// 				}
// 				return nil
// 			})
// 			if err != nil {
// 				l.Errorf("failed to walk directory: %s: %v", dir, err)
// 				return
// 			}
// 			l.Debugf("watching directory: %s", dir)
// 		}()
// 	}
//
// 	l.Debugf("starting watcher: (%v)", w.Directories)
// 	go func() {
// 		// defer watcher.Close()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case event := <-watcher.Events:
// 				l.Debugf("event: %#v\n", event)
// 			case err := <-watcher.Errors:
// 				l.Debugf("error: %v", err)
// 			}
// 		}
// 	}()
//
// 	return nil
// }
