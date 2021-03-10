package template

import (
	// "fmt"
	pongo2 "github.com/flosch/pongo2/v4"
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

// Variables for the template
var tplVar = pongo2.Context{}

// Name of the template file
var fn string

// Add a variable for the template
func variable(L *lua.LState) int {
	k := L.ToString(1)
	v := L.ToString(2)

	tplVar[k] = v

	return 1
}

//Set the name of the template file to use
func setFile(L *lua.LState) int {
	fn = L.ToString(1)

	return 1
}

// Functions that move data around the server application

// Allows the Go code to grab the view definitions
func RetrieveTemplate() string {
	return "Views/" + fn + ".tpl"
}

// Allows the Go code to grab the variables for the templates
func RetrieveVariables() pongo2.Context {
	return tplVar
}
