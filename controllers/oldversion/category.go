package controllers

import (
	"github.com/astaxie/beego"
	"myBlog/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {

	c.Data["IsLogin"] = checkAccount(c.Ctx)

	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 302)
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 302)
	}

	c.Data["IsCategory"] = true
	c.TplNames = "category.html"

	var err error
	c.Data["Categories"], err = models.GetAllCategory()

	if err != nil {
		beego.Error(err)
	}
}
