package area

import (
	"fmt"
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

//根据区域获取area id数组
func GetUidsByPid(pid int) []string {
	AreaList := GetChild(int64(pid), "")
	aids := GetAllAreaIdByPid(AreaList)
	return aids
}

//区域、行业都选，获取交集area id 数组
func GetAidsBoth(aidsByPid, aidsByTid []string) []string {
	both := []string{}
	for _, v1 := range aidsByPid {
		canAdd := false
		for _, v2 := range aidsByTid {
			if v1 == v2 {
				canAdd = true
			}
		}
		if canAdd {
			both = append(both, v1)
		}
	}
	return both
}

//根据area id数组 获取userid 数组
func GetUidsByAids(aids []string) []string {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("area"))
	cond := orm.NewCondition()
	cond = cond.And("id__in", aids)
	qs = qs.SetCond(cond)
	var areas []Area
	num, _ := qs.All(&areas)
	fmt.Println("num", num)
	if num > 0 {
		var uids []string
		for _, v := range areas {
			var uid string
			uid = strconv.Itoa(int(v.Owner))
			uids = append(uids, uid)
		}
		return uids
	} else {
		return nil
	}

}

//获取任务成功率
func GetMissionDoneRate(uids []string) float64 {
	var none, done int64
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("mission_my"))
	cond1 := orm.NewCondition()
	cond1 = cond1.And("userid__in", uids)
	none, _ = qs.SetCond(cond1).Count()

	cond2 := orm.NewCondition()
	cond2 = cond2.And("status", 1) //已完成
	cond2 = cond2.And("userid__in", uids)
	done, _ = qs.SetCond(cond2).Count()

	doneRate := float64(done) / float64(none) * 100
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", doneRate), 64)
	return value

}
