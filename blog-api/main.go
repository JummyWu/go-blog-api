package main

import (
	"encoding/json"
	"go-blog-api/blog-api/models"
	_ "go-blog-api/blog-api/models"
	_ "go-blog-api/blog-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	var filterAdmin = func(c *context.Context) {
		url := c.Input.URL()
		logs.Info("filter url : %s", url)
		token := c.Request.Header.Get("token")
		if token == "" {
			c.Output.SetStatus(403)
			m := models.Message{Code: 403, Result: "你没登陆", Data: nil}
			resBody, err := json.Marshal(m)
			c.Output.Body([]byte(resBody))
			if err != nil {
				logs.Info(err)
			}
		}
		if url != "/api.user.admin.login" {
			claims := models.ParseToken(token)
			if claims == false {
				c.Output.SetStatus(403)
				m := models.Message{Code: 403, Result: "你的token已经过期了,请重新登陆", Data: nil}
				resBody, err := json.Marshal(m)
				c.Output.Body([]byte(resBody))
				if err != nil {
					logs.Info(err)
				}
			}
		}
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/api/post/**", beego.BeforeExec, filterAdmin)
	beego.InsertFilter("/api/tag/**", beego.BeforeExec, filterAdmin)
	beego.InsertFilter("/api/links/**", beego.BeforeExec, filterAdmin)
	beego.Run()
}
