package forms

// This will contain all the helper functions for the controllers

import (
	// "encoding/json"
	"fmt"
	pongo2 "github.com/flosch/pongo2/v4"
	// "bitbucket.org/selenesoftware/humboldt/template"
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
	// "new": newForm,
	"render": render,
}

func render(L *lua.LState) int {
	f := L.ToTable(1)
	var form Form
	if err := gluamapper.Map(f, &form); err != nil {
		panic(err)
	}

	for _, field := range form.Elements {
		fmt.Println(field)
	}
	DefaultLoader := pongo2.MustNewLocalFileSystemLoader("")

	form2 := pongo2.NewSet("NewForm", DefaultLoader)

	formStr, _ := form2.RenderTemplateString("<form name='jason'>This will be a form</form>", pongo2.Context{})
	// fmt.Println(formStr)
	L.Push(lua.LString(formStr))

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
