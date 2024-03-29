package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/iotames/glayui/component"
	"github.com/iotames/glayui/gtpl"
	"github.com/iotames/glayui/web"
)

func main() {
	tpl := gtpl.GetTpl()
	tpl.SetResourceDirPath("resource")
	s := web.NewEasyServer(":1598")
	s.SetData("dt111", "dtv126")
	fmt.Printf("-----You Can SET--ENV: USE_EMBED_FILE=true--toUseEmbedFile---\n")
	cpt := component.NewLayout("")
	s.AddHandler("GET", "/layout", func(ctx web.Context) {
		df := ctx.DataFlow
		dtkeys := df.GetDataKeys()
		err := ctx.DataFlow.SetData("startat", "hello")
		costime := time.Since(ctx.DataFlow.GetStartAt())
		fmt.Printf("-------cors(%s)--dtkeys(%+v)--cost(%.6fs)--startat(%v)--resetErr(%v)---\n", df.GetStr("CorsAllowOrigin"), dtkeys, costime.Seconds(), df.GetStartAt(), err)
		cpt.SetTitle("THIS is TITLE")
		cpt.SetContent("hello This is Content 99999999")
		cpt.Exec(ctx.Writer)
	})
	s.AddHandler("GET", "/table", func(ctx web.Context) {
		dtd := ctx.Server.GetData("dt111")
		fmt.Printf("--------dtd(%v)-----vvvv(%s)---\n", dtd, dtd.Value.(string))

		w := ctx.Writer
		// r := ctx.Request
		cpt.SetTitle("THIS is Data Table Title")
		table := component.NewTable("")
		dt := component.TableData{Name: "Tom", Age: 36}
		table.SetData(dt)
		var bf bytes.Buffer
		table.Exec(&bf)
		cpt.SetContent(bf.String())
		cpt.Exec(w)
	})
	s.ListenAndServe()
}

func HelloWordd() {
	fmt.Println("HELLO GLAYUI")
}
