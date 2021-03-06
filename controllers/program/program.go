package program

import (
	"encoding/json"
	"fmt"
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/area"
	. "github.com/Lcfling/OAcount/models/program"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/bitly/go-simplejson"
	"strconv"
	"time"
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

	this.ServeJSON()
}

//添加项目
type ProgramEditController struct {
	controllers.BaseController
}

func (this *ProgramEditController) Get() {

	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	if !(id > 0) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "测评不存在"}
		this.ServeJSON()
	}
	p, err := GetProgram(id)
	if err != nil {
		this.Ctx.Redirect(404, "/program/manage")
	}
	this.Data["program"] = p
	this.TplName = "program/program-form.tpl"
}
func (this *ProgramEditController) Post() {
	title := this.GetString("title")
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	_, err := UpdateProgram(id, title)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "测评更新成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "测评更新失败"}
	}

	this.ServeJSON()
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

//列表
type ListController struct {
	controllers.IndexController
}

func (this *ListController) Get() {
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
func (this *ListController) Post() {
	lastid, err := this.GetInt("lastid")
	if err != nil {
		lastid = 0
	}

	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	_, _, programs := ApiGetProgram(lastid, offset)

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "消息列表", "data": programs}
	this.ServeJSON()
}

//列表
type ShareController struct {
	controllers.IndexController
}

func (this *ShareController) Get() {
	id, _ := this.GetInt64("id")
	_, err := GetProgram(id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "获取内容失败"}
	} else {
		url := "http://192.168.0.116:8088/static/html/share.html?pid=" + strconv.FormatInt(id, 10)
		//data:=map[string]string{}
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "success", "data": url}
	}
	this.ServeJSON()
}

//测评详情
/*type DetailController struct {
	controllers.IndexController
}
func (this *DetailController) Post() {
	lastid, _ := this.GetInt("id")


	//_, _, programs := ApiGetProgram(lastid, offset)

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "消息列表", "data": programs}
	this.ServeJSON()
}*/
type AnwserList struct {
	controllers.MobileController
}

func (this *AnwserList) Post() {
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
	m["area"] = GetChild(0, "0")
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "请填写项目名称", "data": m}
	this.ServeJSON()
}

type AnswerController struct {
	controllers.MobileController
}

func (this *AnswerController) Post() {

	pid, err := this.GetInt64("pid")
	if err != nil || !(pid > 0) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "获取测评失败"}
	}
	areaid, _ := this.GetInt64("area")
	office, _ := this.GetInt64("office")
	community, _ := this.GetInt64("community")
	addr := this.GetString("addr")
	name := this.GetString("name")
	mobile := this.GetString("mobile")
	if addr == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请完善地址信息"}
		this.ServeJSON()
		return
	}
	if name == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请完善姓名信息"}
		this.ServeJSON()
		return
	}
	if mobile == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请完善手机信息"}
		this.ServeJSON()
		return
	}
	Content := this.GetString("data")
	//var ContentJson map[string]interface{}

	fmt.Println("string=======", Content)
	//err = json.Unmarshal([]byte(Content), &ContentJson)

	ContentJson, err := simplejson.NewJson([]byte(Content))
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "答题数据解析失败+" + err.Error()}
		this.ServeJSON()
		return
	}

	arr, _ := ContentJson.Array()
	var count int64
	var sumscore int64

	for _, v := range arr {
		count += 1

		//var Subject []interface{}
		c := v.(map[string]interface{})
		s := c["Content"]
		Subject := s.(map[string]interface{})
		arrS := Subject["Subject"]
		arrSs := arrS.([]interface{})
		for k, j := range arrSs {
			state, _ := j.(map[string]interface{})["Check"].(json.Number).Int64()
			if state == 1 {
				sumscore += (4 - int64(k))
			}
		}
	}
	fmt.Println("count", count)
	fmt.Println("sumscore", sumscore)
	fmt.Println("最终得分:", 100*sumscore/(count*4))

	pro := new(Sub)
	pro.Area = areaid
	pro.Office = office
	pro.Community = community

	pro.Score = 100 * sumscore / (count * 4)
	if pro.Score >= 80 {
		pro.Status = 1
	} else {
		pro.Status = 0
	}
	pro.Addr = addr
	pro.Name = name
	pro.Mobile = mobile
	pro.Pid = pid
	pro.Creatime = time.Now().Unix()
	pro.Content = Content
	_, err = AddSub(pro)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "入库失败：" + err.Error()}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "提交成功"}
	}
	this.ServeJSON()
}
