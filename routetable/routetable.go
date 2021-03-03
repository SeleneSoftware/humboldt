package routetable

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	// "log"
	"net/http"
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
	for _, v := range RouteTable {
		if v.Route == r {
			return v
		}
	}
	return Route{Name: "Name"}
}

func Loader(L *lua.LState) int {

	// create the route table
	// RouteTable = map[string]Route{}
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	L.SetField(mod, "name", lua.LString("value"))

	// returns the module
	L.Push(mod)

	return 1
}

func route(L *lua.LState) int {
	r := Route{
		Name:        L.ToString(1),
		Route:       L.ToString(2),
		Method:      L.ToString(3),
		ForceSecure: L.ToBool(4),
	}

	if RouteTable[r.Name].Name == "" {
		fmt.Println(r.Route)
		RouteTable[r.Name] = r
		http.HandleFunc(r.Route, handler)
	}

	return 1
}

func handler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/view/"):]
	// p, err := loadRoute(title)
	// if err != nil {
	// 	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	// 	return
	// }
	rt := loadRoute(r.URL.Path[1:])
	fmt.Fprintf(w, "Hi there, I love %s!", rt.Name)
}
