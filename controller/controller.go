package controller

// This will contain all the helper functions for the controllers

import (
	// "encoding/json"
	// "fmt"
	"github.com/yuin/gopher-lua"
	"net/http"
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
	"setheader": setheader,
	"getheader": getheader,
}

func getheader(L *lua.LState) int {
	v := L.ToString(1)
	// In time, I will have this by default push the whole map
	// into a table, but I just don't feel like it right now
	// if v == "" {
	// 	L.Push(req.Header)
	// 	return 1
	// }
	// fmt.Println(v)
	L.Push(lua.LString(req.Header.Get(v)))
	return 1 // Notify that we pushed one value to the stack
}

// Set a header for the response in the header
func setheader(L *lua.LState) int {
	k := L.ToString(1)
	v := L.ToString(2)
	h[k] = v
	return 1
}

// Functions to move data to other parts of the server

var req *http.Request

func SetRequest(r *http.Request) {
	req = r
}

// All the header variables
func RetrieveHeader() map[string]string {
	return h
}
