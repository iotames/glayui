package web

import (
	"net/http"
)

type Routing struct {
	Path    string
	Methods []string
	handler func(w http.ResponseWriter, r *http.Request, dataFlow *DataFlow)
}

func GetDefaultRoutingList() []Routing {
	return []Routing{
		{Path: "/", handler: Home},
		{Path: "/hello", handler: Hello},
		{Path: "/debug", handler: Debug},
	}
}
