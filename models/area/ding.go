package area

import (
	"github.com/astaxie/beego/orm"
	"github.com/virteman/OAcount/models"
	"time"
)

type Ding struct {
	Id int64
	Creatime int64
}

func (this *Ding) TableName() string {
	return models.TableName("ding")
}
func init()  {
	orm.RegisterModel(new(Ding))
}

func InsertNew() error {

	o := orm.NewOrm()
	areaold := new(Ding)
	areaold.Creatime=time.Now().Unix()
	_,err:=o.Insert(areaold)
	return err
}
func GetNew() (d Ding,err error){
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id","creatime").From("pms_ding").
		OrderBy("-id").
		Limit(1)

	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&d)

	return d,nil

}
