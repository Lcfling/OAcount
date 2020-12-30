package news

import (
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/news"
	"github.com/astaxie/beego"
	"strconv"
)

//首页轮播
type ApiBannerClassicController struct {
	controllers.MobileController
}

//首页轮播
func (this *ApiBannerClassicController) Post() {
	//消息类型
	Banner := ApiGetClassic()
	//返回数据
	data := make(map[string]interface{})
	data["Banner"] = Banner
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "首页轮播", "data": data}
	this.ServeJSON()
}

//消息类型
type ApiNewsClassicController struct {
	controllers.MobileController
}

//消息类型
func (this *ApiNewsClassicController) Post() {
	//消息类型
	Classic := ApiGetClassic()
	//返回数据
	data := make(map[string]interface{})
	data["classic"] = Classic
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "消息类型", "data": data}
	this.ServeJSON()
}

//消息列表
type ApiNewsController struct {
	controllers.MobileController
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
	condArr["classid"] = classid

	//类型消息
	_, _, newsList := ApiGetPageNews(condArr, lastid, offset)
	//返回数据
	data := make(map[string]interface{})
	//data["paginator"]=paginator
	data["newsList"] = newsList
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "消息列表", "data": data}

	this.ServeJSON()
}

// 任务单位
type ApiAreaController struct {
	controllers.MobileController
}

//单位列表
func (this *ApiAreaController) Post() {

	//类型消息
	ApiGetTags := ApiGetTags()
	//返回数据
	data := make(map[string]interface{})
	//data["paginator"]=paginator
	data["ApiGetTags"] = ApiGetTags
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "单位列表", "data": data}
	this.ServeJSON()
}

//任务列表
type ApiMissionController struct {
	controllers.MobileController
}

//任务列表
func (this *ApiMissionController) Post() {
	tid := this.GetString("tid")
	tid64, _ := strconv.ParseInt(tid, 10, 64)
	//类型消息
	_, ApiTageFile := TagsFile(tid64)
	//返回数据
	data := make(map[string]interface{})
	data["ApiTageFile"] = ApiTageFile
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "任务列表", "data": data}

	this.ServeJSON()
}

//任务详情
type ApiMissionInfoController struct {
	controllers.MobileController
}

//任务详情
func (this *ApiMissionInfoController) Post() {
	tid := this.GetString("tid")
	tid64, _ := strconv.ParseInt(tid, 10, 64)
	//类型消息
	_, mission := ApiGetMissionInfo(tid64)
	//返回数据
	data := make(map[string]interface{})
	data["mission"] = mission
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "任务详情", "data": data}

	this.ServeJSON()
}
