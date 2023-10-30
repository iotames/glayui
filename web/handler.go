package web

import (
	"fmt"
	"net/http"

	"github.com/iotames/glayui/gtpl"
)

func ResponseNotFound(w http.ResponseWriter, r *http.Request) {
	dt := NewApiDataNotFound()
	w.Write(dt.Bytes())
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
	fpath := `tpl/demo/table1.html`
	err := tpl.SetDataFromResource(fpath, data, w)
	if err != nil {
		resp := NewApiDataServerError(err.Error())
		w.Write(resp.Bytes())
		fmt.Printf("----服务器错误(%v)---\n", resp.String())
	}
}

func debug(w http.ResponseWriter, r *http.Request) {
	tpl := gtpl.GetTpl()
	fpath := `tpl/demo/layout.html`
	err := tpl.SetDataFromResource(fpath, "hello layout", w)
	if err != nil {
		resp := NewApiDataServerError(err.Error())
		w.Write(resp.Bytes())
		fmt.Printf("----服务器错误(%v)---\n", resp.String())
	}
}
