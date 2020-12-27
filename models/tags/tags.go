package tags

import (
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego/orm"
)

type Tags struct {
	Id   int64
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
func Isexsit(name string) (int64, bool) {

	o := orm.NewOrm()
	tag := Tags{Name: name}
	o.Read(&tag, "name")
	if tag.Id > 0 {
		return tag.Id, true
	} else {
		return 0, false
	}

}
func AddTags(name string) (id int64, err error) {
	o := orm.NewOrm()
	if id, ok := Isexsit(name); ok {
		return id, nil
	} else {
		tag := new(Tags)
		tag.Name = name
		return o.Insert(tag)
	}
}
func GetTagNameById(id int64) (string, error) {
	o := orm.NewOrm()
	t := Tags{Id: id}
	err := o.Read(&t)
	return t.Name, err
}
func GetTagIdByName(name string) (int64, error) {
	o := orm.NewOrm()
	t := Tags{Name: name}
	err := o.Read(&t)
	return t.Id, err
}
