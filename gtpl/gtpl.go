package gtpl

import (
	"bytes"
	// "fmt"
	"io"
	// "os"
	// "path/filepath"
	"text/template"
)

const GTPL_NAME = "GTPL_NAME"

const GTPL_DELIM_LEFT = `<{%`
const GTPL_DELIM_RIGHT = `%}>`

type Gtpl struct {
	tplName               string
	t                     template.Template
	delimLeft, delimRight string

	tplText     string
	tplFilepath string
	buff        bytes.Buffer
}

func NewTpl(tplName string, leftDelim, rightDelim string) *Gtpl {
	if leftDelim == "" || rightDelim == "" {
		panic("leftDelim or rightDelim could not be empty")
	}
	if tplName == "" {
		tplName = GTPL_NAME
	}
	t := createTpl(tplName).Delims(leftDelim, rightDelim)
	return &Gtpl{tplName: tplName, t: *t, delimLeft: leftDelim, delimRight: rightDelim}
}

func (g *Gtpl) SetTplText(tpltext string) *Gtpl {
	g.tplText = tpltext
	return g
}

func (g *Gtpl) SetTplFile(filepath string) *Gtpl {
	g.tplFilepath = filepath
	return g
}

func (g *Gtpl) SetData(data any, wr io.Writer) error {
	if g.tplFilepath != "" {
		t, err := g.t.ParseFiles(g.tplFilepath)
		if err != nil {
			return err
		}
		if wr != nil {
			return t.Execute(wr, data)
		}
		return t.Execute(&g.buff, data)
	}
	t, err := g.t.Parse(g.tplText)
	if err != nil {
		return err
	}
	if wr != nil {
		return t.Execute(wr, data)
	}
	return t.Execute(&g.buff, data)
}

func (g Gtpl) GetBuffer() bytes.Buffer {
	return g.buff
}

func (g Gtpl) GetBytes() []byte {
	return g.buff.Bytes()
}

func (g Gtpl) String() string {
	return g.buff.String()
}

// func (g Gtpl) parseFiles(filenames ...string) (*template.Template, error) {
// 	t := &g.t
// 	if len(filenames) == 0 {
// 		// Not really a problem, but be consistent.
// 		return nil, fmt.Errorf("template: no files named in call to ParseFiles")
// 	}

// 	for _, filename := range filenames {
// 		name, b, err := readFileOS(filename)
// 		if err != nil {
// 			return nil, err
// 		}
// 		s := string(b)
// 		// First template becomes return value if not already defined,
// 		// and we use that one for subsequent New calls to associate
// 		// all the templates together. Also, if this file has the same name
// 		// as t, this file becomes the contents of t, so
// 		//  t, err := New(name).Funcs(xxx).ParseFiles(name)
// 		// works. Otherwise we create a new template associated with t.
// 		var tmpl *template.Template
// 		if t == nil {
// 			t = createTpl(name).Delims(g.delimLeft, g.delimRight) // template.New(name)
// 		}
// 		if name == t.Name() {
// 			tmpl = t
// 		} else {
// 			tmpl = t.New(name)
// 		}
// 		_, err = tmpl.Parse(s)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return t, nil
// }

// func readFileOS(file string) (name string, b []byte, err error) {
// 	name = filepath.Base(file)
// 	b, err = os.ReadFile(file)
// 	return
// }

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

func createTpl(name string) *template.Template {
	return template.New(name).Funcs(getDefaultFuncMap())
}
