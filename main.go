package main

import (
	"encoding/gob"
	"github.com/jicg/liteblog/models"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/jicg/liteblog/models"
	_ "github.com/jicg/liteblog/routers"
)

func main() {
	initSession()
	initTemplate()
	beego.Run()

}

func initSession() {
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

//添加初始化模板方法
func initTemplate() {
	//访问路径对比
	beego.AddFuncMap("equrl", func(x, y string) bool {
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1, y1) == 0
	})

	beego.AddFuncMap("add", func(x,y int) int {
		return x + y
	})
}
