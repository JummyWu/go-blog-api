package routers

import (
	"go-blog-api/blog-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/links/add_links", &controllers.LikesController{})
	beego.Router("/api/links/update_links", &controllers.LikesController{})
	beego.Router("/api/links/update_status", &controllers.LikesController{})
	beego.Router("/api/links/delete_links", &controllers.LikesController{})
	beego.Router("/api/links_list", &controllers.LikesController{})
}
