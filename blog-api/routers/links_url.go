package routers

import (
	"go-blog-api/blog-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/links/add_links", &controllers.NewLinks{})
	beego.Router("/api/links/update_links", &controllers.UpdateLinkes{})
	beego.Router("/api/links/update_status", &controllers.UpdateStatus{})
	beego.Router("/api/links/delete_links", &controllers.DeleteLinks{})
	beego.Router("/api/links_list", &controllers.GetLinks{})
}
