package controllers

import (
	"fmt"
	"go-blog-api/blog-api/models"
	"go-blog-api/blog-api/util"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

/*
UserGitHub 请求GitHub
*/
type UserGitHub struct {
	beego.Controller
}

/*
Get github/restgie/
*/
func (c *UserGitHub) Get() {
	code := c.GetString("code")
	token := util.GetGitHubToken(code)
	if token == "" {
		c.Data["json"] = models.Message{Code: 200, Result: "获取access_token失败", Data: nil}
		c.ServeJSON()
	}
	userInfo := util.GetGitHubUser(token)
	uid := userInfo.Uid
	fmt.Println("这是一个用户的ID" + uid)
	jwtToken, err := models.SetToken(uid)
	if err != nil {
		logs.Info(err)
		c.Data["json"] = models.Message{Code: 301, Result: "获取token失败", Data: err}
		c.ServeJSON()
	}
	user := map[string]interface{}{"token": jwtToken, "user": userInfo}
	if userInfo == nil {
		c.Data["json"] = models.Message{Code: 200, Result: "写入失败", Data: user}
		c.ServeJSON()
	}
	c.Data["json"] = models.Message{Code: 200, Result: "获取用户信息成功", Data: user}
	c.ServeJSON()
}

/*
UpdatePassword 修改密码
*/
type UpdatePassword struct {
	beego.Controller
}

/*
Put update/password
*/
func (c *UpdatePassword) Put() {
	token := c.Ctx.Request.Header.Get("token")
	uid := models.GetUID(token)
	pard := c.GetString("password")
	o := orm.NewOrm()
	user := models.User{Uid: uid}
	err := o.Read(&user, "Uid")
	if err == orm.ErrNoRows {
		logs.Info(err)
		c.Data["json"] = models.Message{Code: 200, Result: "没有这个用户", Data: nil}
		c.ServeJSON()
	}
	password, err := models.Encrypt(pard)
	if err != nil {
		logs.Info(err)
		c.Data["json"] = models.Message{Code: 301, Result: "加密错误", Data: nil}
		c.ServeJSON()
	}
	user.Password = password
	logs.Info(o.Update(&user, "Password"))
	c.Data["json"] = models.Message{Code: 200, Result: "修改成功", Data: nil}
	c.ServeJSON()
}

/*
Logout 退出登陆
*/
type Logout struct {
	beego.Controller
}

/*
Post /user/post
*/
func (c *Logout) Post() {
	token := c.Ctx.Request.Header.Get("token")
	uid := models.GetUID(token)
	o := orm.NewOrm()
	user := models.User{Uid: uid}
	err := o.Read(&user, "Uid")
	if err == orm.ErrNoRows {
		logs.Info(err)
		c.Data["json"] = models.Message{Code: 200, Result: "没有这个用户", Data: nil}
		c.ServeJSON()
	}
	c.Data["json"] = models.Message{Code: 200, Result: "已经退出", Data: nil}
}
