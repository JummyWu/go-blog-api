package controllers

import (
	"go-blog-api/blog-api/models"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gofrs/uuid"

	"github.com/astaxie/beego"
)

/*
LikesController 友联代理
*/
type LikesController struct {
	beego.Controller
}

/*
URLMapping 代理全部URL
*/
func (c *LikesController) URLMapping() {
	c.Mapping("NewLinks", c.Post)
	c.Mapping("UpdateLinkes", c.Put)
	c.Mapping("UpdateStatus", c.PutStatus)
	c.Mapping("DeleteLinks", c.Delete)
	c.Mapping("GetLinks", c.Get)
}

/*
Post /api/links/add_links
*/
func (c *LikesController) Post() {
	token := c.Ctx.Request.Header.Get("token")
	_, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		o := orm.NewOrm()
		email := c.GetString("email")
		li := models.Links{Email: email}
		err := o.Read(&li, "Email")
		if err != orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 301, Result: "已经有这个友联了", Data: li}
			c.ServeJSON()
		}
		name := c.GetString("hostname")
		image := c.GetString("hostimage")
		url := c.GetString("hosturl")
		body := c.GetString("body")
		link := new(models.Links)
		uid, err := uuid.NewV4()
		if err != nil {
			logs.Info(err)
		}
		link.Uid = uid.String()
		link.Is = 0
		link.Email = email
		link.HostName = name
		link.HostImage = image
		link.HostUrl = url
		link.Body = body
		link.Time = time.Now()
		logs.Info(o.Insert(link))
		c.Data["json"] = models.Message{Code: 200, Result: "ok", Data: link}
		c.ServeJSON()
	}
}

/*
Put /api/links/update_links
*/
func (c *LikesController) Put() {
	token := c.Ctx.Request.Header.Get("token")
	_, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		id := c.GetString("lid")
		name := c.GetString("hostname")
		image := c.GetString("hostimage")
		url := c.GetString("hosturl")
		body := c.GetString("body")
		email := c.GetString("email")
		is, err := c.GetInt("is")
		if err != nil {
			logs.Info(err)
		}
		o := orm.NewOrm()
		li := models.Links{Email: email}
		error := o.Read(&li, "Email")
		if error != orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 200, Result: "已经有这个友联了", Data: li}
			c.ServeJSON()
		}
		link := models.Links{Uid: id}
		logs.Info(o.Read(&link, "Uid"))
		if name != "" {
			link.HostName = name
		}
		link.Is = is
		if image != "" {
			link.HostImage = image
		}
		if url != "" {
			link.HostUrl = url
		}
		if body != "" {
			link.Body = body
		}
		logs.Info(o.Update(&link))
		c.Data["json"] = models.Message{Code: 200, Result: "修改成功", Data: link}
		c.ServeJSON()
	}
}

/*
PutStatus /api/links/update_status
*/
func (c *LikesController) PutStatus() {
	token := c.Ctx.Request.Header.Get("token")
	_, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		lid := c.GetString("lid")
		links := models.Links{Uid: lid}
		o := orm.NewOrm()
		error := o.Read(&links, "Uid")
		if error == orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 200, Result: "没有这个友联", Data: nil}
			c.ServeJSON()
		}
		is, err := c.GetInt("is")
		if err != nil {
			logs.Info(err)
		}
		links.Is = is
		logs.Info(o.Update(&links, "Uid"))
		c.Data["json"] = models.Message{Code: 200, Result: "已经设置成功", Data: links}
		c.ServeJSON()
	}
}

/*
Delete /api/links/delete_links
*/
func (c *LikesController) Delete() {
	token := c.Ctx.Request.Header.Get("token")
	_, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		lid := c.GetString("lid")
		links := models.Links{Uid: lid}
		o := orm.NewOrm()
		error := o.Read(&links, "Uid")
		if error == orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 301, Result: "没有这个友联", Data: nil}
			c.ServeJSON()
		}
		logs.Info(o.Delete(&links, "Uid"))
		c.Data["json"] = models.Message{Code: 200, Result: "删除成功", Data: nil}
		c.ServeJSON()
	}
}

/*
Get /api/links_list
*/
func (c *LikesController) Get() {
	var links []*models.Links
	o := orm.NewOrm()
	number, err := o.QueryTable("links").OrderBy("-Id").Filter("is", 1).All(&links)
	if err != nil {
		logs.Info(number)
	}
	c.Data["json"] = models.Message{Code: 200, Result: "成功", Data: links}
	c.ServeJSON()
}
