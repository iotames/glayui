package gtpl

import (
	"bytes"
	"fmt"

	// "fmt"
	"io"
	// "os"
	"path/filepath"
	"text/template"
)

const GTPL_NAME = "GTPL_NAME"

const GTPL_DELIM_LEFT = `<{%`
const GTPL_DELIM_RIGHT = `%}>`

type Gtpl struct {
	t                     *template.Template
	funcMap               template.FuncMap
	delimLeft, delimRight string
	tplText               string
	tplFilepath           string
	buff                  bytes.Buffer
}

func NewTpl(leftDelim, rightDelim string) *Gtpl {
	if leftDelim == "" || rightDelim == "" {
		panic("leftDelim or rightDelim could not be empty")
	}
	return &Gtpl{delimLeft: leftDelim, delimRight: rightDelim}
}

func (g *Gtpl) SetTemplate(t *template.Template) *Gtpl {
	g.t = t
	return g
}

func (g *Gtpl) SetTplText(tpltext string) *Gtpl {
	g.tplText = tpltext
	return g
}

func (g *Gtpl) SetTplFile(fpath string) *Gtpl {
	g.tplFilepath = fpath
	return g
}

func (g *Gtpl) SetData(data any, wr io.Writer) error {
	if g.tplFilepath != "" {
		if g.t == nil {
			g.t = createTpl(filepath.Base(g.tplFilepath), g.funcMap).Delims(g.delimLeft, g.delimRight)
		}
		t, err := g.t.ParseFiles(g.tplFilepath)
		if err != nil {
			return err
		}
		result := t.Execute(&g.buff, data)
		if wr != nil {
			return t.Execute(wr, data)
		}
		return result
	}

	if g.t == nil {
		g.t = createTpl(GTPL_NAME, g.funcMap).Delims(g.delimLeft, g.delimRight)
	}
	t, err := g.t.Parse(g.tplText)
	if err != nil {
		return err
	}
	result := t.Execute(&g.buff, data)
	if wr != nil {
		return t.Execute(wr, data)
	}
	return result
}

func (g Gtpl) GetBuffer() bytes.Buffer {
	return g.buff
}

func (g Gtpl) Bytes() []byte {
	return g.buff.Bytes()
}

func (g Gtpl) String() string {
	return g.buff.String()
}

func (g *Gtpl) SetFuncMap(mp template.FuncMap) *Gtpl {
	g.funcMap = mp
	return g
}

func (g *Gtpl) AddFunc(k string, v any) error {
	if g.funcMap == nil {
		g.funcMap = template.FuncMap{k: v}
		return nil
	}
	_, ok := g.funcMap[k]
	if ok {
		return fmt.Errorf("the func of key(%s) has exist", k)
	}
	g.funcMap[k] = v
	return nil
}

func createTpl(name string, funcmp template.FuncMap) *template.Template {
	tpl := template.New(name)
	if funcmp != nil {
		tpl = tpl.Funcs(funcmp)
	}
	return tpl
}

var gtpl *Gtpl

func GetTpl() *Gtpl {
	if gtpl == nil {
		gtpl = NewTpl(GTPL_DELIM_LEFT, GTPL_DELIM_RIGHT)
		gtpl.SetFuncMap(getDefaultFuncMap())
	}
	return gtpl
}

func getDefaultFuncMap() template.FuncMap {
	return template.FuncMap{
		// "getDataTypeForJS":  getDataTypeForJS,
		// "getFormFieldsHtml": getFormFieldsHtml,
		// "toObjStr":          database.TableColToObj,
		// "dbtype":            dbtype,
		// "dbdefault":         dbdefault,
		"gotype": gotype,
	}
}
