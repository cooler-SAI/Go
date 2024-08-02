package routers

import (
	"BeeGo/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.HelloController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/user/:id", &controllers.UserController{})

	beego.Router("/api/v1/user", &controllers.UserAPIController{}, "get:GetAll;post:Create")
	beego.Router("/api/v1/user/:id", &controllers.UserAPIController{}, "get:Get;put:Update;delete:Delete")
}
