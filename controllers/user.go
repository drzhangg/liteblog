package controllers

import (
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserror"
	"github.com/pkg/errors"
	"strings"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (this *UserController) Login() {

	//email
	email := this.GetMustString("email","邮箱不能为空！")
	//password
	pwd := this.GetMustString("password","密码不能为空！")

	user,err := models.QueryUserByEmailAndPwd(email,pwd)
	if err != nil {
		this.Abort500(syserror.New("登录失败！",err))
	}

	this.SetSession(SESSION_USER_KEY,user)
	this.Data["json"] = map[string]interface{}{
		"code":0,
		"action":"/",
	}
	this.ServeJSON()

}

// @router /logout [get]
func (this *UserController) Logout() {
	this.MustLogin()
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/",302)
}

// @router /reg [post]
func (this *UserController) Reg() {
	name := this.GetMustString("name","昵称不能为空")
	email := this.GetMustString("email","邮箱不能为空")
	password := this.GetMustString("password","密码不能为空")
	password1 := this.GetMustString("password1","确认不能为空")

	if strings.Compare(password, password1) != 0 {
		this.Abort500(errors.New("两次输入的密码不一致！"))
	}

	if u,err := models.QueryUserByName(name); err == nil && u.ID > 0{
		this.Abort500(errors.New("用户昵称已存在"))
	}

	if u,err := models.QueryUserByEmail(email); err == nil && u.ID > 0{
		this.Abort500(errors.New("用户邮箱已存在"))
	}

	if err := models.SaveUser(&models.User{
		Name:name,
		Email:email,
		Pwd:password,
		Avatar:"/static/image/info-img.png",
		Role:1,
	}); err != nil {
		this.Abort500(syserror.New("用户保存失败",err))
	}
	this.JsonOk("注册成功","/user")
}
