package forms

// This will contain all the helper functions for the controllers

import (
	// "encoding/json"
	// "fmt"
	"github.com/yuin/gopher-lua"
)

var h = map[string]string{}

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	L.SetField(mod, "name", lua.LString("value"))

	// returns the module
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	// "request":   request,
	"formbuilder": formbuilder,
	"renderform":  renderform,
}

func formbuilder(L *lua.LState) int {
	return 1
}
