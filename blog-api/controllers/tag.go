package controllers

import (
	"go-blog-api/blog-api/models"
	"go-blog-api/blog-api/util"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

/*
NewTag 新添标签
*/
type NewTag struct {
	beego.Controller
}

/*
Post /api/tag/add_tag
*/
func (c *NewTag) Post() {
	token := c.Ctx.Request.Header.Get("token")
	uid, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		name := c.GetString("name")
		ta := models.Tag{Name: name}
		o := orm.NewOrm()
		error := o.Read(&ta, "Name")
		if error != orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 200, Result: "你已经有这个分类了", Data: nil}
			c.ServeJSON()
		}
		tag := new(models.Tag)
		tag.Name = name
		tag.UserId = uid
		tag.Time = time.Now()
		logs.Info(o.Insert(tag))
		var user models.User
		_, err := o.QueryTable("user").Filter("Uid", uid).All(&user)
		if err != nil {
			logs.Info(err)
		}
		userView := util.UserToViews(user)
		tagView := util.TagToView(*tag, *userView)
		c.Data["json"] = models.Message{Code: 200, Result: "添加成功", Data: tagView}
		c.ServeJSON()
	}
}

/*
UpdateTag 修改标签
*/
type UpdateTag struct {
	beego.Controller
}

/*
Put /api/tag/update_tag
*/
func (c *UpdateTag) Put() {
	token := c.Ctx.Request.Header.Get("token")
	_, isAdmin := models.IsAdmin(token)
	name := c.GetString("name")
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		o := orm.NewOrm()
		ta := models.Tag{Name: name}
		er := o.Read(&ta, "Name")
		if er != orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 200, Result: "已经有这个分类的名字", Data: ta}
			c.ServeJSON()
		}
		tid, err := c.GetInt("tid")
		if err != nil {
			logs.Info(err)
		}
		tag := models.Tag{Id: tid}
		error := o.Read(&tag, "Id")
		if error == orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 200, Result: "没有这个分类", Data: nil}
			c.ServeJSON()
		}
		tag.Name = name
		logs.Info(o.Update(&tag, "Name"))

		var user models.User
		_, err = o.QueryTable("user").Filter("Uid", tag.UserId).All(&user)
		if err != nil {
			logs.Info(err)
		}
		userView := util.UserToViews(user)
		tagView := util.TagToView(tag, *userView)

		c.Data["json"] = models.Message{Code: 200, Result: "修改成功", Data: tagView}
		c.ServeJSON()
	}
}

/*
DeleteTag 删除分类
*/
type DeleteTag struct {
	beego.Controller
}

/*
Delete /api/tag/delete_tag
*/
func (c *DeleteTag) Delete() {
	token := c.Ctx.Request.Header.Get("token")
	_, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		tid, err := c.GetInt("tid")
		if err != nil {
			logs.Info(err)
		}
		o := orm.NewOrm()
		tag := models.Tag{Id: tid}
		error := o.Read(&tag, "Id")
		if error == orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 200, Result: "没有这个分类", Data: nil}
			c.ServeJSON()
		}
		logs.Info(o.Delete(&tag, "Id"))
		c.Data["json"] = models.Message{Code: 200, Result: "删除成功", Data: nil}
		c.ServeJSON()
	}
}

/*
GetTagList 查询所有分类
*/
type GetTagList struct {
	beego.Controller
}

/*
Get /api/tag_list
*/
func (c *GetTagList) Get() {
	o := orm.NewOrm()
	var tags []*models.Tag
	var tagViews []models.TagView
	var user models.User
	_, error := o.QueryTable("tag").OrderBy("Id").All(&tags)
	if error == orm.ErrArgs {
		c.Data["json"] = models.Message{Code: 200, Result: "没有分类", Data: nil}
		c.ServeJSON()
	}
	for _, v := range tags {
		_, err := o.QueryTable("user").Filter("Uid", v.UserId).All(&user)
		if err != nil {
			logs.Info(err)
		}
		userView := util.UserToViews(user)
		tagView := util.TagToView(*v, *userView)
		tagViews = append(tagViews, *tagView)
	}
	c.Data["json"] = models.Message{Code: 200, Result: "ok", Data: tagViews}
	c.ServeJSON()
}

/*
GetTagPosts 查询该分类下的文章
*/
type GetTagPosts struct {
	beego.Controller
}

/*
Get /api/tag_posts/:id([0-9]+)?page=1&size=5
*/
func (c *GetTagPosts) Get() {
	str := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(str)
	size, err := c.GetInt("size")
	page, err := c.GetInt("page")
	offset := (page - 1) * size
	if err != nil {
		logs.Info(err)
	}
	o := orm.NewOrm()
	tag := models.Tag{Id: id}
	error := o.Read(&tag, "Id")
	if error == orm.ErrNoRows {
		c.Data["json"] = models.Message{Code: 200, Result: "没有分类", Data: nil}
		c.ServeJSON()
	}
	var posts []*models.Post
	num, err := o.QueryTable("post").OrderBy("-Id").Filter("tid__icontains", &tag.Id).Limit(size, offset).All(&posts)
	logs.Info(&num)
	if err != nil {
		logs.Info(err)
	}
	if int(num) < size {
		next := "无"
		previous := page - 1
		data := map[string]interface{}{"previous": previous, "next": next, "data": posts}
		c.Data["json"] = models.Message{Code: 200, Result: "查询成功", Data: data}
		c.ServeJSON()
	} else if int(num) > size {
		c.Data["json"] = models.Message{Code: 301, Result: "出错了", Data: nil}
		c.ServeJSON()
	} else if int(num) == 0 {
		c.Data["json"] = models.Message{Code: 200, Result: "没有文章", Data: nil}
		c.ServeJSON()
	} else {
		next := page + 1
		previous := page - 1
		data := map[string]interface{}{"previous": previous, "next": next, "data": posts}
		c.Data["json"] = models.Message{Code: 200, Result: "查询成功", Data: data}
		c.ServeJSON()
	}
}
