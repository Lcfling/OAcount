package mission

import (
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Examination struct {
	Id       int64
	Pid      int64
	Name     string
	Content  string
	Need     string
	Creatime int64
}

func (this *Examination) TableName() string {
	return models.TableName("examination")
}
func init() {
	orm.RegisterModel(new(Examination))
}
func AddExamination(data Examination) (id int64, err error) {
	o := orm.NewOrm()
	pro := new(Examination)
	//pro.Id = updPro.Id
	pro.Pid = data.Pid
	pro.Name = data.Name
	pro.Content = data.Content
	pro.Need = data.Need
	pro.Creatime = time.Now().Unix()
	id, err = o.Insert(pro)
	return id, err
}

func CountExamination(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("examination"))
	cond := orm.NewCondition()
	if condArr["parentid"] != "" {
		cond = cond.And("pid", condArr["parentid"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}
func ListExamination(condArr map[string]string, page int, offset int) (num int64, err error, ops []Examination) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("examination"))
	cond := orm.NewCondition()

	if condArr["parentid"] != "" {
		cond = cond.And("pid", condArr["parentid"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	qs = qs.OrderBy("id")
	var areas []Examination
	num, errs := qs.Limit(offset, start).All(&areas)
	return num, errs, areas
}
