package controllers

import (
	"fmt"
)

type User struct {
	Name string `json:"name" valid:"Required;MinSize(6)"`
	Age  int    `json:"age"`
	Sex  string `json:"sex" valid:"Required;MaxSize(6)"`
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
	this.response.Success(&user)
	this.response.Msg = "I'm struct"
}
