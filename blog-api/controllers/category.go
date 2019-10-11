package controllers

import (
	"go-blog-api/blog-api/models"

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
		c.Data["json"] = models.Message{Code: 200, Result: "添加分类成功", Data: category}
		c.ServeJSON()
	}
}
