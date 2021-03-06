package mission

import (
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/mission"
	"github.com/astaxie/beego"
	"strconv"
)

//点位信息
type ApiAreaInfoController struct {
	controllers.IndexController
}

//点位信息
func (this *ApiAreaInfoController) Post() {
	//userid := this.UserBaseController.UserUserId
	aid := this.GetString("aid")
	aid64, _ := strconv.ParseInt(aid, 10, 64)
	//点位员信息
	areaInfo := ApiGetAreaInfo(aid64)
	areaInfo.MissionCount = MissionCount(areaInfo.Owner)    //任务数量
	areaInfo.MissionOver = MissionOverCount(areaInfo.Owner) //任务完成数量
	//返回数据
	data := make(map[string]interface{})
	data["areaInfo"] = areaInfo
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "消息列表", "data": data}
	this.ServeJSON()
}

//我的任务列表
type ApiMissionMyController struct {
	controllers.UserBaseController
}

//我的任务列表
func (this *ApiMissionMyController) Post() {

	userid := this.UserBaseController.UserUserId
	types := this.GetString("types") // 0 未完成  1已完成
	if types == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "参数错误!", "data": ""}
		this.ServeJSON()
	}
	types64, _ := strconv.ParseInt(types, 10, 64)
	lastid, err := this.GetInt("lastid")
	if err != nil {
		lastid = 0
	}
	// 每页显示数量
	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	//获取我的任务
	_, _, missionmy := ApiGetMyMission(userid, lastid, offset, types64)
	//返回数据
	data := make(map[string]interface{})
	data["missionmy"] = missionmy
	data["countmissionmy"] = len(missionmy)
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "消息列表", "data": data}
	this.ServeJSON()
}

//任务详情
type ApiMissionInfoController struct {
	controllers.UserBaseController
}

//任务详情
func (this *ApiMissionInfoController) Post() {

	userid := this.UserBaseController.UserUserId
	missionId := this.GetString("missionId") // 任务id
	missionId64, _ := strconv.ParseInt(missionId, 10, 64)
	if !(missionId64 > int64(0)) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "参数错误!", "data": ""}
		this.ServeJSON()
	}
	//获取任务详情
	missionInfo := ApiGetMissionMy(missionId64)

	//更改任务查阅状态
	go UpdateCheck(missionId64, int64(userid))

	//返回数据
	data := make(map[string]interface{})
	data["missionInfo"] = missionInfo
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "任务详情", "data": data}
	this.ServeJSON()
}
