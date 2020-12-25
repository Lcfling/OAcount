package area

import (
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/area"
	"strconv"
)

//打卡数据
type DakaController struct {
	controllers.IndexController
}

//区域任务成功率
func (this *DakaController) Post() {
	mid := this.GetString("missionid") //区域ID
	missionId, _ := strconv.Atoi(mid)
	area := GetAreaDaka(missionId)
	//返回数据
	data := make(map[string]interface{})
	data["area"] = area
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "点位打卡数据", "data": data}
	this.ServeJSON()
}
