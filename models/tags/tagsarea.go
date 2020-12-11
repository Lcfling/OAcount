package tags

import (
	"github.com/astaxie/beego/orm"
	"github.com/virteman/OAcount/models"
)

type Tagsarea struct {
	Id int64
	Aid int64
	Tid int64
}
func (this *Tagsarea) TableName() string {
	return models.TableName("tags_area")
}
func init() {
	orm.RegisterModel(new(Tagsarea))
}

func AddTagsarea(aid int64,tid int64) (int64,error)  {
	//err:=nil
	if aid>0 && tid>0{
		o := orm.NewOrm()
		t:=new(Tagsarea)
		t.Aid=aid
		t.Tid=tid
		return o.Insert(t)

	}else {
		return 0,nil
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
	id,err:=GetTagIdByName(name)
	if !(err==nil&&id>0){
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