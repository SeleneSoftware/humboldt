package routetable

import (
	"fmt"
	pongo2 "github.com/flosch/pongo2/v4"
	"github.com/selenesoftware/humboldt/controller"
	"github.com/selenesoftware/humboldt/template"
	"github.com/yuin/gopher-lua"
	"net/http"
	"os"
	"strings"
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

func init() {
	fmt.Println("Initalizing the RouteTable")
	fileServer := http.FileServer(FileSystem{http.Dir("public")})
	http.Handle("/public/", http.StripPrefix(strings.TrimRight("/public/", "/"), fileServer))
}

// Not sure if this is still needed, but I'll leave it here JUST IN CASE
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
		// Everytime the route Lua file is changed,
		// this is run to make sure it is added to the server
		fmt.Println(rt.Route + " Added")
		RouteTable[rt.Name] = rt

		// The most important part of the system,
		// this is where each route runs the Lua
		// controller and sets up the output.
		http.HandleFunc(rt.Route, func(w http.ResponseWriter, r *http.Request) {
			controller.SetRequest(r)
			if err := L.DoFile("Controller/" + rt.Name + ".lua"); err != nil {
				// I would rather this throw a 502 or something of that sort.
				// But for now, this will do.
				// Don't judge me, this is still heavy development

				// Check to see if there is a static file before throwing everything away
				// if Exists("Public" + rt.Route) {
				// 	http.ServeFile(w, r, "public"+rt.Route)
				// 	return
				// }
				panic(err)
			}

			// Set the headers from the Lua files
			responseHeaders := controller.RetrieveHeader()
			for k, v := range responseHeaders {
				w.Header().Set(k, v)
			}

			// Template compilation and rendering
			tpl, _ := pongo2.FromFile(template.RetrieveTemplate())
			err := tpl.ExecuteWriter(template.RetrieveVariables(), w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	}

	return 1
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
