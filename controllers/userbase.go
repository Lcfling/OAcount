package controllers

import (
	//"github.com/virteman/OPMS/initial"

	. "github.com/Lcfling/OAcount/models/users"
	"github.com/astaxie/beego"
	"strconv"
)

type UserBaseController struct {
	beego.Controller
	IsLogin bool
	//UserInfo string
	UserUserId   int64
	UserUsername string
	UserAvatar   string
}

func (this *UserBaseController) Prepare() {
	token := this.Ctx.Request.Header.Get("token")
	useridstr := this.Ctx.Request.Header.Get("userid")
	if token == "" || useridstr == "" {
		this.Data["json"] = map[string]interface{}{"code": 2, "message": "登录效验失败", "data": ""}
		this.ServeJSON()
	}
	userid, _ := strconv.ParseInt(useridstr, 10, 64)
	U, err := GetUser(userid)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 2, "message": "登录效验失败", "data": ""}
		this.ServeJSON()
	} else {
		if U.Token != token {
			this.Data["json"] = map[string]interface{}{"code": 2, "message": "登录效验失败", "data": ""}
			this.ServeJSON()
		} else {
			this.UserUserId = U.Id
			this.UserUsername = U.Username
			this.UserAvatar = U.Avatar
		}
	}

}
