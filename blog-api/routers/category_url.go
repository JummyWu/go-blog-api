package routers

import (
	"go-blog-api/blog-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/category/add_category", &controllers.NewCategory{})
	beego.Router("/api/category/update_category/", &controllers.UpdateCategory{})
	beego.Router("/api/category/delete_category", &controllers.DeleteCategory{})
	beego.Router("/api/category/list", &controllers.CategoryList{})
	beego.Router("/api/category/:category_idcategory_[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}", &controllers.CategoryPost{})
}
