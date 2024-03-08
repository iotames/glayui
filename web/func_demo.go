package web

import (
	"fmt"
	"net/http"

	"github.com/iotames/glayui/gtpl"
	"github.com/iotames/glayui/web/response"
)

func Home(w http.ResponseWriter, r *http.Request, dataFlow *DataFlow) {
	query := r.URL.Query()
	for k, v := range query {
		fmt.Println("----url--query---", k, v)
		fmt.Println("--------------------", r.URL.Host, r.URL.Path)
	}

	tpl := gtpl.GetTpl()
	fpath := `tpl/demo/layout1.html`
	// err := tpl.SetDataFromResource(fpath, "hello layout", w)
	err := tpl.SetDataByTplFile(fpath, "hello layout", w)
	if err != nil {
		resp := response.NewApiDataServerError(err.Error())
		w.Write(resp.Bytes())
		fmt.Printf("----服务器错误(%v)---\n", resp.String())
	}
}

func Hello(w http.ResponseWriter, r *http.Request, dataFlow *DataFlow) {
	tpl := gtpl.GetTpl()
	type GameStatus struct {
		Name   string
		IsWin  bool
		Age    int
		Weight float64
	}
	data := GameStatus{Name: "Victor", Age: 399}
	fpath := `tpl/demo/table1.html`
	err := tpl.SetDataByTplFile(fpath, data, w)
	if err != nil {
		resp := response.NewApiDataServerError(err.Error())
		w.Write(resp.Bytes())
		fmt.Printf("----服务器错误(%v)---\n", resp.String())
	}
}

func Debug(w http.ResponseWriter, r *http.Request, dataFlow *DataFlow) {
	tpl := gtpl.GetTpl()
	fpath := `tpl/demo/layout1.html`
	err := tpl.SetDataFromResource(fpath, "hello layout", w)
	if err != nil {
		resp := response.NewApiDataServerError(err.Error())
		w.Write(resp.Bytes())
		fmt.Printf("----服务器错误(%v)---\n", resp.String())
	}
}
