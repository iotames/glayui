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

	data := GameStatus{Name: "Victor21", Age: 321}

	tpl := gtpl.GetTpl()
	f, err := os.OpenFile("runtime/debug.html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	fpath := `tpl/demo/table1.html`
	// fullpath, fss, err := resource.GetResourceFile(fpath)
	// t.Logf("---fullpath(%s)2----fss(%+v)--err(%v)----\n", fullpath, fss, err)

	// err = tpl.SetDataByTplFile(fpath, data, f)
	// err = tpl.SetDataByTplFS(fpath, fss, data, f)
	err = tpl.SetDataFromResource(fpath, data, f)
	if err != nil {
		panic(err)
	}
	t.Logf("------getString(%s)---\n", tpl.String())
}

func TestMain(t *testing.T) {
	t.Logf("---------MaxHeaderBytes(%d)------\n", 1<<20)
}
