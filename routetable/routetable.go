package routetable

import (
	"fmt"
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
	"route": route,
}

// r.route(uri,forceHttps)
func route(L *lua.LState) int {
	// Name - This will correlate with the controller file
	nm := L.ToString(1)
	// Route - The URI
	rt := L.ToString(2)
	// Method - GET or POST, GET by default
	mt := L.ToString(3)
	// Force Secure - Do you want to run this as TLS?  Default is false
	fs := L.ToBool(4)

	fmt.Println(nm, rt, mt, fs)

	return 1
}
