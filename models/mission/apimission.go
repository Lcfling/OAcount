package mission

import (
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego/orm"
	"time"
)

type ApiMissionMydata struct {
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

type ApiMissionAreaInfo struct {
	Id       int64
	Name     string
	Realname string
	Phone    int64
}

func (this *ApiMissionMydata) TableName() string {
	return models.TableName("mission_my")
}

//我的任务
func ApiGetMyMission(userId int64, lastid, offset int, types int64) (num int64, err error, ops []MissionMydata) {
	var my []MissionMydata

	qb, _ := orm.NewQueryBuilder("mysql")

	//判断是第几页
	if lastid > 0 {
		qb.Select("t.id", "t.userid", "p.name", "t.missionid", "a.name as areaname", "p.started", "p.ended", "p.desc", "p.creatime", "t.status", "t.check", "t.checktime").From("pms_mission_my AS t").
			LeftJoin("pms_mission AS p").On("p.id = t.missionid").
			LeftJoin("pms_area AS a").On("a.id = t.areaid").
			Where("t.userid=? and t.status=? and t.id<?").OrderBy("-t.`id` ").
			Limit(offset)
		sql := qb.String()
		o := orm.NewOrm()
		nums, err := o.Raw(sql, userId, types, lastid).QueryRows(&my)
		return nums, err, my
	}

	qb.Select("t.id", "t.userid", "p.name", "t.missionid", "a.name as areaname", "p.started", "p.ended", "p.desc", "p.creatime", "t.status", "t.check", "t.checktime").From("pms_mission_my AS t").
		LeftJoin("pms_mission AS p").On("p.id = t.missionid").
		LeftJoin("pms_area AS a").On("a.id = t.areaid").
		Where("t.userid=? and t.status=?").OrderBy("-t.`id` ").
		Limit(offset)
	sql := qb.String()
	o := orm.NewOrm()
	nums, err := o.Raw(sql, userId, types).QueryRows(&my)
	return nums, err, my

}

//任务详情
func ApiGetMissionMy(id int64) ApiMissionMydata {
	var my ApiMissionMydata
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

//点位员信息
func ApiGetAreaUserInfo(owner int64) ApiMissionAreaInfo {
	var my ApiMissionAreaInfo
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("a.name", "a.owner", "p.realname", "p.phone").From("pms_area AS a").
		LeftJoin("pms_users_profile AS p").On("a.owner = p.userid").
		Where("a.owner=?").
		Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql, owner).QueryRow(&my)
	return my
}

//更改我的任务查阅状态
func UpdateCheck(id int64, userid int64) error {
	checktime := time.Now().Unix()
	o := orm.NewOrm()
	_, err := o.Raw("update  "+models.TableName("mission_my")+" SET `check`=1,checktime=?  WHERE id = ? AND userid=?", checktime, id, userid).Exec()
	return err
}
