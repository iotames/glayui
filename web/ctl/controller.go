package ctl

import (
	"fmt"
	"net/http"

	"github.com/iotames/glayui/gtpl"
	"github.com/iotames/glayui/web/response"
)

func Home(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	for k, v := range query {
		fmt.Println("----url--query---", k, v)
		fmt.Println("--------------------", r.URL.Host, r.URL.Path)
	}

	tpl := gtpl.GetTpl()
	fpath := `tpl/demo/layout.html`
	// err := tpl.SetDataFromResource(fpath, "hello layout", w)
	err := tpl.SetDataByTplFile(fpath, "hello layout", w)
	if err != nil {
		resp := response.NewApiDataServerError(err.Error())
		w.Write(resp.Bytes())
		fmt.Printf("----服务器错误(%v)---\n", resp.String())
	}
}
