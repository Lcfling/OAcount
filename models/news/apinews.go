package news

import (
	"github.com/Lcfling/OAcount/models"
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
func ApiGetClassic() []ApiClassic {
	var Classic []ApiClassic
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("c.id,c.classname,c.parentid,c.sorts").From("pms_classic AS c")
	sql := qb.String()
	_, _ = o.Raw(sql).QueryRows(&Classic)
	return Classic
}
