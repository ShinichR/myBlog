package controllers

import (
	"github.com/astaxie/beego"
	"myBlog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplNames = "home.html"
	var err error
	//beego.Debug("cate=", c.Input().Get("cate"))
	c.Data["Topics"], err = models.GetAllTopic(c.Input().Get("cate"), true)
	//beego.Debug("date=", c.Data["Topics"].Created)
	if err != nil {
		beego.Error(err)
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)

	categories, err := models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories
}
