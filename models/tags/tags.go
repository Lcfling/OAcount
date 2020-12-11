package tags

import (
	"github.com/astaxie/beego/orm"
	"github.com/virteman/OAcount/models"
)

type Tags struct {
	Id int64
	Name string
}

func (this *Tags) TableName() string {
	return models.TableName("tags")
}
func init() {
	orm.RegisterModel(new(Tags))
}
func GetTagsAll() []Tags {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("tags"))
	var tags []Tags
	qs.All(&tags)
	return tags
}
func Isexsit(name string) (int64,bool){

	return 1,false
}
func AddTags(name string) (id int64,err error) {
	o := orm.NewOrm()
	if id,ok:=Isexsit(name);ok{
		return id,nil
	}else{
		tag:=new(Tags)
		tag.Name=name
		return o.Insert(tag)
	}
}
func GetTagNameById(id int64) (string,error) {
	o := orm.NewOrm()
	t:= Tags{Id: id}
	err := o.Read(&t)
	return t.Name,err
}
func GetTagIdByName(name string) (int64,error) {
	o := orm.NewOrm()
	t:= Tags{Name: name}
	err := o.Read(&t)
	return t.Id,err
}