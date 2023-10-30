package main

import (
	"bytes"
	"fmt"

	"github.com/iotames/glayui/component"
	"github.com/iotames/glayui/gtpl"
	"github.com/iotames/glayui/web"
	"github.com/iotames/glayui/web/response"
)

func main() {
	s := web.NewEasyServer(":1598")
	cpt := component.NewLayout("")
	s.AddHandler("GET", "/layout", func(ctx web.Context) {
		cpt.SetTitle("THIS is TITLE")
		cpt.SetContent("hello This is Content 99999999")
		cpt.Handler(ctx.Writer, ctx.Request)
	})
	s.AddHandler("GET", "/table", func(ctx web.Context) {
		w := ctx.Writer
		r := ctx.Request
		cpt.SetTitle("THIS is Data Table")
		tpl := gtpl.GetTpl()
		type GameStatus struct {
			Name   string
			IsWin  bool
			Age    int
			Weight float64
		}
		data := GameStatus{Name: "Victor", Age: 399}
		fpath := `tpl/table.html`
		var bf bytes.Buffer
		err := tpl.SetDataByTplFile(fpath, data, &bf)
		if err != nil {
			resp := response.NewApiDataServerError(err.Error())
			w.Write(resp.Bytes())
			fmt.Printf("----服务器错误(%v)---\n", resp.String())
			return
		}
		cpt.SetContent(bf.String())
		cpt.Handler(w, r)
	})
	s.ListenAndServe()
}

func HelloWordd() {
	fmt.Println("HELLO GLAYUI")
}
