package controllers

import (
	"github.com/astaxie/beego"
	"myBlog/models"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsTopic"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelTopic(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/topic", 302)
	}

	var err error
	c.Data["Topics"], err = models.GetAllTopic("", false)

	if err != nil {
		beego.Error(err)
	}

	c.TplNames = "topic.html"
}

func (this *TopicController) Add() {
	this.TplNames = "topic_add.html"

}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	label := this.Input().Get("label")
	category := this.Input().Get("category")

	if len(tid) != 0 {
		var err error
		beego.Debug("modify", title, content, label, category)
		err = models.ModifyTopic(tid, title, category, label, content)
		if err != nil {
			beego.Debug(err)
			this.Redirect("/topic", 302)
			return
		}
		this.Redirect("/topic", 302)

	} else {
		op := this.Input().Get("op")
		switch op {
		case "add":
			title := this.Input().Get("title")
			content := this.Input().Get("content")
			category := this.Input().Get("category")
			var err error
			err = models.AddTopic(title, category, label, content)
			if err != nil {
				return
			}
			this.Redirect("/topic", 302)
		case "del":
			id := this.Input().Get("id")
			if len(id) == 0 {
				break
			}

			err := models.DelTopic(id)
			if err != nil {
				beego.Error(err)
			}
			this.Redirect("/topic", 302)
		}
	}

}

func (this *TopicController) View() {
	this.TplNames = "topic_view.html"

	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))

	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Labels"] = strings.Split(topic.Labels, " ")
	this.Data["Tid"] = this.Ctx.Input.Param("0")

	replies, err := models.GetAllReplies(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	categories, err := models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
}

func (this *TopicController) Modify() {
	this.TplNames = "topic_modify.html"
	tid := this.Input().Get("tid")

	topic, err := models.GetTopic(tid)

	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}
