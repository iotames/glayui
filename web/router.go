package web

import (
	"net/http"
	"strings"

	"github.com/iotames/glayui/web/ctl"
)

type Routing struct {
	Path    string
	Methods []string
	handler func(w http.ResponseWriter, r *http.Request)
}

func GetRoutingList() []Routing {
	return []Routing{
		{Path: "/", handler: ctl.Home},
		{Path: "/hello", handler: hello},
	}
}

func MiddleRouter(w http.ResponseWriter, r *http.Request) (subNext bool, mainNext bool) {
	routings := GetRoutingList()
	isMatch := false
	rpath := r.URL.Path
	rmethod := r.Method

	for _, rt := range routings {
		if rt.Path == rpath {
			if len(rt.Methods) == 0 {
				isMatch = true
			} else {
				for _, m := range rt.Methods {
					// strings.ToUpper(m) == strings.ToUpper(rmethod)
					if strings.EqualFold(m, rmethod) {
						isMatch = true
					}
				}
			}
			if isMatch {
				rt.handler(w, r)
				break
			}
		}
	}

	if !isMatch {
		PageNotFound(w, r)
	}

	return true, true
}
