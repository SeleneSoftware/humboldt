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
	"setFile":  setFile,
}

var tmpVar = map[string]string{}

var fn string

func variable(L *lua.LState) int {
	k := L.ToString(1)
	v := L.ToString(2)

	tmpVar[k] = v

	return 1
}

func setFile(L *lua.LState) int {
	fn = L.ToString(1)

	return 1
}

func RetrieveTemplate() string {
	return "Views/" + fn + ".tpl"
}
