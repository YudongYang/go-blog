package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	params   interface{}
	response Response
}

func (this *BaseController) GetJSON(obj interface{}) {
	data := this.Ctx.Input.RequestBody
	this.params = obj
	err := json.Unmarshal(data, &this.params)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
		this.Fail("json.Unmarshal is err:" + err.Error())
		this.Finish()
		this.StopRun()
	}
}

func (this *BaseController) Success(data interface{}) {
	this.response.Success(&data)
}

func (this *BaseController) Fail(msg string) {
	this.response.Fail(msg)
}

func (this *BaseController) Prepare() {
	fmt.Println("Prepare")
	this.EnableRender = false
}

func (this *BaseController) Finish() {
	fmt.Println("Finish")
	this.Data["json"] = this.response
	this.ServeJSON()
}
