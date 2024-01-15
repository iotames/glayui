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
	"github.com/iotames/glayui/component"
	"github.com/iotames/glayui/gtpl"
	"github.com/iotames/glayui/web"
)

func main() {
	tpl := gtpl.GetTpl()
	tpl.SetResourceDirPath("glayui/resource")
	s := web.NewEasyServer(":1598")
	cpt := component.NewLayout("")
	s.AddHandler("GET", "/layout", func(ctx web.Context) {
		cpt.SetTitle("THIS is TITLE")
		cpt.SetContent("hello This is Content 99999999")
		cpt.Exec(ctx.Writer)
	})
	s.ListenAndServe()
}

```
