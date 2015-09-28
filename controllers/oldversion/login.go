package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exit") == "True"
	beego.Debug("login get =", isExit)
	if isExit {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("passwd", "", -1, "/")
		uname := c.Ctx.GetCookie("uname")
		passwd := c.Ctx.GetCookie("passwd")
		beego.Debug("blog[debug] exit", uname, passwd)
		c.Redirect("/", 301)
		return
	}

	c.TplNames = "login.html"

}

func (c *LoginController) Post() {
	uname := c.Input().Get("uname")
	passwd := c.Input().Get("passwd")
	autoLogin := c.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("passwd") == passwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("uname", uname, maxAge, "/")
		c.Ctx.SetCookie("passwd", passwd, maxAge, "/")
		c.Data["LoginErr"] = false

	}

	c.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	uname := ctx.GetCookie("uname")
	/*if err != nil {
		return false
	}*/
	//uname := ok.Value
	passwd := ctx.GetCookie("passwd")
	/*if err != nil {
		return false
	}*/
	//passwd := ok.Value
	beego.Debug("blog[debug]", uname, passwd)
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("passwd") == passwd {
		return true
	}
	return false
}
