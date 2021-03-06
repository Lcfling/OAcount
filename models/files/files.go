package files

import (
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Files struct {
	Id          int64
	Missionmyid int64
	Aid         int64
	Types       int
	Oldname     string
	Path        string
	Missionid   int64
	Creatime    int64
	Status      int64
}

func (this *Files) TableName() string {
	return models.TableName("files")
}
func init() {
	orm.RegisterModel(new(Files))
}

func Addfile(Aid int64, missionmyid int64, types int, missionid int64, filename string, path string) (id int64, err error) {

	o := orm.NewOrm()
	pro := new(Files)
	//pro.Id = updPro.Id
	pro.Missionmyid = missionmyid
	pro.Aid = Aid
	pro.Types = types
	pro.Oldname = filename
	pro.Missionid = missionid
	pro.Path = path
	pro.Creatime = time.Now().Unix()
	return o.Insert(pro)
}

//统计数量
func CountFiles(condArr map[string]interface{}) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("files"))
	qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("oldname__icontains", condArr["keywords"]))
	}
	if condArr["types"].(int) != 0 {
		cond = cond.And("types", condArr["types"])
	}
	if condArr["missionmyid"].(int64) != 0 {
		cond = cond.And("missionmyid", condArr["missionmyid"])
	}
	if condArr["aid"].(int64) != 0 {
		cond = cond.And("aid", condArr["aid"])
	}
	if condArr["missionid"].(int64) != 0 {
		cond = cond.And("missionid", condArr["missionid"])
	}
	cond = cond.And("status", 0)
	num, _ := qs.SetCond(cond).Count()
	return num
}

//项目列表
func ListFiles(condArr map[string]interface{}, page int, offset int) (num int64, err error, user []Files) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("files"))
	cond := orm.NewCondition()
	if condArr["keywords"].(string) != "" {
		cond = cond.AndCond(cond.And("oldname__icontains", condArr["keywords"]))
	}
	if condArr["types"].(int) != 0 {
		cond = cond.And("types", condArr["types"])
	}
	if condArr["missionmyid"].(int64) != 0 {
		cond = cond.And("missionmyid", condArr["missionmyid"])
	}
	if condArr["aid"].(int64) != 0 {
		cond = cond.And("aid", condArr["aid"])
	}
	if condArr["missionid"].(int64) != 0 {
		cond = cond.And("missionid", condArr["missionid"])
	}
	cond = cond.And("status", 0)
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	qs = qs.RelatedSel()

	var files []Files
	qs = qs.OrderBy("-id")
	num, err1 := qs.Limit(offset, start).All(&files)
	return num, err1, files
}
func DeleteFile(id int64) error {
	var file Files
	o := orm.NewOrm()
	file = Files{Id: id}
	err := o.Read(&file)
	if err != nil {
		return err
	}
	file.Status = 1
	_, err = o.Update(&file)
	return err
}

//根据任务获取对应的文件
func MissionGetFile(condArr map[string]interface{}) (num int64, err error, user []Files) {

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("files"))
	cond := orm.NewCondition()
	if condArr["missionmyid"].(int64) != 0 {
		cond = cond.And("missionmyid", condArr["missionmyid"])
	}
	if condArr["aid"].(int64) != 0 {
		cond = cond.And("aid", condArr["aid"])
	}
	cond = cond.And("status", 0)
	qs = qs.SetCond(cond)

	qs = qs.RelatedSel()
	var files []Files
	qs = qs.OrderBy("-id")
	num, err1 := qs.Limit(1).All(&files)
	return num, err1, files
}

//文件进行修改关联
func TageFile(tid, missionid int64) error {
	o := orm.NewOrm()
	_, err := o.Raw("update  "+models.TableName("files")+" SET tid=?  WHERE missionid = ?", tid, missionid).Exec()
	return err

}
