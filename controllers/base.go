package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserror"
	"github.com/satori/go.uuid"
)

const SESSION_USER_KEY  = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	User models.User
	IsLogin bool
}

type NestPreparer interface {
	NextPrepare()
}


func (this *BaseController) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	u,ok := this.GetSession(SESSION_USER_KEY).(models.User)
	this.IsLogin = false
	if ok {
		this.User = u
		this.IsLogin = true
		this.Data["User"] = this.User
	}
	this.Data["islogin"] = this.IsLogin
	if a,ok := this.AppController.(NestPreparer);ok{
		a.NextPrepare()
	}
}

func (this *BaseController)Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}

func (this *BaseController) GetMustString(key, msg string) string {
	val := this.GetString(key)
	if len(val) == 0 {
		this.Abort500(errors.New(msg))
	}
	return val
}

func (this *BaseController) MustLogin() {
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
}

type H map[string]interface{}

func (this *BaseController) JsonOk(msg,action string) {
	this.Data["json"] = H{
		"code":0,
		"msg":msg,
		"action":action,
	}
	this.ServeJSON()
}

func (this *BaseController) JsonOkH(msg string, data H) {
	if data == nil {
		data = H{}
	}
	data["code"] = 0
	data["msg"] = msg
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *BaseController) UUID() string {
	u,err := uuid.NewV4()
	if err != nil {
		this.Abort500(syserror.New("系统错误",err))
	}
	return u.String()
}