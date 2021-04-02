package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/hello", &controllers.HelloControllers{})
	beego.Include(&controllers.HelloControllers{})
	beego.Router("/Api/?:id", &controllers.RegularController{})
}
