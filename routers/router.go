package routers

import (
	"github.com/astaxie/beego"
	"github.com/jicg/liteblog/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.UserController{})
	beego.AddNamespace(
		beego.NewNamespace("note",
			beego.NSInclude(&controllers.NoteController{}),
		),
		beego.NewNamespace("message",
			beego.NSInclude(&controllers.MessageController{})),
	)

}
