package controllers

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

type HelloControllers struct {
	BaseController
}

// @router /hello [get]
func (hello *HelloControllers) Get() {
	hello.Ctx.WriteString("GET hello world")
}

func (hello *HelloControllers) Post() {
	hello.Ctx.WriteString("POST hello world")
}

// @router /hello/JsonFunc [get]
func (this *HelloControllers) JsonFunc() {
	user := &User{"uuu", 20, "m"}
	this.Data["json"] = user
	this.ServeJSON()
}

// @router /hello/PostStruct [post]
func (this *HelloControllers) PostStruct() {
	fmt.Println("PostStruct")
	var user User
	this.GetJSON(&user)
	valid := validation.Validation{} //创建验证数据对象
	valid.Required(user.Name, "Name").Message("真实姓名不能为空")
	this.response.Success(&user)
	this.response.Msg = "I'm struct"
}
