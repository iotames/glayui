package gtpl

import (
	"bytes"
	"fmt"
	"io/fs"

	"io"
	// "os"
	"path/filepath"
	"text/template"

	"github.com/iotames/glayui/resource"
)

const GTPL_NAME = "GTPL_NAME"

const GTPL_DELIM_LEFT = `<{%`
const GTPL_DELIM_RIGHT = `%}>`

type Gtpl struct {
	t                                      *template.Template
	funcMap                                template.FuncMap
	resourceDirPath, delimLeft, delimRight string
	buff                                   bytes.Buffer
}

func NewTpl(leftDelim, rightDelim string) *Gtpl {
	if leftDelim == "" || rightDelim == "" {
		panic("leftDelim or rightDelim could not be empty")
	}
	return &Gtpl{delimLeft: leftDelim, delimRight: rightDelim}
}
func (g *Gtpl) SetResourceDirPath(dpath string) *Gtpl {
	g.resourceDirPath = dpath
	return g
}
func (g *Gtpl) SetTemplate(t *template.Template) *Gtpl {
	g.t = t
	return g
}

func (g *Gtpl) SetDataByTplText(tpltext string, data any, wr io.Writer) error {
	g.t = g.newTpl(GTPL_NAME)
	t, err := g.t.Parse(tpltext)
	if err != nil {
		return err
	}
	return g.execTpl(t, data, wr)
}

// SetDataByTplFile 使用模板文件设置返回值
// fpath 相对于resource目录的相对路径。
func (g *Gtpl) SetDataByTplFile(fpath string, data any, wr io.Writer) error {
	g.t = g.newTpl(filepath.Base(fpath))
	t, err := g.t.ParseFiles(g.getResourceFullPath(fpath))
	if err != nil {
		return err
	}
	return g.execTpl(t, data, wr)
}

func (g Gtpl) getResourceFullPath(fpath string) string {
	resourcePath := g.resourceDirPath
	if resourcePath == "" {
		resourcePath = resource.RESOURCE_DIR
	}
	return filepath.Join(resourcePath, fpath)
}

// SetDataFromResource 使用内嵌的静态资源系统设置返回值
// fpath 相对于resource目录的相对路径。
func (g *Gtpl) SetDataFromResource(fpath string, data any, wr io.Writer) error {
	// fss, err := resource.GetResourceFile(g.getResourceFullPath(fpath))
	// if err != nil {
	// 	fmt.Printf("-----resource.GetResourceFile--err(%v)\n", err)
	// 	return err
	// }
	return g.SetDataByTplFS(fpath, resource.ResourceFs, data, wr)
}
func (g *Gtpl) SetDataByTplFS(fpath string, fsfs fs.FS, data any, wr io.Writer) error {
	g.t = g.newTpl(filepath.Base(fpath))
	t, err := g.t.ParseFS(fsfs, fpath)
	if err != nil {
		return err
	}
	return g.execTpl(t, data, wr)
}

func (g *Gtpl) newTpl(tplName string) *template.Template {
	return createTpl(tplName, g.funcMap).Delims(g.delimLeft, g.delimRight)
}

func (g *Gtpl) execTpl(t *template.Template, data any, wr io.Writer) error {
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
