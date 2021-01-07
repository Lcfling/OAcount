package mission

import (
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Document struct {
	Id       int64
	Pid      int64
	Name     string
	Content  string
	Need     string
	Creatime int64
}

func (this *Document) TableName() string {
	return models.TableName("document")
}
func init() {
	orm.RegisterModel(new(Document))
}
func AddDocument(data Document) (id int64, err error) {
	o := orm.NewOrm()
	pro := new(Document)
	//pro.Id = updPro.Id
	pro.Pid = data.Pid
	pro.Name = data.Name
	pro.Content = data.Content
	pro.Need = data.Need
	pro.Creatime = time.Now().Unix()
	id, err = o.Insert(pro)
	return id, err
}

func CountDocument(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("document"))
	cond := orm.NewCondition()
	if condArr["parentid"] != "" {
		cond = cond.And("pid", condArr["parentid"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}
func ListDocument(condArr map[string]string, page int, offset int) (num int64, err error, ops []Document) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("document"))
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
	var areas []Document
	num, errs := qs.Limit(offset, start).All(&areas)
	return num, errs, areas
}
func GetDocument(id int64) (Document, error) {

	var doc Document
	var err error

	//err = utils.GetCache("GetProject.id."+fmt.Sprintf("%d", id), &project)
	//if err != nil {
	o := orm.NewOrm()
	doc = Document{Id: id}
	err = o.Read(&doc)
	//utils.SetCache("GetProject.id."+fmt.Sprintf("%d", id), project, 600)
	//}
	return doc, err
}
