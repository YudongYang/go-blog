package controllers

import "github.com/astaxie/beego"

type RegularController struct {
	beego.Controller
}

func (reg *RegularController) Get() {
	id := reg.Ctx.Input.Param(":id")
	reg.Ctx.WriteString("id=" + id + "\n")
}
