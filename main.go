package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myBlog/controllers/admin"
	"myBlog/controllers/blog"
	"myBlog/models"
)

func init() {
	models.Init()

}

func main() {
	//orm.Debug = true
	orm.RunSyncdb("default", false, true)
	/*beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	*/
	//前台路由
	beego.Router("/", &blog.MainController{}, "*:Index")
	beego.Router("/page/:page:int", &blog.MainController{}, "*:Index")
	beego.Router("/article/:id:int", &blog.MainController{}, "*:Show")      //ID访问
	beego.Router("/article/:urlname(.+)", &blog.MainController{}, "*:Show") //别名访问
	beego.Router("/archives", &blog.MainController{}, "*:Archives")
	beego.Router("/archives/page/:page:int", &blog.MainController{}, "*:Archives")
	beego.Router("/category/:name(.+?)", &blog.MainController{}, "*:Category")
	beego.Router("/category/:name(.+?)/page/:page:int", &blog.MainController{}, "*:Category")
	beego.Router("/:urlname(.+)", &blog.MainController{}, "*:Show") //别名访问
	beego.Router("/view/:Title(.+?)", &blog.MainController{}, "*:View")

	//后台路由
	beego.Router("/admin", &admin.IndexController{}, "*:Index")
	beego.Router("/admin/main", &admin.IndexController{}, "*:Main")
	beego.Router("/admin/login", &admin.AccountController{}, "*:Login")
	//beego.Router("/admin/logout", &admin.AccountController{}, "*:Logout")
	//beego.Router("/admin/account/profile", &admin.AccountController{}, "*:Profile")

	//内容管理
	beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")
	beego.Router("/admin/article/save", &admin.ArticleController{}, "post:Save")
	beego.Router("/admin/article/delete", &admin.ArticleController{}, "*:Delete")
	beego.Router("/admin/article/batch", &admin.ArticleController{}, "*:Batch")
	beego.Router("/admin/article/upload", &admin.ArticleController{}, "*:Upload")

	//用户管理
	beego.Router("/admin/user/list", &admin.UserController{}, "*:List")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")
	beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")

	beego.Run()
}
