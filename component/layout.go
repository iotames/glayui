package component

import (
	"fmt"
	"net/http"

	"github.com/iotames/glayui/gtpl"
	"github.com/iotames/glayui/web/response"
)

type LayoutData struct {
	Title, Content string
}

type Layout struct {
	gtpl            *gtpl.Gtpl
	resourceDirPath string
	tplpath         string
	title, content  string
}

func (l *Layout) SetResourceDirPath(dpath string) *Layout {
	l.resourceDirPath = dpath
	return l
}

func NewLayout(fpath string) *Layout {
	if fpath == "" {
		fpath = "tpl/layout.html"
	}
	return &Layout{tplpath: fpath}
}

func (l *Layout) SetContent(content string) {
	l.content = content
}
func (l *Layout) SetTitle(title string) {
	l.title = title
}

func (l Layout) Handler(w http.ResponseWriter, r *http.Request) {
	if l.gtpl == nil {
		l.gtpl = gtpl.GetTpl()
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
