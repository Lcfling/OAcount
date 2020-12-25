package area

import (
	"github.com/Lcfling/OAcount/models"
	. "github.com/Lcfling/OAcount/models/checkworks"
	"github.com/astaxie/beego/orm"
	"time"
)

//获取今天区域打卡数据
func GetAreaDaka() []Daka {
	area := getArea()
	daka := getTodayDaka()

	for _, av := range area {
		for ak, dv := range daka {
			if dv.Aid == av.Id {
				area[ak].Daka = 1
				continue
			}
		}
	}

	return area
}

//获取所有区域信息
func getArea() []Daka {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("area"))
	cond := orm.NewCondition()
	qs = qs.SetCond(cond)
	var areas []Area
	num, _ := qs.All(&areas) //, "Id", "Parentid", "Jstatus", "Name", "Tags", "Locations", "Owner", "Coler", "Creatime")
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
