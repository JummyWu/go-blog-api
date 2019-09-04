package main

import (
	_ "go-blog-api/blog-api/models"
	_ "go-blog-api/blog-api/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
