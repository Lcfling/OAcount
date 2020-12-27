package program

import (
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego/orm"
)

type Sub struct {
	Id        int64
	Pid       int64
	Area      int64
	Office    int64
	Community int64
	Addr      string
	Name      string
	Mobile    string
	Content   string
	Creatime  int64
	Status    int
	Score     int64
}

func (this *Sub) TableName() string {
	return models.TableName("sub")
}
func init() {
	orm.RegisterModel(new(Sub))
}

func AddSub(pro *Sub) (int64, error) {
	p := new(Sub)
	p = pro
	o := orm.NewOrm()
	return o.Insert(p)
}
