package main

import (
	"bitbucket.org/selenesoftware/humboldt/controller"
	"bitbucket.org/selenesoftware/humboldt/template"
	"github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("controller", controller.Loader)
	L.PreloadModule("template", template.Loader)

	if err := L.DoFile("test.lua"); err != nil {
		panic(err)
	}
}
