package controllers

import (
	//"github.com/virteman/OPMS/initial"

	"github.com/astaxie/beego"
)

type MobileController struct {
	beego.Controller
	IsLogin bool
	//UserInfo string
	UserUserId   int64
	UserUsername string
	UserAvatar   string
}


//hocker
func (this *MobileController) Prepare() {
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.IsLogin = false
	} else {
		this.IsLogin = true
	}
	this.Data["IsLogin"] = this.IsLogin
	//this.Data["IsLogin"] = this.IsLogin
}
