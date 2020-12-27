package area

import (
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/area"
	. "github.com/Lcfling/OAcount/models/tags"
	"strconv"
)

//任务成功率
type ApiDoneRateController struct {
	controllers.IndexController
}

//区域任务成功率
func (this *ApiDoneRateController) Post() {
	pid := this.GetString("pid") //区域ID
	tid := this.GetString("tid") //行业ID
	aid, _ := strconv.Atoi(pid)
	tagId, _ := strconv.Atoi(tid)

	var aids []string
	var uids []string
	var rate float64
	//返回数据
	data := make(map[string]interface{})
	if pid == "" && tid == "" {
		rate = 0
		data["rate"] = rate
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "区域成功率", "data": data}
		this.ServeJSON()
	}
	//未选区域
	if pid == "" {
		aids = GetUidsByTid(int64(tagId))
	}
	//未选行业
	if tid == "" {
		aids = GetUidsByPid(aid)
	}
	//区域、行业都选
	if pid != "" && tid != "" {
		aidsByPid := GetUidsByPid(aid)
		aidsByTid := GetUidsByTid(int64(tagId))
		aids = GetAidsBoth(aidsByPid, aidsByTid)
	}

	uids = GetUidsByAids(aids)
	if uids == nil {
		rate = 0
	} else {
		rate = GetMissionDoneRate(uids)
	}
	//返回数据
	data["rate"] = rate
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "区域成功率", "data": data}
	this.ServeJSON()
}

//任务达标率
type ApiPassRateController struct {
	controllers.IndexController
}

//区域任务达标率率
func (this *ApiPassRateController) Post() {
	pid := this.GetString("pid")  //区域上级ID
	tid := this.GetString("type") //任务类型
	parentId, _ := strconv.Atoi(pid)
	kind, _ := strconv.Atoi(tid)

	area := GetPassRate(parentId, kind)

	//返回数据
	data := make(map[string]interface{})
	data["pass"] = area
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "任务达标率", "data": data}
	this.ServeJSON()
}

//问卷调查
type ApiQuestionController struct {
	controllers.IndexController
}

//区域 问卷平均分
func (this *ApiQuestionController) Post() {
	aid := this.GetString("pid") //区域D
	areaid, _ := strconv.Atoi(aid)
	area := GetAreaQuestionAver(areaid)
	//返回数据
	data := make(map[string]interface{})
	data["question"] = area
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "问卷平均分", "data": data}
	this.ServeJSON()
}
