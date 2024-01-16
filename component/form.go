package component

import (
	"fmt"
	"io"

	"github.com/iotames/glayui/web/response"
)

type FormData struct {
	Title        string
	SubmitButton Button
	SubmitMethod string
	SubmitUrl    string
	JsText       string
	LayFilter    string
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
	if l.data.LayFilter == "" {
		l.data.LayFilter = "demo-reg"
	}
	if l.data.JsText == "" {
		sendSmsJs := ""
		for _, fitem := range l.data.FormItems {
			if fitem.ComponentType() == COMPONENT_SEND_MSG {
				sendSmsJs = fitem.Js()
			}
		}
		l.data.JsText = fmt.Sprintf(`
		layui.use(function(){
			var $ = layui.$;
			var form = layui.form;
			var layer = layui.layer;
			var util = layui.util;

			// 提交事件
			form.on('submit(%s)', function(data){
			  var field = data.field; // 获取表单字段值
			  var xhr = new XMLHttpRequest();
			  xhr.open('%s', '%s', true);
			  xhr.setRequestHeader('Content-Type', 'application/json');
			  xhr.send(JSON.stringify(field));
			  xhr.onreadystatechange = function() {
				if (xhr.readyState === 4 && xhr.status === 200) {
				  console.log(xhr.responseText);
				  console.log(xhr);
				  layer.alert(xhr.responseText, {
					  title: '提交结果'
					});
				}
			  };

			  return false; // 阻止默认 form 跳转
			});
			
			// 普通事件
			util.on('lay-on', {
					  // 获取验证码
					  'reg-get-vercode': function(othis){
						var isvalid = form.validate('#reg-cellphone'); // 主动触发验证，v2.7.0 新增 
						// 验证通过
						if(isvalid){
						  %s
						  // 此处可继续书写「发送验证码」等后续逻辑
						  // …
						}
					  }
					});
		  });
`, l.data.LayFilter, l.data.SubmitMethod, l.data.SubmitUrl, sendSmsJs)
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
