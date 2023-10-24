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
	fpath := `resource/tpl/demo/table1.html`
	err := tpl.SetDataByTplFile(fpath, data, w)
	if err != nil {
		fmt.Printf("----服务器错误(%v)---", err)
	}
}

func debug(w http.ResponseWriter, r *http.Request) {
	tpl := gtpl.GetTpl()
	fpath := `resource/tpl/demo/layout.html`
	err := tpl.SetDataByTplFile(fpath, "hello layout", w)
	if err != nil {
		fmt.Printf("----服务器错误(%v)---", err)
	}
}
