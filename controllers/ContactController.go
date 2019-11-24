package controllers

import(
	"github.com/astaxie/beego"

)
type ContactController struct{
	beego.Controller
}

func(this *ContactController) Get(){
	this.Data["headline"]="Contact Page"
	this.TplName="contact.tpl"
}
func (this *ContactController) MyName(){
	this.Data["headline"]=this.Ctx.Input.Param(":id")
	this.Data["id"]=this.Ctx.Input.Param(":id")

	this.TplName="contact.tpl"
}