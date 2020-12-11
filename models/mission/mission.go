package mission

import (
	"fmt"
	"github.com/Lcfling/OAcount/models"
	"github.com/Lcfling/OAcount/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Mission struct {
	Id       int64
	Userid   int64
	Name     string
	Types    int
	Mid      int64
	Started  int64
	Ended    int64
	Desc     string
	Creatime int64
	Status   int
}
type MissionMydata struct {
	Id        int64
	Userid    int64
	Name      string
	Areaname  string
	Missionid int64
	Started   int64
	Ended     int64
	Desc      string
	Feedback  string
	Creatime  int64
	Status    int
	Check     int
	Checktime int64
}
type MissionMy struct {
	Id        int64
	Missionid int64
	Types     int64
	Userid    int64
	Areaid    int64
	Check     int
	Creatime  int64
	Checktime int64
	Feedback  string
	Detail    string
	Status    int
}

func (this *Mission) TableName() string {
	return models.TableName("mission")
}
func (this *MissionMy) TableName() string {
	return models.TableName("mission_my")
}
func init() {
	orm.RegisterModel(new(Mission), new(MissionMy))
}
func GetMission(id int64) (Mission, error) {
	var mission Mission
	var err error

	//err = utils.GetCache("GetProject.id."+fmt.Sprintf("%d", id), &project)
	//if err != nil {
	o := orm.NewOrm()
	mission = Mission{Id: id}
	err = o.Read(&mission)
	//utils.SetCache("GetProject.id."+fmt.Sprintf("%d", id), project, 600)
	//}
	return mission, err
}
func GetMissionName(id int64) string {
	var err error
	var name string

	err = utils.GetCache("GetMissionName.id."+fmt.Sprintf("%d", id), &name)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var mission Mission
		o := orm.NewOrm()
		o.QueryTable(models.TableName("mission")).Filter("id", id).One(&mission, "name")
		name = mission.Name
		utils.SetCache("GetMissionName.id."+fmt.Sprintf("%d", id), name, cache_expire)
	}
	return name
}
func AddMission(updPro Mission) (int64, error) {
	o := orm.NewOrm()
	pro := new(Mission)

	//pro.Id = updPro.Id
	pro.Userid = updPro.Userid
	pro.Name = updPro.Name
	pro.Types = updPro.Types
	pro.Started = updPro.Started
	pro.Ended = updPro.Ended
	pro.Desc = updPro.Desc
	pro.Creatime = time.Now().Unix()
	pro.Status = 1
	id, err := o.Insert(pro)
	return id, err
}

//项目列表
func ListMission(condArr map[string]string, page int, offset int) (num int64, err error, user []Mission) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("mission"))
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("name__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	qs = qs.RelatedSel()

	var missions []Mission
	qs = qs.OrderBy("-id")
	num, err1 := qs.Limit(offset, start).All(&missions)
	return num, err1, missions
}

//统计数量
func CountMission(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("mission"))
	qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("name__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}
func ChangeMissionStatus(id int64, status int) error {
	o := orm.NewOrm()

	pro := Mission{Id: id}
	err := o.Read(&pro, "id")
	if nil != err {
		return err
	} else {
		pro.Status = status
		_, err := o.Update(&pro)
		return err
	}
}
func GetMyMission(userId int64, page int, offset int) (num int64, err error, ops []MissionMydata) {
	var my []MissionMydata
	start := (page - 1) * offset
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("t.id", "t.userid", "p.name", "t.missionid", "a.name as areaname", "p.started", "p.ended", "p.desc", "p.creatime", "t.status", "t.check", "t.checktime").From("pms_mission_my AS t").
		LeftJoin("pms_mission AS p").On("p.id = t.missionid").
		LeftJoin("pms_area AS a").On("a.id = t.areaid").
		Where("t.userid=?").
		Limit(offset).Offset(start)
	sql := qb.String()
	o := orm.NewOrm()
	nums, err := o.Raw(sql, userId).QueryRows(&my)
	return nums, err, my
}
func GetMissionMy(id int64) MissionMydata {
	var my MissionMydata
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("p.id", "t.userid", "p.name", "a.name as areaname", "p.started", "p.ended", "p.desc", "p.creatime", "t.status", "t.check", "t.checktime").From("pms_mission_my AS t").
		LeftJoin("pms_mission AS p").On("p.id = t.missionid").
		LeftJoin("pms_area AS a").On("a.id = t.areaid").
		Where("t.id=?").
		Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql, id).QueryRow(&my)
	return my
}
func UpdateMissionMy(m MissionMy, id int64) error {
	var my MissionMy

	o := orm.NewOrm()
	my = MissionMy{Id: id}
	o.Read(&my)
	my.Feedback = m.Feedback
	my.Detail = m.Detail
	_, err := o.Update(&my)
	return err
}

//添加我的任务
func AddMyMission(
	Missionid, Userid, Areaid int64) (int64, error) {
	o := orm.NewOrm()
	MissionMy := new(MissionMy)
	MissionMy.Missionid = Missionid
	MissionMy.Userid = Userid
	MissionMy.Areaid = Areaid
	MissionMy.Creatime = time.Now().Unix()
	MissionMy.Status = 0
	id, err := o.Insert(MissionMy)
	return id, err
}
