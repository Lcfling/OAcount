package area

import (
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/area"
	. "github.com/Lcfling/OAcount/models/tags"
	"strconv"
)

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
