package news

import (
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/news"
	"github.com/astaxie/beego"
)

type ApiNewsClassicController struct {
	controllers.UserBaseController
}

type ApiNewsController struct {
	controllers.UserBaseController
}

//消息类型
func (this *ApiNewsClassicController) Post() {

	//获利消息类型
	Classic := ApiGetClassic()
	//返回数据
	data := make(map[string]interface{})
	data["classic"] = Classic
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "消息列表", "data": data}
	this.ServeJSON()
}

//消息列表
func (this *ApiNewsController) Post() {

	lastid, err := this.GetInt("lastid")
	//status := this.GetString("status")
	classid := this.GetString("classid")
	if err != nil {
		lastid = 0
	}

	// 每页显示数量
	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}
	condArr := make(map[string]string)
	//condArr["status"] = status
	condArr["classid"] = classid
	//countProject := CountNews(condArr)

	//类型消息
	_, _, newsList := ApiGetPageNews(condArr, lastid, offset)

	//返回数据
	data := make(map[string]interface{})
	//data["paginator"]=paginator
	data["newsList"] = newsList

	//data["countNews"]=countProject
	//this.Data["condArr"] = condArr

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "消息列表", "data": data}

	this.ServeJSON()
}
