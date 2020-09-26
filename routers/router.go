package routers

import (
	"onework/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/", &controllers.MainController{})
}
