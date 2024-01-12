package component

import (
	"fmt"
	"io"

	"github.com/iotames/glayui/web/response"
)

type LayoutData struct {
	Title, Content string
}

type Layout struct {
	BaseComponent
	data LayoutData
}

// NewLayout 布局文件
// fpath 相对于resource目录的相对路径。
// 使用 USE_EMBED_TPL 环境变量 设置是否使用嵌入静态资源文件。USE_EMBED_TPL=1 使用嵌入的资源文件。否则读取外部静态文件。
func NewLayout(fpath string) *Layout {
	l := &Layout{}
	l.tplpath = fpath
	if l.tplpath == "" {
		l.tplpath = "tpl/layout.html"
	}
	l.name = "LAYOUT"
	l.SetGtpl(defaultGtpl)
	l.UseEmbedTpl(gUseEmbedTpl)
	return l
}

func (l *Layout) SetContent(content string) {
	l.data.Content = content
}

func (l *Layout) SetTitle(title string) {
	l.data.Title = title
}

func (l *Layout) SetData(dt LayoutData) {
	l.data = dt
}

func (l Layout) Exec(w io.Writer) {
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
