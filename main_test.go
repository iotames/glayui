package main

import (
	"os"
	"testing"

	// "text/template"
	"github.com/iotames/glayui/gtpl"
)

func TestTpl(t *testing.T) {

	type GameStatus struct {
		Name   string
		IsWin  bool
		Age    int
		Weight float64
	}

	data := GameStatus{Name: "Harvey", Age: 23}
	fpath := `resource/tpl/hello.html`
	tpl := gtpl.NewTpl("hello.html", gtpl.GTPL_DELIM_LEFT, gtpl.GTPL_DELIM_RIGHT)
	tpl = tpl.SetTplFile(fpath)
	// // tpl = tpl.SetTplText(`fhlaksdjflakdjf---<{% .Name %}> : <{% .Age %}>--fadsfadfadfa------`)
	err := tpl.SetData(data, nil)
	if err != nil {
		panic(err)
	}
	bf := tpl.GetBuffer()

	f, err := os.OpenFile("runtime/hello.html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	// tt, err := template.ParseFiles(fpath)
	// if err != nil {
	// 	panic(err)
	// }
	// err = tt.Execute(f, data)

	_, err = f.Write(bf.Bytes())
	if err != nil {
		panic(err)
	}
}
