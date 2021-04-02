package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	beego.Controller
	params   interface{}
	response Response
	valid    validation.Validation
}

var MessageTmpls = map[string]string{
	"Required": "不能为空",
	"MinSize":  "最短长度为 %d",
	"Length":   "长度必须为 %d",
	"Numeric":  "必须是有效的数字",
	"Email":    "必须是有效的电子邮件地址",
	"Mobile":   "必须是有效的手机号码",
}

func init() {
	fmt.Println("init")
	validation.SetDefaultMessage(MessageTmpls)
}

func (this *BaseController) GetJSON(obj interface{}) {
	data := this.Ctx.Input.RequestBody
	this.params = obj
	err := json.Unmarshal(data, this.params)
	if err == nil {
		err = this.Validate()
	}
	if err != nil {
		fmt.Println("err:", err.Error())
		this.Fail("err:" + err.Error())
		this.Finish()
		this.StopRun()
	}
}

func (this *BaseController) Validate() error {
	if valid, err := this.valid.Valid(this.params); !valid {
		fmt.Println(valid)
		fmt.Println(err)
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		var err_list []string
		for _, err := range this.valid.Errors {
			err_list = append(err_list, "[ "+err.Key+" , "+err.Message+" ]")
		}
		fmt.Println(strings.Join(err_list, " , "))
		return errors.New(strings.Join(err_list, " , "))
	} else {
		return nil
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
