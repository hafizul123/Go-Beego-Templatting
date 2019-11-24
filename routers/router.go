package routers

import (
	"first_go_web_app/controllers"
	"github.com/astaxie/beego"

)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/contact", &controllers.ContactController{})
	beego.Router("/me/?:id", &controllers.ContactController{},"get:MyName")

}
