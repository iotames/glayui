package component

import (
	"fmt"
)

type Button struct {
	Text string
}

type IFromItem interface {
	String() string
	Label() string
	Name() string
}

type BaseFormItem struct {
	label, name, componentType string
}

type InputItem struct {
	BaseFormItem
	inputType string
}

func (i InputItem) Name() string {
	return i.name
}
func (i InputItem) Label() string {
	return i.label
}
func (i InputItem) String() string {
	result := fmt.Sprintf(`
	<div class="layui-form-item">
    <label class="layui-form-label">%s</label>
    <div class="layui-input-block">
      <%s type="%s" name="%s" lay-verify="required" placeholder="请输入" autocomplete="off" class="layui-input">
    </div>
  </div>
`, i.Label(), i.componentType, i.inputType, i.Name())

	if i.componentType == "mobilesms" {
		result = fmt.Sprintf(mobileSmsFmt, i.Name())
	}

	return result
}

func NewTextInputItem(name, label string) *InputItem {
	i := &InputItem{inputType: "text"}
	i.name = name
	i.label = label
	i.componentType = "input"
	return i
}

func NewTelInputItem(name, label string) *InputItem {
	i := &InputItem{inputType: "tel"}
	i.name = name
	i.label = label
	i.componentType = "input"
	return i
}

// NewMobileMsgInputItem("cellphone", "手机号")
func NewMobileMsgInputItem(name, label string) *InputItem {
	i := &InputItem{inputType: "text"}
	i.name = name
	i.label = label
	i.componentType = "mobilesms"
	return i
}
func NewNumberInputItem(name, label string) *InputItem {
	i := &InputItem{inputType: "number"}
	i.name = name
	i.label = label
	i.componentType = "input"
	return i
}

func NewPasswordInputItem(name, label string) *InputItem {
	i := &InputItem{inputType: "password"}
	i.name = name
	i.label = label
	i.componentType = "input"
	return i
}

var mobileSmsFmt = `
<div class="layui-form-item">
<div class="layui-row">
  <div class="layui-col-xs7">
	<div class="layui-input-wrap">
	  <div class="layui-input-prefix">
		<i class="layui-icon layui-icon-cellphone"></i>
	  </div>
	  <input type="text" name="%s" value="" lay-verify="required|phone" placeholder="手机号" lay-reqtext="请填写手机号" autocomplete="off" class="layui-input" id="reg-cellphone">
	</div>
  </div>
  <div class="layui-col-xs5">
	<div style="margin-left: 11px;">
	  <button type="button" class="layui-btn layui-btn-fluid layui-btn-primary" lay-on="reg-get-vercode">获取验证码</button>
	</div>
  </div>
</div>
</div>
`
