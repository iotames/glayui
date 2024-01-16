package component

import (
	"fmt"
	"strings"
)

type Button struct {
	Text string
}

const COMPONENT_SEND_MSG = "SEND_MOBILE_SMS"

type IFromItem interface {
	String() string
	Label() string
	Name() string
	ComponentType() string
	Js() string
}

type BaseFormItem struct {
	icon, label, name, value, placeholder string
	componentType, layVerify, jsText      string
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
func (i InputItem) ComponentType() string {
	return i.componentType
}
func (i InputItem) Js() string {
	return i.jsText
}
func (i InputItem) String() string {
	if i.icon == "" {
		i.icon = "layui-icon-edit"
	}
	if i.placeholder == "" {
		i.placeholder = i.label
	}
	reqtext := ""
	if strings.Contains(i.layVerify, "required") {
		reqtext = fmt.Sprintf(`lay-reqtext="请填写%s"`, i.label)
	}
	result := ""
	if i.componentType == COMPONENT_SEND_MSG {
		result = fmt.Sprintf(`
		<div class="layui-form-item">
		<div class="layui-row">
		  <div class="layui-col-xs7">
			<div class="layui-input-wrap">
			  <div class="layui-input-prefix">
				<i class="layui-icon layui-icon-cellphone"></i>
			  </div>
			  <input type="text" name="%s" value="" lay-verify="%s" placeholder="%s" %s autocomplete="off" class="layui-input" id="reg-cellphone">
			</div>
		  </div>
		  <div class="layui-col-xs5">
			<div style="margin-left: 11px;">
			  <button type="button" class="layui-btn layui-btn-fluid layui-btn-primary" lay-on="reg-get-vercode">获取验证码</button>
			</div>
		  </div>
		</div>
		</div>
		`, i.name, i.layVerify, i.placeholder, reqtext)
	} else {
		result = fmt.Sprintf(`
		<div class="layui-form-item">
		<div class="layui-input-wrap">
		  <div class="layui-input-prefix">
			<i class="layui-icon %s"></i>
		  </div>
		  <%s type="%s" name="%s" value="%s" lay-verify="%s" placeholder="%s" %s autocomplete="off" class="layui-input">
		</div>
		</div>
		`, i.icon, i.componentType, i.inputType, i.name, i.value, i.layVerify, i.placeholder, reqtext)
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

// NewMobileInputItemForSendMsg("cellphone", "手机号")
func NewMobileInputItemForSendMsg(name, label, posturl string) *InputItem {
	i := &InputItem{inputType: "text"}
	i.name = name
	i.label = label
	i.componentType = COMPONENT_SEND_MSG
	i.layVerify = "required|phone"
	i.jsText = fmt.Sprintf(`
	var xhr = new XMLHttpRequest();
	xhr.open('post', '%s', true);
	xhr.setRequestHeader('Content-Type', 'application/json');
	var mobile = $("#reg-cellphone").val()
	data = {mobile: mobile}
	xhr.send(JSON.stringify(data));
	xhr.onreadystatechange = function() {
	  if (xhr.readyState === 4 && xhr.status === 200) {
		console.log(xhr);
		layer.msg(JSON.parse(xhr.responseText).msg);
	  }
	};
`, posturl)
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
