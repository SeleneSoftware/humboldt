module github.com/SeleneSoftware/humboldt

go 1.15

require (
	github.com/SeleneSoftware/humboldt/controller v0.0.0
	github.com/SeleneSoftware/humboldt/routetable v0.0.0
	github.com/SeleneSoftware/humboldt/template v0.0.0
	github.com/flosch/pongo2/v4 v4.0.2 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/radovskyb/watcher v1.0.7
	github.com/yuin/gopher-lua v0.0.0-20200816102855-ee81675732da
)

replace github.com/SeleneSoftware/humboldt/controller v0.0.0 => ./controller

replace github.com/SeleneSoftware/humboldt/forms v0.0.0 => ./forms

replace github.com/SeleneSoftware/humboldt/routetable v0.0.0 => ./routetable

replace github.com/SeleneSoftware/humboldt/template v0.0.0 => ./template

replace github.com/SeleneSoftware/humboldt/sessions v0.0.0 => ./sessions
