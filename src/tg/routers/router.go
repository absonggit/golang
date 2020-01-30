package routers

import (
	"tg/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/send", &controllers.SendController{})
}
