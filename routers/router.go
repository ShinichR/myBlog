package routers

import (
	"github.com/astaxie/beego"
	"myBlog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
