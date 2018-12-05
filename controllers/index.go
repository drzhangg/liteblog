package controllers

import (
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserror"
)

type IndexController struct {
	BaseController
}

// @router / [get]
func (this *IndexController) Get() {
	limit := 10

	//page
	page,err := this.GetInt("page",1)
	if err != nil || page <= 0 {
		page = 1
	}

	//得到页面的数据
	notes,_ := models.QueryNoteByPage(page,limit)
	this.Data["notes"] = notes

	//总页数
	count,err := models.QueryNotesCount()
	if err != nil {
		this.Abort500(err)
	}
	totpage := count / limit
	if count%limit != 0 {
		totpage = totpage + 1
	}

	this.Data["totpage"] = totpage
	this.Data["page"] = page
	this.TplName = "index.html" //首页模板
}

// @router /about [get]
func (this *IndexController) GetAbort() {
	this.TplName = "about.html"
}

// @router /message [get]
func (this *IndexController) GetMessage() {
	//this.Abort("404")
	this.TplName = "message.html"
}

// @router /user [get]
func (this *IndexController) GetUser() {
	this.TplName = "user.html"
}

// @router /reg [get]
func (this *IndexController) GetReg() {
	this.TplName = "reg.html"
}

// @router /comment/:key [get]
func (this *IndexController) GetComment() {
	key := this.Ctx.Input.Param(":key")
	note,err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.New("文章不存在",err))
	}
	this.Data["note"] = note
	this.TplName = "comment.html"
}

//@router /details/:key [get]
func (this *IndexController) GetDetails()  {
	key:= this.Ctx.Input.Param(":key")
	note,err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.New("文章不存在",err))
	}
	this.Data["note"] = note

	ms,err := models.QueryMessageByNoteKey(key)
	if err != nil {
		this.Abort500(syserror.New("文章不存在",err))
	}
	this.Data["messages"] = ms
	this.TplName = "details.html"
}