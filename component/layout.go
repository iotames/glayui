package component

import (
	"fmt"
	"net/http"

	"github.com/iotames/glayui/web/response"
)

type LayoutData struct {
	Title, Content string
}

type Layout struct {
	BaseComponent
	title, content string
}

// NewLayout 布局文件
// 使用 USE_EMBED_TPL 环境变量 设置是否使用嵌入静态资源文件。USE_EMBED_TPL=1 使用嵌入的资源文件。否则读取外部静态文件。
func NewLayout(fpath string) *Layout {
	l := &Layout{}
	if fpath == "" {
		l.tplpath = "tpl/layout.html"
	}
	l.name = "LAYOUT"
	return l
}

func (l *Layout) SetContent(content string) {
	l.content = content
}

func (l *Layout) SetTitle(title string) {
	l.title = title
}

func (l Layout) Handler(w http.ResponseWriter, r *http.Request) {
	if l.gtpl == nil {
		panic("gtpl can not be empty")
	}
	dt := LayoutData{Title: l.title, Content: l.content}
	// err := l.gtpl.SetDataFromResource(l.tplpath, l.bd, w)
	err := l.gtpl.SetDataByTplFile(l.tplpath, dt, w)
	if err != nil {
		resp := response.NewApiDataServerError(err.Error())
		w.Write(resp.Bytes())
		fmt.Printf("----服务器错误(%v)---\n", resp.String())
	}
}
