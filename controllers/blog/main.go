package blog

import (
	"github.com/astaxie/beego"
	"myBlog/models"
	"strconv"
	"strings"
)

type MainController struct {
	baseController
}

//首页, 只显示前N条
func (this *MainController) Index() {
	var (
		list     []*models.Post
		pagesize int
		err      error
		page     int
	)

	if page, err = strconv.Atoi(this.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}

	if pagesize, err = strconv.Atoi(this.getOption("pagesize")); err != nil || pagesize < 1 {
		pagesize = 10
	}
	beego.Debug("pagesize=", pagesize, "offset=", (page-1)*pagesize)
	query := new(models.Post).Query().Filter("status", 0).Filter("urltype", 0)
	count, _ := query.Count()

	beego.Debug("count=", count)
	if count > 0 {
		query.OrderBy("-is_top", "-post_time").Limit(pagesize, (page-1)*pagesize).All(&list)
	}

	this.Data["count"] = count
	this.Data["list"] = list
	beego.Debug("list=", list)
	this.Data["pagebar"] = 1 //util.NewPager(page, int(count), pagesize, "").ToString()
	this.setHeadMetas()
	this.display("home")
	beego.Debug("list=", list)
}

//文章显示
func (this *MainController) Show() {
	var (
		post models.Post
		err  error
	)

	urlname := this.Ctx.Input.Param(":urlname")
	//beego.Debug("urlname=", urlname)
	if urlname != "" {
		post.UrlName = urlname
		err = post.Read("urlname")
	} else {
		id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
		post.Id = id
		err = post.Read()
	}
	if err != nil || post.Status != 0 {
		beego.Debug("show err:", err)
		this.Abort("404")
		return
	}

	post.Views++
	post.Update("Views")

	post.Content = strings.Replace(post.Content, "_ueditor_page_break_tag_", "", -1)

	this.Data["post"] = post
	//beego.Debug("show,post", post)
	this.setHeadMetas(post.Title, strings.Trim(post.Tags, ","), post.Title)
	this.display("article")
}

//历史归档
func (this *MainController) Archives() {
	var (
		page     int
		pagesize int
		err      error
		count    int64
		result   map[string][]*models.Post
	)

	if page, err = strconv.Atoi(this.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}

	if pagesize, err = strconv.Atoi(this.getOption("pagesize")); err != nil || pagesize < 1 {
		pagesize = 20
	} else {
		pagesize *= 2
	}

	query := new(models.Post).Query().Filter("status", 0).Filter("urltype", 0)

	count, _ = query.Count()
	result = make(map[string][]*models.Post)
	if count > 0 {
		var list []*models.Post
		query.OrderBy("-post_time").Limit(pagesize, (page-1)*pagesize).All(&list)
		for _, v := range list {
			year := v.PostTime.Format("2006")
			if _, ok := result[year]; !ok {
				result[year] = make([]*models.Post, 0)
			}
			result[year] = append(result[year], v)
		}
	}

	this.Data["page"] = page
	this.Data["count"] = count
	this.Data["pagesize"] = pagesize
	this.Data["pagebar"] = 1 //util.NewPager(page, int(count), pagesize, "/archives").ToString()
	this.Data["result"] = result

	this.setHeadMetas("归档")
	this.display("archives")
}

func (this *MainController) View() {
	var (
		post models.Post
		err  error
	)

	Title := this.Ctx.Input.Param(":Title")
	beego.Debug(";Title=", Title)
	if Title != "" {
		post.Title = Title
		err = post.Read("Title")
	} else {
		id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
		post.Id = id
		err = post.Read()
	}
	if err != nil || post.Status != 0 {
		beego.Debug("show err:", err)
		this.Abort("404")
		return
	}

	post.Views++
	post.Update("Views")

	post.Content = strings.Replace(post.Content, "_ueditor_page_break_tag_", "", -1)

	this.Data["post"] = post
	//beego.Debug("show,post", post)
	this.setHeadMetas(post.Title, strings.Trim(post.Tags, ","), post.Title)
	this.display("post")
}

//分类查看
func (this *MainController) Category() {
	var (
		page     int
		pagesize int
		name     string
		err      error
		count    int64
		result   map[string][]*models.Post
	)
	name = this.Ctx.Input.Param(":name")
	beego.Debug("name", name)
	if page, err = strconv.Atoi(this.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}
	if pagesize, err = strconv.Atoi(this.getOption("pagesize")); err != nil || pagesize < 1 {
		pagesize = 20
	} else {
		pagesize *= 2
	}

	tagpost := new(models.TagPost)
	tag := new(models.Tag)
	tag.Name = name

	if tag.Read("Name") != nil {
		this.Abort("404")
	}

	query := tagpost.Query().Filter("tag_id", tag.Id).Filter("post_status", 0)
	count, _ = query.Count()
	result = make(map[string][]*models.Post)
	if count > 0 {
		var tp []*models.TagPost
		var list []*models.Post
		var pids []int = make([]int, 0)

		query.OrderBy("-post_time").Limit(pagesize, (page-1)*pagesize).All(&tp)
		for _, v := range tp {
			pids = append(pids, v.PostId)
		}

		new(models.Post).Query().Filter("id__in", pids).All(&list)

		for _, v := range list {
			year := v.PostTime.Format("2006")
			if _, ok := result[year]; !ok {
				result[year] = make([]*models.Post, 0)
			}
			result[year] = append(result[year], v)
		}
	}

	this.Data["tag"] = tag
	this.Data["page"] = page
	this.Data["pagesize"] = pagesize
	this.Data["count"] = count
	this.Data["result"] = result
	this.Data["pagebar"] = 1 //util.NewPager(page, int(count), pagesize, tag.Link()).ToString()

	this.setHeadMetas(tag.Name, tag.Name, tag.Name)
	this.display("category")
}
