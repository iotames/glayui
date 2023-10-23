package web

import (
	"fmt"
	"net/http"

	"github.com/iotames/glayui/gtpl"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	if !middleware(w, r) {
		return
	}
}

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PageNotFound"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	tpl := gtpl.GetTpl()
	type GameStatus struct {
		Name   string
		IsWin  bool
		Age    int
		Weight float64
	}
	data := GameStatus{Name: "Victor", Age: 399}
	fpath := `resource/tpl/hello.html`
	tpl = tpl.SetTplFile(fpath)

	err := tpl.SetData(data, w)
	if err != nil {
		fmt.Printf("----服务器错误(%v)---", err)
	}
}
