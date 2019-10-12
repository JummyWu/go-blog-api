package controllers

import (
	"go-blog-api/blog-api/models"
	"go-blog-api/blog-api/util"

	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gofrs/uuid"
)

//NewCategory 管理员添加分类
type NewCategory struct {
	beego.Controller
}

// Post /api/category/add_category
func (c *NewCategory) Post() {
	token := c.Ctx.Request.Header.Get("token")
	uid, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		o := orm.NewOrm()
		name := c.GetString("name")
		pid, err := uuid.NewV4()
		if err != nil {
			logs.Info(err)
		}
		var puuid = "category_" + pid.String()
		category := new(models.Category)
		category.Name = name
		category.Uid = puuid
		category.UserId = uid
		category.Time = time.Now()
		logs.Info(o.Insert(category))
		var user models.User
		_, err = o.QueryTable("user").Filter("Uid", uid).All(&user)
		if err != nil {
			logs.Info(nil)
		}
		userView := util.UserToViews(user)
		categoryView := util.CategoryToView(*category, *userView)

		c.Data["json"] = models.Message{Code: 200, Result: "添加分类成功", Data: categoryView}
		c.ServeJSON()
	}
}

//UpdateCategory 修改文章
type UpdateCategory struct {
	beego.Controller
}

//Put /api/category/update_category
func (c *UpdateCategory) Put() {
	token := c.Ctx.Request.Header.Get("token")
	_, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		categoryId := c.GetString("category_id")
		category := models.Category{Uid: categoryId}
		o := orm.NewOrm()
		err := o.Read(&category, "Uid")
		if err == orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 200, Result: "没有这分类", Data: &category}
			c.ServeJSON()
		}
		name := c.GetString("name")
		category.Name = name
		logs.Info(o.Update(&category))

		var user models.User
		_, err = o.QueryTable("user").Filter("uid", category.UserId).All(&user)
		if err != nil {
			logs.Info(err)
		}
		userView := util.UserToViews(user)
		categoryView := util.CategoryToView(category, *userView)

		c.Data["json"] = models.Message{Code: 200, Result: "修改成功", Data: &categoryView}
		c.ServeJSON()
	}
}

//DeleteCategory 删除文章
type DeleteCategory struct {
	beego.Controller
}

//Delete /api/category/delete_category
func (c *DeleteCategory) Delete() {
	token := c.Ctx.Request.Header.Get("token")
	_, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		categoryId := c.GetString("category_id")
		category := models.Category{Uid: categoryId}
		o := orm.NewOrm()
		err := o.Read(&category, "Uid")
		if err == orm.ErrNoRows {
			c.Data["json"] = models.Message{Code: 200, Result: "没有这分类", Data: &category}
			c.ServeJSON()
		}
		logs.Info(o.Delete(&category, "Uid"))
		c.Data["json"] = models.Message{Code: 200, Result: "删除成功", Data: &category}
		c.ServeJSON()
	}
}

//CategoryList 查询所有的分类
type CategoryList struct {
	beego.Controller
}

//Get /api/category/list
func (c *CategoryList) Get() {
	var categorys []models.Category
	var categoryViews []*models.CategoryView
	o := orm.NewOrm()
	_, err := o.QueryTable("category").OrderBy("-Id").All(&categorys)
	if err != nil {
		logs.Info(err)
	}
	for _, v := range categorys {
		var user models.User
		_, err = o.QueryTable("user").Filter("Uid", v.UserId).All(&user)
		if err != nil {
			logs.Info(err)
		}
		userView := util.UserToViews(user)
		categoryView := util.CategoryToView(v, *userView)

		categoryViews = append(categoryViews, categoryView)
	}
	c.Data["json"] = models.Message{Code: 200, Result: "ok", Data: categoryViews}
	c.ServeJSON()
}

//CategoryPost 获取分类下的文章
type CategoryPost struct {
	beego.Controller
}

/*
Get /api/category/:id([0-9]+)
category_[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}
*/
func (c *CategoryPost) Get() {
	str := c.Ctx.Input.Param("category_id")
	o := orm.NewOrm()
	category := models.Category{Uid: str}
	error := o.Read(&category, "Uid")
	if error == orm.ErrNoRows {
		c.Data["json"] = models.Message{Code: 301, Result: "分类下没有文章", Data: nil}
		c.ServeJSON()
	}
	var posts []*models.Post
	_, err := o.QueryTable("post").OrderBy("-Id").Filter("categoryId__icontains", &category.Uid).All(&posts)
	if err != nil {
		logs.Info(err)
	}
	var user models.User
	var tag models.Tag
	var postsViews []*models.PostView
	for _, v := range posts {
		_, err = o.QueryTable("user").Filter("uid", v.UserId).All(&user)
		if err != nil {
			logs.Info(err)
		}
		userView := util.UserToViews(user)
		_, err := o.QueryTable("tag").Filter("id", v.Tid).All(&tag)
		if err != nil {
			logs.Info(err)
		}
		tagView := util.TagToView(tag, *userView)
		postView := util.PostToViews(v, *userView, *tagView)

		postsViews = append(postsViews, postView)
	}
	c.Data["json"] = models.Message{Code: 200, Result: "成功", Data: postsViews}
	c.ServeJSON()
}
