package forms

// This will contain all the helper functions for the controllers

import (
	// "encoding/json"
	"fmt"
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

type Form struct {
	Name     string
	Elements []FormElement
}

type FormElement struct {
	Name  string
	Type  string
	Value string
}

var exports = map[string]lua.LGFunction{
	"new": newForm,
}

func newForm(L *lua.LState) int {
	f := L.ToTable(1)
	var form Form
	if err := gluamapper.Map(f, &form); err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", f)
	fmt.Println(form)
	return 1
}

func handleRequest(L *lua.LState) int {
	return 1
}

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	L.SetField(mod, "name", lua.LString("value"))

	// returns the module
	L.Push(mod)
	return 1
}
