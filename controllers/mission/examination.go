package mission

import (
	"fmt"
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/mission"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
	"time"
)

type ExaClassManageController struct {
	controllers.BaseController
}

func (this *ExaClassManageController) Get() {
	parentid, _ := this.GetInt64("pid")

	page, err := this.GetInt("p")
	if err != nil {
		page = 1
	}
	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	condArr := make(map[string]string)

	condArr["parentid"] = strconv.FormatInt(parentid, 10)
	CountClass := CountExaclass(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, CountClass)
	_, _, class := ListExaClass(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["class"] = class
	this.Data["CountClass"] = CountClass

	this.Data["parentid"] = parentid
	this.TplName = "mission/exaclass.tpl"
}

type ExaClassAddController struct {
	controllers.BaseController
}

func (this *ExaClassAddController) Get() {
	pid, _ := this.GetInt64("pid")
	this.Data["pid"] = pid
	this.TplName = "mission/exaclass-form.tpl"
}
func (this *ExaClassAddController) Post() {

	pid, _ := this.GetInt64("pid")
	title := this.GetString("title")

	var pro Exaclass
	pro.Pid = pid
	pro.Title = title
	pro.Creatime = time.Now().Unix()
	id, err := AddExaClass(pro)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败"}

	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "添加成功", "id": fmt.Sprintf("%d", id)}
	}
	this.ServeJSON()

}

type ExaminationManageController struct {
	controllers.BaseController
}

func (this *ExaminationManageController) Get() {
	pid, _ := this.GetInt64("pid")

	page, err := this.GetInt("p")
	if err != nil {
		page = 1
	}
	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	condArr := make(map[string]string)

	condArr["parentid"] = strconv.FormatInt(pid, 10)
	CountExa := CountExamination(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, CountExa)
	_, _, exa := ListExamination(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["exa"] = exa
	this.Data["CountExa"] = CountExa

	this.Data["parentid"] = pid
	this.TplName = "mission/examination.tpl"
}

type ExaminationAddController struct {
	controllers.BaseController
}

func (this *ExaminationAddController) Get() {
	pid, _ := this.GetInt64("pid")
	this.Data["pid"] = pid

	this.TplName = "mission/examination-form.tpl"
}
func (this *ExaminationAddController) Post() {

	pid, _ := this.GetInt64("pid")
	name := this.GetString("name")
	content := this.GetString("content")
	need := this.GetString("need")

	var pro Examination
	pro.Pid = pid
	pro.Name = name
	pro.Content = content
	pro.Need = need
	pro.Creatime = time.Now().Unix()
	id, err := AddExamination(pro)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "添加成功", "id": fmt.Sprintf("%d", id)}
	}
	this.ServeJSON()

}

type ExaclassTree struct {
	controllers.IndexController
}

func (this *ExaclassTree) Get() {
	pid, _ := this.GetInt64("pid")
	tree := GetExaChildTree(pid, "")
	if len(tree) > 0 {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "success", "data": tree}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "feild"}
	}
	this.ServeJSON()
}

type ExaminationList struct {
	controllers.IndexController
}

func (this *ExaminationList) Get() {
	pid, _ := this.GetInt64("pid")

	condArr := make(map[string]string)

	condArr["parentid"] = strconv.FormatInt(pid, 10)
	_, _, exa := ListExamination(condArr, 1, 30)

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "success", "data": exa}
	this.ServeJSON()
}
