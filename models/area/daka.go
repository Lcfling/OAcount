package area

import (
	"github.com/Lcfling/OAcount/models"
	. "github.com/Lcfling/OAcount/models/checkworks"
	. "github.com/Lcfling/OAcount/models/mission"
	"github.com/astaxie/beego/orm"
	"time"
)

//获取今天区域打卡数据
func GetAreaDaka(mid int) []Daka {
	var newArea []Daka        //最后的数据
	area := getArea()         //全部数据
	daka := getTodayDaka()    //打卡数据
	aids := getAidsByMid(mid) //我的任务获取aids
	for _, av1 := range area {
		add := false
		for _, v2 := range aids {
			if v2 == av1.Id {
				add = true
			}
		}
		if add {
			newArea = append(newArea, av1)
		}
	}
	//根据aids详细区域点位，获取打卡数据
	for _, av := range newArea {
		for ak, dv := range daka {
			if dv.Aid == av.Id {
				newArea[ak].Daka = 1
				continue
			}
		}
	}
	return newArea
}

//根据missionID 获取aid 数组
func getAidsByMid(mid int) []int64 {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("mission_my"))
	cond := orm.NewCondition()
	cond = cond.And("missionid", mid)
	qs = qs.SetCond(cond)
	var areas []MissionMy
	num, _ := qs.All(&areas, "Areaid")
	if num > 0 {
		var aids []int64
		for _, v := range areas {
			var aid int64
			aid = v.Areaid
			aids = append(aids, aid)
		}
		return aids
	} else {
		return nil
	}
}

//获取所有区域信息
func getArea() []Daka {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("area"))
	cond := orm.NewCondition()
	qs = qs.SetCond(cond)
	var areas []Area
	num, _ := qs.All(&areas)
	if num > 0 {
		var areaLists []Daka
		for _, v := range areas {
			var area Daka
			area.Parentid = v.Parentid
			area.Tags = v.Tags
			area.Jstatus = v.Jstatus
			area.Id = v.Id
			area.Name = v.Name
			area.Creatime = v.Creatime
			area.Owner = v.Owner
			area.Coler = v.Coler
			area.Locations = v.Locations
			area.Daka = 0
			areaLists = append(areaLists, area)
		}
		return areaLists
	} else {
		return nil
	}
}

//获取今天所有打卡数据
func getTodayDaka() []Checkworks {
	begin, end := getToday()
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("checkworks"))
	cond := orm.NewCondition()
	cond = cond.And("Created__gte", begin)
	cond = cond.And("Created__lte", end)
	qs = qs.SetCond(cond)
	var dakas []Checkworks
	num, _ := qs.All(&dakas)
	if num > 0 {
		var dakaLists []Checkworks
		for _, v := range dakas {
			var daka Checkworks
			daka.Id = v.Id
			daka.Userid = v.Userid
			daka.Clock = v.Clock
			daka.Aid = v.Aid
			daka.Missionmyid = v.Missionmyid
			daka.Lng = v.Lng
			daka.Lat = v.Lat
			daka.Type = v.Type
			daka.Ip = v.Ip
			daka.Created = v.Created
			dakaLists = append(dakaLists, daka)
		}
		return dakaLists
	} else {
		return nil
	}
}

//00:00 23:59 今天时间戳
func getToday() (int64, int64) {
	const timeLayout = "20060102150405"
	t := time.Now()
	pre := t.Format(timeLayout)
	suf := []rune(pre)
	dateT := string(suf[0:8])
	start := dateT + "000000"
	end := dateT + "235900"
	startTime, _ := time.ParseInLocation(timeLayout, start, time.Local)
	endTime, _ := time.ParseInLocation(timeLayout, end, time.Local)
	return startTime.Unix(), endTime.Unix()
}
