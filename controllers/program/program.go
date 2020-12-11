package program

import (
	"fmt"
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/program"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

//列表
type ProgramController struct {
	controllers.BaseController
}

func (this *ProgramController) Get() {
	//权限检测
	/*if !strings.Contains(this.GetSession("userPermission").(string), "program-manage") {
		this.Redirect("/my/task", 302)
		return
		//this.Abort("401")
	}*/
	page, err := this.GetInt("p")
	status := this.GetString("status")
	keywords := this.GetString("keywords")
	if err != nil {
		page = 1
	}

	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	condArr := make(map[string]string)
	condArr["status"] = status
	condArr["keywords"] = keywords

	countProject := CountProgram(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, countProject)
	_, _, programs := ListProgram(condArr, page, offset)
	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["program"] = programs
	this.Data["countProject"] = countProject
	this.TplName = "program/program.tpl"
}

//添加项目
type ProgramAddController struct {
	controllers.BaseController
}

func (this *ProgramAddController) Get() {
	this.TplName = "program/program-form.tpl"
}
func (this *ProgramAddController) Post() {
	title := this.GetString("title")
	id, err := AddProgram(title)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "项目信息添加成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "项目信息添加失败"}
	}

}
func (this *ProgramAddController) Delete() {

}

type ProgramIndexController struct {
	controllers.BaseController
}

func (this *ProgramIndexController) Post() {
	pidstr := this.GetString("pid")
	pid, _ := strconv.ParseInt(pidstr, 10, 64)
	program, err := GetProgram(pid)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "测评不存在", "data": ""}
		this.ServeJSON()
		return
	}
	_, s := GetList(pid)
	//var m map[string]interface{}
	m := make(map[string]interface{})

	m["project"] = program
	m["subject"] = s
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "请填写项目名称", "data": m}
	this.ServeJSON()
}

type SubjectManageController struct {
	controllers.BaseController
}

func (this *SubjectManageController) Get() {
	id, _ := this.GetInt64("pid")
	this.Data["pid"] = id
	this.TplName = "program/ifram.tpl"
}