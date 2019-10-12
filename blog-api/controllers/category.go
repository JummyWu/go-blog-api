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
		_, err = o.QueryTable("user").Filter("Uid", uid).All(user)
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
		categoryId := c.GetString("categoryId")
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
