package main

import (
	"bitbucket.org/selenesoftware/humboldt/controller"
	"bitbucket.org/selenesoftware/humboldt/forms"
	"bitbucket.org/selenesoftware/humboldt/routetable"
	"bitbucket.org/selenesoftware/humboldt/template"
	// "context"
	"fmt"
	"github.com/radovskyb/watcher"
	"github.com/yuin/gopher-lua"
	"log"
	"regexp"
	"sync"
	"time"
)

func main() {

	httpServerExitDone := &sync.WaitGroup{}

	httpServerExitDone.Add(1)
	startHttpServer(httpServerExitDone)

	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("controller", controller.Loader)
	L.PreloadModule("template", template.Loader)
	L.PreloadModule("route", routetable.Loader)
	L.PreloadModule("forms", forms.Loader)

	w := watcher.New()
	w.FilterOps(watcher.Rename, watcher.Move, watcher.Write)
	if err := w.Add("Config"); err != nil {
		log.Fatalln(err)
	}

	w.FilterOps(watcher.Create)

	// Only files that match the regular expression during file listings
	// will be watched.
	r := regexp.MustCompile("^Routing.lua$")
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event)
				// if err := srv.Shutdown(context.TODO()); err != nil {
				// 	panic(err)
				// }

				if err := L.DoFile("Config/Routing.lua"); err != nil {
					panic(err)
				}

				// httpServerExitDone = &sync.WaitGroup{}
				//
				// httpServerExitDone.Add(1)
				// srv = startHttpServer(httpServerExitDone)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				fmt.Println("Closed, foo!")
				return
			}
		}
	}()

	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	go func() {
		w.Wait()
		w.TriggerEvent(watcher.Write, nil)
	}()

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
