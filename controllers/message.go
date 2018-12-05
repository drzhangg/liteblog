package controllers

import (
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserror"
)

type MessageController struct {
	BaseController
}

//@router /new/:key [post]
func (this *MessageController) NewMessage() {
	this.MustLogin()
	key := this.Ctx.Input.Param(":key")
	content := this.GetMustString("content","请输入内容！")
	k := this.UUID()
	m := &models.Message{
		Key:k,
		NoteKey:key,
		User:this.User,
		UserId:int(this.User.ID),
		Content:content,
	}
	if err := models.SaveMessage(m); err != nil {
		this.Abort500(syserror.New("保存失败",err))
	}

	this.JsonOkH("保存成功！",nil)
}
