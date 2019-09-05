package controllers

import (
	"go-blog-api/blog-api/models"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/gofrs/uuid"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

/*
NewPost 管理员添加文章
*/
type NewPost struct {
	beego.Controller
}

/*
Post /api/post/add_post
*/
func (c *NewPost) Post() {
	token := c.Ctx.Request.Header.Get("token")
	uid, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		o := orm.NewOrm()
		title := c.GetString("title")
		po := models.Post{Title: title}
		error := o.Read(&po, "Title")
		if error == orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 301, Result: "已经有这个文章", Data: &po}
		}
		tid, err := c.GetInt("tid")
		desc := c.GetString("desc")
		categoryID, err := c.GetInt("categoryId")
		image := c.GetString("image")
		content := c.GetString("content")
		markdown := c.GetString("markdown")
		post := models.Post{}
		pid, err := uuid.NewV4()
		if err != nil {
			logs.Info(err)
		}
		post.Uid = pid.String()
		post.Title = title
		post.Tid = tid
		post.UserId = uid
		post.Desc = desc
		post.Image = image
		post.CategoryId = categoryID
		post.Markdown = markdown
		post.Content = content
		post.Likes = 0
		post.Pv = 0
		post.Time = time.Now()
		logs.Info(o.Insert(post))
		c.Data["json"] = models.Message{Code: 200, Result: "添加文章成功", Data: post}
	}
}

/*
UpdatePost 修改文章
*/
type UpdatePost struct {
	beego.Controller
}

/*
Put /api/post/update_post
*/
func (c *UpdatePost) Put() {
	token := c.Ctx.Request.Header.Get("token")
	uid, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		pid := c.GetString("pid")
		post := models.Post{Uid: pid}
		o := orm.NewOrm()
		err := o.Read(&post, "Uid")
		if err == orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 200, Result: "没有这文章", Data: &post}
			c.ServeJSON()
		}
		title := c.GetString("title")
		tid, err := c.GetInt("tid")
		desc := c.GetString("desc")
		categoryID, err := c.GetInt("categoryId")
		image := c.GetString("image")
		content := c.GetString("content")
		markdown := c.GetString("markdown")
		if err != nil {
			logs.Info(err)
		}
		post.Title = title
		post.Tid = tid
		post.UserId = uid
		post.Desc = desc
		post.Image = image
		post.CategoryId = categoryID
		post.Markdown = markdown
		post.Content = content
		logs.Info(o.Update(&post))
		c.Data["json"] = models.Message{Code: 200, Result: "修改成功", Data: &post}
		c.ServeJSON()
	}
}

/*
DeletePost 删除文章
*/
type DeletePost struct {
	beego.Controller
}

/*
Delete /api/post/delete_post
*/
func (c *DeletePost) Delete() {
	token := c.Ctx.Request.Header.Get("token")
	uid, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		pid := c.GetString("pid")
		o := orm.NewOrm()
		post := models.Post{Uid: pid}
		error := o.Read(&post, "Uid")
		if error == orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 200, Result: "没有这文章", Data: &post}
			c.ServeJSON()
		}
		logs.Info(o.Delete(&post, "Uid"))
		logs.Info(uid)
		c.Data["json"] = models.Message{Code: 200, Result: "删除成功", Data: &post}
		c.ServeJSON()
	}
}

/*
PostList 博客列表
*/
type PostList struct {
	beego.Controller
}

/*
Get /api/post_list
*/
func (c *PostList) Get() {
	o := orm.NewOrm()
	var posts []*models.Post
	size, err := c.GetInt("size")
	str := c.Ctx.Input.Param(":id")
	page, err := strconv.Atoi(str)
	offset := (page - 1) * size
	num, err := o.QueryTable("post").OrderBy("-Id").Limit(size, offset).All(&posts)
	if err != nil {
		logs.Info(err)
	}
	if int(num) < size {
		next := "无"
		previous := page - 1
		po := map[string]interface{}{"num": num, "previous": previous, "next": next, "data": posts}
		c.Data["json"] = models.Message{Code: 200, Result: "成功", Data: po}
		c.ServeJSON()
	} else if int(num) > size {
		c.Data["json"] = models.Message{Code: 200, Result: "error", Data: nil}
		c.ServeJSON()
	} else {
		next := page + 1
		previous := page - 1
		po := map[string]interface{}{"num": num, "previous": previous, "next": next, "data": posts}
		c.Data["json"] = models.Message{Code: 200, Result: "成功", Data: po}
		c.ServeJSON()
	}
}

/*
PostController 查看单篇文章
*/
type PostController struct {
	beego.Controller
}

/*
Get /api/post_id/:id([0-9]+)
*/
func (c *PostController) Get() {
	str := c.Ctx.Input.Param(":id")
	is := c.GetString("is")
	id, err := strconv.Atoi(str)
	if err != nil {
		logs.Info(err)
	}
	o := orm.NewOrm()
	post := models.Post{Id: id}
	logs.Info(o.Read(&post, "Id"))
	if is == "admin" {
		c.Data["json"] = models.Message{Code: 200, Result: "管理员看文章", Data: post}
		c.ServeJSON()
	} else {
		post.Pv++
		logs.Info(o.Update(&post))
		c.Data["json"] = models.Message{Code: 200, Result: "用户看文章", Data: post}
		c.ServeJSON()
	}
}
