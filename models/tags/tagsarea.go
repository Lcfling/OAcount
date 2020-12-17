package tags

import (
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Tagsarea struct {
	Id  int64
	Aid int64
	Tid int64
}

func (this *Tagsarea) TableName() string {
	return models.TableName("tags_area")
}
func init() {
	orm.RegisterModel(new(Tagsarea))
}

func AddTagsarea(aid int64, tid int64) (int64, error) {
	//err:=nil
	if aid > 0 && tid > 0 {
		o := orm.NewOrm()
		t := new(Tagsarea)
		t.Aid = aid
		t.Tid = tid
		return o.Insert(t)

	} else {
		return 0, nil
	}

}

func GetAreaBytagid(id int64) []Tagsarea {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("tags_area"))
	cond := orm.NewCondition()
	cond = cond.And("tid", id)
	var tagsarea []Tagsarea
	qs.All(&tagsarea)
	return tagsarea
}
func GetAreaBytagName(name string) []Tagsarea {
	id, err := GetTagIdByName(name)
	if !(err == nil && id > 0) {
		return nil
	}
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("tags_area"))
	cond := orm.NewCondition()
	cond = cond.And("tid", id)
	var tagsarea []Tagsarea
	qs.All(&tagsarea)
	return tagsarea
}

//返回area ID数组
func GetUidsByTid(tid int64) []string {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("tags_area"))
	cond := orm.NewCondition()
	cond = cond.And("tid", tid)
	qs = qs.SetCond(cond)
	var tagsarea []Tagsarea
	num, _ := qs.All(&tagsarea)
	if num > 0 {
		var areaLists []string
		for _, v := range tagsarea {
			var aid string
			aid = strconv.Itoa(int(v.Aid))
			areaLists = append(areaLists, aid)
		}
		return areaLists
	} else {
		return nil
	}

}
