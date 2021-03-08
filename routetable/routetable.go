package routetable

import (
	"bitbucket.org/selenesoftware/humboldt/controller"
	"fmt"
	"github.com/yuin/gopher-lua"
	"net/http"
	// "sync"
)

var exports = map[string]lua.LGFunction{
	"route": route,
}

type Page struct {
	Title string
	Body  []byte
}

type Route struct {
	Name        string
	Route       string
	Method      string
	ForceSecure bool
}

var RouteTable = map[string]Route{}

func loadRoute(r string) Route {
	r = "/" + r
	for _, v := range RouteTable {
		if v.Route == r {
			fmt.Println(r)
			return v
		}
	}
	return Route{Name: "Name"}
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

func route(L *lua.LState) int {
	rt := Route{
		Name:        L.ToString(1),
		Route:       L.ToString(2),
		Method:      L.ToString(3),
		ForceSecure: L.ToBool(4),
	}

	if RouteTable[rt.Name].Name == "" {
		fmt.Println(rt.Route + " Added")
		RouteTable[rt.Name] = rt
		http.HandleFunc(rt.Route, func(w http.ResponseWriter, r *http.Request) {
			if err := L.DoFile("Controller/" + rt.Name + ".lua"); err != nil {
				// I would rather this throw a 502 or something of that sort.
				// But for now, this will do.
				// Don't judge me, this is still heavy development
				panic(err)
			}
			responseHeaders := controller.RetrieveHeader()
		})
	}

	return 1
}
