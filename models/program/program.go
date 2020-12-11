package program

import (
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Program struct {
	Id       int64
	Userid   int64
	Title    string
	Temps    string
	Publish  string
	Pstatus  int
	Counts   int64
	Creatime int64
	Updatime int64
}

func (this *Program) TableName() string {
	return models.TableName("program")
}
func init() {
	orm.RegisterModel(new(Program))
}

//项目列表
func ListProgram(condArr map[string]string, page int, offset int) (num int64, err error, user []Program) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("program"))
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("Title__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("pstatus", condArr["status"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	qs = qs.RelatedSel()

	var program []Program
	qs = qs.OrderBy("-id")
	num, err1 := qs.Limit(offset, start).All(&program)
	return num, err1, program
}
func GetAllpro() []Program {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("program"))
	var program []Program
	qs = qs.OrderBy("-id")
	_, err1 := qs.Limit(20).All(&program)

	if err1 != nil {
		return nil
	} else {
		return program
	}
}

//统计数量
func CountProgram(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("program"))
	qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("title__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("pstatus", condArr["status"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}
func AddProgram(t string) (int64, error) {
	o := orm.NewOrm()
	var p Program
	p.Title = t
	p.Creatime = time.Now().Unix()
	return o.Insert(p)
}
func GetProgram(id int64) (Program, error) {

	var p Program
	o := orm.NewOrm()
	p = Program{Id: id}
	err := o.Read(&p)
	return p, err
}
