package component

import (
	"fmt"
	"io"

	"github.com/iotames/glayui/web/response"
)

type TableData struct {
	Name   string
	IsWin  bool
	Age    int
	Weight float64
}

type Table struct {
	BaseComponent
	data TableData
}

// NewTable 数据表格
// fpath 相对于resource目录的相对路径。
// 使用 USE_EMBED_TPL 环境变量 设置是否使用嵌入静态资源文件。USE_EMBED_TPL=1 使用嵌入的资源文件。否则读取外部静态文件。
func NewTable(fpath string) *Table {
	l := &Table{}
	l.tplpath = fpath
	if l.tplpath == "" {
		l.tplpath = "tpl/table.html"
	}
	l.name = "TABLE"
	l.SetGtpl(defaultGtpl)
	l.UseEmbedTpl(gUseEmbedTpl)
	return l
}

func (l *Table) SetData(dt TableData) {
	l.data = dt
}

func (l Table) Exec(w io.Writer) {
	var err error
	if l.gtpl == nil {
		panic("gtpl can not be empty")
	}
	if l.useEmbedTpl {
		err = l.gtpl.SetDataFromResource(l.tplpath, l.data, w)
	} else {
		err = l.gtpl.SetDataByTplFile(l.tplpath, l.data, w)
	}
	if err != nil {
		resp := response.NewApiDataServerError(err.Error())
		w.Write(resp.Bytes())
		fmt.Printf("----服务器错误(%v)---\n", resp.String())
	}
}
