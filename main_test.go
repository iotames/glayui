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
	fpath := `resource/tpl/demo/table1.html`
	tpl := gtpl.GetTpl()
	f, err := os.OpenFile("runtime/debug.html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	err = tpl.SetDataByTplFile(fpath, data, f) //SetData(data, f)
	if err != nil {
		panic(err)
	}
	t.Logf("------getString(%s)---\n", tpl.String())
}

func TestMain(t *testing.T) {
	t.Logf("---------MaxHeaderBytes(%d)------\n", 1<<20)
}
