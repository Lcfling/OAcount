package area

import (
	"fmt"
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

//达标率
func GetPassRate(pid, kind int) []Pass {
	area := getAreaListByPid(pid)
	for k, v := range area {
		aids := getAidsByPid(v.Id)
		if aids != nil {
			pass := missionPassRate(aids, kind)
			fmt.Printf("pass:%+v\n", pass)
			area[k].PassRate = pass
			continue
		}

	}
	return area
}

//根据pid获取区域信息
func getAreaListByPid(aid int) []Pass {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("area"))
	cond := orm.NewCondition()
	cond = cond.And("parentid", aid)
	qs = qs.SetCond(cond)
	var areas []Area
	num, _ := qs.All(&areas)
	if num > 0 {
		var areaLists []Pass
		for _, v := range areas {
			var area Pass
			area.Parentid = v.Parentid
			area.Tags = v.Tags
			area.Jstatus = v.Jstatus
			area.Id = v.Id
			area.Name = v.Name
			area.Creatime = v.Creatime
			area.Owner = v.Owner
			area.Coler = v.Coler
			area.Locations = v.Locations
			area.PassRate = 0
			areaLists = append(areaLists, area)
		}
		return areaLists
	} else {
		return nil
	}
}

//根据aid获取一级街道aids where parentid=id
func getAidsByPid(aid int64) []string {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("area"))
	cond := orm.NewCondition()
	cond = cond.And("parentid", aid)
	qs = qs.SetCond(cond)
	var areas []Area
	num, _ := qs.All(&areas, "Id")
	if num > 0 {
		var aids []string
		for _, v := range areas {
			var aid string
			aid = strconv.Itoa(int(v.Id))
			aids = append(aids, aid)
		}
		return aids
	} else {
		return nil
	}
}

//根据类型、aids计算达标率  count(all)=where aid in aids and type=1||type=2 and where arraignment=2
func missionPassRate(aids []string, kind int) float64 {
	var none, done int64
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("mission_my"))
	cond1 := orm.NewCondition()
	cond1 = cond1.And("types", kind)      //类型
	cond1 = cond1.And("areaid__in", aids) //aids数组
	none, _ = qs.SetCond(cond1).Count()

	cond2 := orm.NewCondition()
	cond2 = cond2.And("types", kind)    //类型
	cond2 = cond2.And("arraignment", 2) //审核通过
	cond2 = cond2.And("areaid__in", aids)
	done, _ = qs.SetCond(cond2).Count()

	if none == 0 {
		return 0
	}
	doneRate := float64(done) / float64(none) * 100
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", doneRate), 64)

	return value
}
