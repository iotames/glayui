package web

import (
	"net/http"

	"github.com/iotames/glayui/web/response"
)

func ResponseNotFound(w http.ResponseWriter, r *http.Request) {
	dt := response.NewApiDataNotFound()
	w.Write(dt.Bytes())
}
