package web

import (
	"net/http"

	"github.com/iotames/glayui/web/ctl"
)

type Routing struct {
	Path    string
	Methods []string
	handler func(w http.ResponseWriter, r *http.Request)
}

func GetDefaultRoutingList() []Routing {
	return []Routing{
		{Path: "/", handler: ctl.Home},
		{Path: "/hello", handler: ctl.Hello},
		{Path: "/debug", handler: ctl.Debug},
	}
}
