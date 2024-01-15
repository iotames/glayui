package component

import (
	"fmt"
	"io"

	"github.com/iotames/glayui/web/response"
)

type FormData struct {
	Title, Url   string
	SubmitButton Button
	SubmitMethod string
	SubmitUrl    string
	FormItems    []IFromItem
}

type Form struct {
	BaseComponent
	data FormData
}

// NewForm 表单
// fpath 相对于resource目录的相对路径。
// 使用 USE_EMBED_TPL 环境变量 设置是否使用嵌入静态资源文件。USE_EMBED_TPL=1 使用嵌入的资源文件。否则读取外部静态文件。
func NewForm(fpath string) *Form {
	l := &Form{}
	l.tplpath = fpath
	if l.tplpath == "" {
		l.tplpath = "tpl/form.html"
	}
	l.name = "FORM"
	l.SetGtpl(defaultGtpl)
	l.UseEmbedTpl(gUseEmbedTpl)
	return l
}

func (l *Form) SetData(dt FormData) {
	l.data = dt
}

func (l *Form) SetTitle(title string) *Form {
	l.data.Title = title
	return l
}

func (l *Form) SetSubmitButton(btn Button) *Form {
	l.data.SubmitButton = btn
	return l
}
func (l *Form) AddFormItem(item IFromItem) *Form {
	l.data.FormItems = append(l.data.FormItems, item)
	return l
}
func (l *Form) SetSubmitUrl(suburl string) *Form {
	l.data.SubmitUrl = suburl
	return l
}
func (l *Form) setDefaultData() {
	if l.data.SubmitButton.Text == "" {
		l.data.SubmitButton.Text = "提交"
	}
	if l.data.SubmitMethod == "" {
		l.data.SubmitMethod = "post"
	}
}
func (l Form) Exec(w io.Writer) {
	var err error
	if l.gtpl == nil {
		panic("gtpl can not be empty")
	}
	l.setDefaultData()
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
