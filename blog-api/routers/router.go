package routers

import (
	"go-blog-api/blog-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/github/redirect", &controllers.UserGitHub{})
	beego.Router("/api/post/add_post", &controllers.NewPost{})
	beego.Router("/api/post/update_post", &controllers.UpdatePost{})
	beego.Router("/api/post/delete_post", &controllers.DeletePost{})
	beego.Router("/api/post_list", &controllers.PostList{})
	beego.Router("/api/post_list/page/:id([0-9]+)", &controllers.PostList{})
	beego.Router("/api/post_id/:id([0-9]+)", &controllers.PostController{})
}
