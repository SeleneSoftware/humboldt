package template

import (
	// "fmt"
	"github.com/yuin/gopher-lua"
)

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
	"variable": variable,
}

func variable(L *lua.LState) int {
	// k := L.ToString(1)
	// v := L.ToString(2)

	return 1
}
