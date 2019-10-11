package routers

import (
	"go-blog-api/blog-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/post/comment/", &controllers.CommentController{}, "post:PostComment")
	beego.Router("/api/post/comment_reply/", &controllers.CommentController{}, "post:PostCommentReply")
	beego.Router("/api/post_comment/", &controllers.CommentController{}, "get:Get")
}
