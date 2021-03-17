module bitbucket.org/selenesoftware/humboldt

go 1.15

require (
	bitbucket.org/selenesoftware/humboldt/controller v0.0.0
	bitbucket.org/selenesoftware/humboldt/routetable v0.0.0
	bitbucket.org/selenesoftware/humboldt/template v0.0.0
	github.com/flosch/pongo2/v4 v4.0.2 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/radovskyb/watcher v1.0.7
	github.com/yuin/gopher-lua v0.0.0-20200816102855-ee81675732da
)

replace bitbucket.org/selenesoftware/humboldt/controller v0.0.0 => ./controller

replace bitbucket.org/selenesoftware/humboldt/forms v0.0.0 => ./forms

replace bitbucket.org/selenesoftware/humboldt/routetable v0.0.0 => ./routetable

replace bitbucket.org/selenesoftware/humboldt/template v0.0.0 => ./template

replace bitbucket.org/selenesoftware/humboldt/sessions v0.0.0 => ./sessions
