package routers

import (
	"W2OlineWinterAssignmentTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//-----主页部分-----
    beego.Router("/", &controllers.HomeController{})
	beego.Router("/index", &controllers.HomeController{})
	beego.Router("/exit",&controllers.ExitController{})
	//-----登录注册部分-----
	beego.Router("/register",&controllers.RegisterController{})
    beego.Router("/login",&controllers.LoginController{})
	//-----书籍部分-----
	beego.Router("/bookinfo/:id",&controllers.ShowBookController{})
    beego.Router("/tags",&controllers.TagsController{})
	beego.Router("/bookshelf",&controllers.UserBookShelfController{})
    beego.Router("/delete",&controllers.DeleteController{})
    beego.Router("/addfavo/:id",&controllers.AddFavoriteBookController{})
}
