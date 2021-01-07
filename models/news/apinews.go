package news

import (
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//消息表
type ApiNews struct {
	Id       int64
	Classid  int64
	Title    string
	Content  string
	Creatime int64
}

//类型表
type ApiClassic struct {
	Id        int64
	Classname string
	Parentid  int64
	Sorts     int64
}

//单位表
type ApiArea struct {
	Id       int64
	Parentid int64
	Name     string
	Owner    int64
}

//社区表
type ApiTags struct {
	Id   int64
	Name string
}

//任务表
type ApiMission struct {
	Id        int64
	Missionid int64
	Types     int64
	Userid    int64
	Areaid    int64
	Name      string
	Desc      string
	Creatime  int64
}

func (this *ApiNews) TableName() string {
	return models.TableName("news")
}

//分页获取消息类型列表
func ApiGetPageNews(condArr map[string]string, lastid int, offset int) (int64, error, []News) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("news"))
	cond := orm.NewCondition()
	if condArr["classid"] != "" {
		cond = cond.And("classid", condArr["classid"])
	}
	//查询是否为第一页
	if lastid > 0 {
		cond = cond.And("id__lt", lastid)
	}
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	var News []News
	qs = qs.OrderBy("-id")
	num, err1 := qs.Limit(offset).All(&News)
	return num, err1, News
}

//获取消息类型列表
func ApiGetNewsInfo(id int) []ApiNews {
	var ApiNews []ApiNews
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("c.id,c.classid,c.title,c.content,c.creatime").From("pms_news AS c").Where("c.id=?")
	sql := qb.String()
	_, _ = o.Raw(sql, id).QueryRows(&ApiNews)
	return ApiNews
}

//获取消息类型列表
func ApiGetClassic() []ApiClassic {
	var Classic []ApiClassic
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("c.id,c.classname,c.parentid,c.sorts").From("pms_classic AS c")
	sql := qb.String()
	_, _ = o.Raw(sql).QueryRows(&Classic)
	return Classic
}

//获取任务单位
func (this *ApiArea) TableName() string {
	return models.TableName("area")
}

//获取社区单位
func ApiGetTags() []ApiTags {
	var ApiTags []ApiTags
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("t.id,t.name").From("pms_tags AS t").OrderBy("id")
	sql := qb.String()
	_, _ = o.Raw(sql).QueryRows(&ApiTags)
	return ApiTags
}

//社区任务文件表
type ApiTageFile struct {
	Id       int64
	Tid      int64
	Oldname  string
	Path     string
	Missonid int64
	Name     string
	Desc     string
	Creatime int64
}

//根据社区单位获取对应的文件任务
func TagsFile(tid int64) (error, []ApiTageFile) {
	var ApiTageFile []ApiTageFile
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("f.id,f.tid,f.oldname,f.path,f.missionid,m.name,m.desc,m.creatime").From("pms_files AS f").
		InnerJoin("pms_mission AS m").On("m.id = f.missionid").Where("f.tid=?")
	sql := qb.String()
	_, err := o.Raw(sql, tid).QueryRows(&ApiTageFile)
	return err, ApiTageFile
}

func (this *ApiMission) TableName() string {
	return models.TableName("mission_my")
}

//任务列表
func ApiGetMission(page int, offset int, aid int64) (error, []ApiMission) {

	var ApiMission []ApiMission
	o := orm.NewOrm()
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("my.id,my.missionid,my.types,my.userid,my.areaid,m.name,m.desc,m.creatime").From("pms_mission AS m").
		InnerJoin("pms_mission_my AS my").On("my.missionid = m.id").Where("my.areaid=?").OrderBy("m.id").Desc().Limit(offset).Offset(start)

	sql := qb.String()
	_, err := o.Raw(sql, aid).QueryRows(&ApiMission)
	return err, ApiMission
}

//任务信息
type ApiMissionInfo struct {
	Id       int64
	Name     string
	Types    int64
	Desc     string
	Userid   int64
	Creatime int64
}

//任务详情
func ApiGetMissionInfo(tid int64) (error, []ApiTageFile) {

	var ApiTageFile []ApiTageFile
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("f.id,f.tid,f.oldname,f.path,f.missionid,m.name,m.desc,m.creatime").From("pms_files AS f").
		InnerJoin("pms_mission AS m").On("m.id = f.missionid").Where("f.tid=?")
	sql := qb.String()
	_, err := o.Raw(sql, tid).QueryRows(&ApiTageFile)
	return err, ApiTageFile
}
