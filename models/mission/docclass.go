package mission

import (
	"fmt"
	"github.com/Lcfling/OAcount/models"
	"github.com/Lcfling/OAcount/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Docclass struct {
	Id       int64
	Pid      int64
	Title    string
	Creatime int64
	Show     int64      `orm:"-"`
	Child    []Docclass `orm:"-"`
}

func (this *Docclass) TableName() string {
	return models.TableName("docclass")
}
func init() {
	orm.RegisterModel(new(Docclass))
}

func AddClass(data Docclass) (id int64, err error) {
	o := orm.NewOrm()
	pro := new(Docclass)

	//pro.Id = updPro.Id
	pro.Pid = data.Pid
	pro.Title = data.Title
	pro.Creatime = time.Now().Unix()
	id, err = o.Insert(pro)
	return id, err
}

func CountDocclass(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("docclass"))
	cond := orm.NewCondition()
	if condArr["parentid"] != "" {
		cond = cond.And("pid", condArr["parentid"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}
func ListClass(condArr map[string]string, page int, offset int) (num int64, err error, ops []Docclass) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("docclass"))
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
	var areas []Docclass
	num, errs := qs.Limit(offset, start).All(&areas)
	return num, errs, areas
}
func GetDocclassName(id int64) string {
	var err error
	var title string

	err = utils.GetCache("GetDocclassName.id."+fmt.Sprintf("%d", id), &title)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var class Docclass
		o := orm.NewOrm()
		o.QueryTable(models.TableName("docclass")).Filter("id", id).One(&class, "title")
		title = class.Title
		if "" == title {
			title = "一级分类"
		}
		utils.SetCache("GetDocclassName.id."+fmt.Sprintf("%d", id), title, cache_expire)
	}
	return title
}
func GetChildTree(pid int64, types string) []Docclass {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("docclass"))
	cond := orm.NewCondition()
	cond = cond.And("pid", pid)
	qs = qs.SetCond(cond)
	var class []Docclass
	num, _ := qs.All(&class)
	if num > 0 {
		var docClass []Docclass
		for _, v := range class {
			var classic Docclass
			classic.Title = v.Title
			classic.Pid = v.Pid
			classic.Id = v.Id
			classic.Creatime = v.Creatime

			Child := GetDocumentsToClassByPid(v.Id)
			if Child != nil {
				classic.Child = Child
				classic.Show = 0
			} else {
				classic.Child = GetChildTree(v.Id, types)
				classic.Show = 0
			}

			docClass = append(docClass, classic)
		}
		return docClass
	} else {
		return nil
	}

}

func GetDocumentsToClassByPid(pid int64) []Docclass {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("document"))
	cond := orm.NewCondition()
	cond = cond.And("pid", pid)
	qs = qs.SetCond(cond)
	var doc []Document
	num, _ := qs.All(&doc)
	if num > 0 {
		var docClass []Docclass
		for _, v := range doc {
			var classic Docclass
			classic.Title = v.Name
			classic.Pid = v.Pid
			classic.Id = v.Id
			classic.Creatime = v.Creatime
			classic.Show = 1
			docClass = append(docClass, classic)
		}
		return docClass
	} else {
		return nil
	}
}

func GetMyChildTree(pid int64, userid int64, types int64) []Docclass {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("docclass"))
	cond := orm.NewCondition()
	cond = cond.And("pid", pid)
	qs = qs.SetCond(cond)
	var class []Docclass
	num, _ := qs.All(&class)
	if num > 0 {
		var docClass []Docclass
		for _, v := range class {
			var classic Docclass
			classic.Title = v.Title
			classic.Pid = v.Pid
			classic.Id = v.Id
			classic.Creatime = v.Creatime

			Child := GetMyDocumentsToClassByPid(v.Id, userid, types)
			if Child != nil {
				classic.Child = Child
				classic.Show = 0
			} else {
				classic.Child = GetMyChildTree(v.Id, userid, types)
				classic.Show = 0
			}

			docClass = append(docClass, classic)
		}
		return docClass
	} else {
		return nil
	}

}
func GetMyDocumentsToClassByPid(pid int64, userid int64, types int64) []Docclass {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("document"))
	cond := orm.NewCondition()
	cond = cond.And("pid", pid)
	qs = qs.SetCond(cond)
	var doc []Document
	num, _ := qs.All(&doc)
	if num > 0 {
		var docClass []Docclass
		for _, v := range doc {
			var classic Docclass

			missionmy := GetMymissionByType(userid, v.Id, types)

			if missionmy.Id > 0 {
				classic.Title = v.Name
				classic.Pid = v.Id
				classic.Id = missionmy.Id
				classic.Creatime = v.Creatime
				classic.Show = 1

				docClass = append(docClass, classic)
			}
		}
		return docClass
	} else {
		return nil
	}
}

func TreeFlater(tree []Docclass) []Docclass {
	var newTree []Docclass

	for _, v := range tree {
		var docclass Docclass

		if v.Child != nil {
			Child := TreeFlater(v.Child)
			if Child != nil {
				docclass.Child = Child
				docclass.Id = v.Id
				docclass.Show = v.Show
				docclass.Pid = v.Pid
				docclass.Creatime = v.Creatime
				docclass.Title = v.Title
				newTree = append(newTree, docclass)
			}

		} else if v.Show == 1 && v.Child == nil {
			docclass.Child = nil
			docclass.Id = v.Id
			docclass.Show = v.Show
			docclass.Pid = v.Pid
			docclass.Creatime = v.Creatime
			docclass.Title = v.Title
			newTree = append(newTree, docclass)
		}
	}
	return newTree
}
