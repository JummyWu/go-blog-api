package controllers

import (
	"fmt"
	"go-blog-api/blog-api/models"
	"go-blog-api/blog-api/util"

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
Get http;//localhost:8001/github/restgie/
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
