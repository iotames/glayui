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

	data := GameStatus{Name: "Victor", Age: 399}
	fpath := `resource/tpl/hello.html`
	tpl := gtpl.GetTpl()
	tpl = tpl.SetTplFile(fpath)
	// // tpl = tpl.SetTplText(`fhlaksdjflakdjf---<{% .Name %}> : <{% .Age %}>--fadsfadfadfa------`)

	f, err := os.OpenFile("runtime/hello.html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	err = tpl.SetData(data, f)
	if err != nil {
		panic(err)
	}
	t.Logf("------getString(%s)---\n", tpl.String())
}

func TestMain(t *testing.T) {
	t.Logf("---------MaxHeaderBytes(%d)------\n", 1<<20)
}
