package models

import (
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
	"time"
)


type Category struct {
	Id              int64
	Title           string
	Created         time.Time `arm:"index"`
	Views           int64     `arm:"index"`
	TopicTime       time.Time `arm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Labels          string
	Content         string `arm:"size(5000)"`
	Attachment      string
	Created         time.Time `arm:"index"`
	Updated         time.Time `arm:"index"`
	Views           int64
	Auther          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `arm:"index"`
}

func GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}



func AddReply(tid, nickname, content string) error {
	tidnum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	reply := &Comment{
		Tid:     tidnum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)

	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("Id", tidnum).One(topic)

	if err == nil {
		topic.ReplyCount++
		topic.ReplyTime = reply.Created
		_, err = o.Update(topic)
		if err != nil {
			beego.Debug(err)
		}
	}
	return err

}

func GetAllReplies(tid string) ([]*Comment, error) {

	o := orm.NewOrm()
	comment := make([]*Comment, 0)

	qs := o.QueryTable("comment")
	_, err := qs.Filter("Tid", tid).All(&comment)

	return comment, err

}

func DeleteReply(rid string) error {
	ridnum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	reply := &Comment{Id: ridnum}
	_, err = o.Delete(reply)

	return err
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name, Created: time.Now(), Views: 0, TopicTime: time.Now(), TopicCount: 0}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)

	if err == nil {
		return nil
	}
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)

	qs := o.QueryTable("category")

	_, err := qs.All(&cates)

	return cates, err
}

func DelCategory(id string) error {
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: uid}
	_, err = o.Delete(cate)
	return err
}

func AddTopic(title, category, label, content string) error {

	label = "$" + strings.Join(
		strings.Split(label, " "), "#$") + "#"

	o := orm.NewOrm()
	topic := &Topic{
		Title:     title,
		Content:   content,
		Labels:    label,
		Category:  category,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}

	_, err := o.Insert(topic)
	if err != nil {
		beego.Debug(err)
		return err
	}

	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)

	if err == nil {

		cate.TopicCount++
		if cate.Created.IsZero() {
			cate.Created = time.Now()
		}
		_, err = o.Update(cate)
		if err != nil {
			beego.Debug(err)
		}

	} else {
		cate.Title = category
		cate.TopicCount = 1
		cate.TopicTime = time.Now()
		cate.Views = 0
		cate.Created = time.Now()

		_, err = o.Insert(cate)
		if err != nil {
			beego.Debug(err)
			return err
		}
	}
	beego.Debug("add four", err)
	return err
}

func GetAllTopic(cate string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		if len(cate) > 0 {
			beego.Debug("GetAllTopic=", len(cate))
			qs = qs.Filter("Category", cate)
		}
		_, err = qs.OrderBy("-created").All(&topics)

	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func DelTopic(id string) error {
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	var oldCate string
	topic := &Topic{Id: uid}
	_, err = o.Delete(topic)
	if o.Read(topic) == nil {
		oldCate = topic.Category

	}
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	return err
}

func GetTopic(id string) (*Topic, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("Id", tid).One(topic)
	if err != nil {
		return nil, err
	}
	//_, err = o
	topic.Views++
	_, err = o.Update(topic)

	topic.Labels = strings.Replace(
		strings.Replace(topic.Labels, "#", " ", -1), "$", "", -1)

	return topic, err
}

func ModifyTopic(tid, title, category, label, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	var oldCate string
	topic := &Topic{Id: tidNum}

	label = "$" + strings.Join(
		strings.Split(label, " "), "#$") + "#"

	if o.Read(topic) == nil {
		oldCate = topic.Category
		topic.Title = title
		topic.Labels = label
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now()
		_, err = o.Update(topic)
		if err != nil {
			beego.Debug("modify err", err)
			return err
		}

	}
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
			if err != nil {
				beego.Debug("modify update", err)
			}
		}
	}

	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", oldCate).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
		if err != nil {
			beego.Debug("modify update", err)
		}
	}
	return nil
}
