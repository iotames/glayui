## 快速开始

```
# 新建项目文件夹并进入
mkdir myproject
cd myproject/

# 下载glayui资源包，路径为 ./glayui/resource
git clone https://github.com/iotames/glayui.git

# 启用 workspace(>= go1.18) 模式，添加glayui目录
go work init glayui

# 
```

编辑入口文件 `main.go`:

```
package main

import (
	"fmt"
	"io"

	"github.com/iotames/glayui/component"
	"github.com/iotames/glayui/gtpl"
	"github.com/iotames/glayui/web"
)

func main() {
	tpl := gtpl.GetTpl()
	tpl.SetResourceDirPath("glayui/resource")
	s := web.NewEasyServer(":1598")
	s.AddHandler("POST", "/login", func(ctx web.Context) {
		postdata, _ := io.ReadAll(ctx.Request.Body)
		fmt.Printf("---postdata(%s)---\n", string(postdata))
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Write([]byte(`{"msg":"登录成功","status":200}`))
	})
	s.AddHandler("POST", "/sendsms", func(ctx web.Context) {
		postdata, _ := io.ReadAll(ctx.Request.Body)
		fmt.Printf("---postdata(%s)---\n", string(postdata))
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Write([]byte(`{"msg":"发送成功","status":200}`))
	})
	s.AddHandler("GET", "/login", func(ctx web.Context) {
		fm := ctx.NewForm().SetTitle("用户登录").SetSubmitButton(component.Button{Text: "登录"})
		fm.SetSubmitUrl("/login")
		fm.AddFormItem(component.NewMobileInputItemForSendMsg("cellphone", "手机号", "/sendsms"))
		fm.AddFormItem(component.NewTextInputItem("smscode", "验证码"))
		fm.Exec(ctx.Writer)
	})
	// 访问 http://127.0.0.1:1598/login
	s.ListenAndServe()
}

```
