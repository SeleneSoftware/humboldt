package controller

// This will contain all the helper functions for the controllers

import (
	// "fmt"
	"github.com/yuin/gopher-lua"
)

var M = map[string]string{}

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
	"request":   request,
	"setheader": setheader,
}

func request(L *lua.LState) int {
	i := L.ToInt(1)          // get first (1) function argument and convert to int
	ln := lua.LNumber(i * i) // make calculation and cast to LNumber
	L.Push(ln)               // Push it to the stack
	return 1                 // Notify that we pushed one value to the stack
}

func setheader(L *lua.LState) int {
	k := L.ToString(1)
	v := L.ToString(2)
	M[k] = v
	return 1
}

func RetrieveHeader() map[string]string {
	return M
}
