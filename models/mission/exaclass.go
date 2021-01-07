package mission

import (
	"fmt"
	"github.com/Lcfling/OAcount/models"
	"github.com/Lcfling/OAcount/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Exaclass struct {
	Id       int64
	Pid      int64
	Title    string
	Creatime int64
	Show     int64      `orm:"-"`
	Child    []Exaclass `orm:"-"`
}

func (this *Exaclass) TableName() string {
	return models.TableName("exaclass")
}
func init() {
	orm.RegisterModel(new(Exaclass))
}

func AddExaClass(data Exaclass) (id int64, err error) {
	o := orm.NewOrm()
	pro := new(Exaclass)

	//pro.Id = updPro.Id
	pro.Pid = data.Pid
	pro.Title = data.Title
	pro.Creatime = time.Now().Unix()
	id, err = o.Insert(pro)
	return id, err
}

func CountExaclass(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("exaclass"))
	cond := orm.NewCondition()
	if condArr["parentid"] != "" {
		cond = cond.And("pid", condArr["parentid"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}
func ListExaClass(condArr map[string]string, page int, offset int) (num int64, err error, ops []Exaclass) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("exaclass"))
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
	var areas []Exaclass
	num, errs := qs.Limit(offset, start).All(&areas)
	return num, errs, areas
}
func GetExaclassName(id int64) string {
	var err error
	var title string

	err = utils.GetCache("GetExaclassName.id."+fmt.Sprintf("%d", id), &title)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var class Exaclass
		o := orm.NewOrm()
		o.QueryTable(models.TableName("exaclass")).Filter("id", id).One(&class, "title")
		title = class.Title
		if "" == title {
			title = "一级分类"
		}
		utils.SetCache("GetExaclassName.id."+fmt.Sprintf("%d", id), title, cache_expire)
	}
	return title
}

func GetExaChildTree(pid int64, types string) []Exaclass {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("exaclass"))
	cond := orm.NewCondition()
	cond = cond.And("pid", pid)
	qs = qs.SetCond(cond)
	var class []Exaclass
	num, _ := qs.All(&class)
	if num > 0 {
		var exaClass []Exaclass
		for _, v := range class {
			var classic Exaclass
			classic.Title = v.Title
			classic.Pid = v.Pid
			classic.Id = v.Id
			classic.Creatime = v.Creatime

			nums := GetExaNums(v.Id)
			if nums > 0 {

				classic.Show = 1
			} else {

				classic.Show = 0
			}
			classic.Child = GetExaChildTree(v.Id, types)

			exaClass = append(exaClass, classic)
		}
		return exaClass
	} else {
		return nil
	}

}
func GetExaNums(pid int64) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("examination"))
	cond := orm.NewCondition()
	cond = cond.And("pid", pid)
	qs = qs.SetCond(cond)
	var exa []Examination
	num, _ := qs.All(&exa)
	return num
}
