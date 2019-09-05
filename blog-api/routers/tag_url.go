package routers

import (
	"go-blog-api/blog-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/tag/add_tag", &controllers.NewTag{})
	beego.Router("/api/tag/update_tag", &controllers.UpdateTag{})
	beego.Router("/api/tag/delete_tag", &controllers.DeleteTag{})
	beego.Router("/api/tag_list", &controllers.GetTagList{})
	beego.Router("/api/tag_posts/:id([0-9]+)", &controllers.GetTagPosts{})
}
