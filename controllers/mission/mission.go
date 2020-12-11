package mission

import (
	"fmt"
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/files"
	. "github.com/Lcfling/OAcount/models/mission"
	. "github.com/Lcfling/OAcount/models/program"
	"github.com/Lcfling/OAcount/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type MissionManageController struct {
	controllers.BaseController
}

func (this *MissionManageController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "mission-manage") {
		this.Redirect("/", 302)
		return
		//this.Abort("401")
	}
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

	countMission := CountMission(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, countMission)
	_, _, missions := ListMission(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["missions"] = missions
	this.Data["countMission"] = countMission

	this.TplName = "mission/mission.tpl"
}
func (this *MissionManageController) Post() {

}

type MissionAddController struct {
	controllers.BaseController
}

func (this *MissionAddController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "mission-add") {
		this.Abort("401")
	}
	var mission Mission
	mission.Started = time.Now().Unix()
	mission.Ended = time.Now().Unix()
	mission.Types = 0
	mission.Mid = 0
	this.Data["program"] = GetAllpro()
	this.Data["mission"] = mission
	this.TplName = "mission/mission-form.tpl"
}
func (this *MissionAddController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "mission-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	name := this.GetString("name")
	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写任务名称"}
		this.ServeJSON()
		return
	}
	startedstr := this.GetString("started")
	if "" == startedstr {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写开始时间"}
		this.ServeJSON()
		return
	}
	startedtime := utils.GetDateParse(startedstr)
	types, _ := this.GetInt("types")
	var aid int64
	aid = 0
	if types == 2 {
		aid, _ = this.GetInt64("aid")
		if !(aid > 0) {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择相关测评"}
			this.ServeJSON()
		}
	}

	endedstr := this.GetString("ended")
	if "" == endedstr {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写结束时间"}
		this.ServeJSON()
		return
	}
	endedtime := utils.GetDateParse(endedstr)

	desc := this.GetString("desc")
	if "" == desc {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写项目描述"}
		this.ServeJSON()
		return
	}

	userid := this.BaseController.UserUserId

	var err error
	//雪花算法ID生成
	//id := utils.SnowFlakeId()

	var pro Mission
	pro.Userid = userid
	pro.Name = name
	pro.Mid = aid //此处测评id
	pro.Types = types
	pro.Started = startedtime
	pro.Ended = endedtime
	pro.Desc = desc

	id, err := AddMission(pro)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "项目信息添加成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "项目信息添加失败"}
	}
	this.ServeJSON()
}

type MissionDetailController struct {
	controllers.BaseController
}

func (this *MissionDetailController) Get() {
	id := this.Ctx.Input.Param(":id")
	var mission Mission
	id64, _ := strconv.ParseInt(id, 10, 64)
	mission, err := GetMission(id64)
	if err != nil {

	}
	this.Data["mission"] = mission
	this.TplName = "mission/mission-detail.tpl"
}

type MissionMyController struct {
	controllers.BaseController
}

func (this *MissionMyController) Get() {
	userid := this.BaseController.UserUserId
	_, _, missionmy := GetMyMission(userid, 1, 100)
	this.Data["missionmy"] = missionmy
	this.Data["countmissionmy"] = len(missionmy)

	this.TplName = "mission/missionmy.tpl"
}

type MissionMySubController struct {
	controllers.BaseController
}

func (this *MissionMySubController) Get() {

	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	missionmy := GetMissionMy(id)
	fmt.Println(missionmy)
	this.Data["missionmy"] = missionmy
	this.TplName = "mission/missionmy-form.tpl"
}
func (this *MissionMySubController) Post() {
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	//headers,err:=this.GetFiles("attachment")
	headers, err := this.Ctx.Request.MultipartForm.File["attachment"]
	//fmt
	feedback := this.GetString("feedback")
	detail := this.GetString("detail")
	if len(headers) > 0 {
		var filepath string
		for _, h := range headers {
			now := time.Now()
			dir := "./static/uploadfile/" + strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
			err1 := os.MkdirAll(dir, 0755)
			if err1 != nil {
				this.Data["json"] = map[string]interface{}{"code": 1, "message": "目录权限不够"}
				this.ServeJSON()
				return
			}
			//生成新的文件名
			filename := h.Filename
			fileExt := path.Ext(filepath)
			var fileType int
			switch fileExt {
			case "jpg":
				fileType = 0
			case "png":
				fileType = 0
			case "gif":
				fileType = 0
			case "doc":
				fileType = 1
			case "docx":
				fileType = 1
			default:
				fileType = -1
			}
			//ext := utils.SubString(filename, strings.LastIndex(filename, "."), 5)
			//filename = utils.GetGuid() + ext

			time64 := time.Now().Unix()
			timestr := strconv.FormatInt(time64, 10) + utils.RandChar(6)
			newName := timestr + filename
			if err != true {
				this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
				this.ServeJSON()
				return
			} else {
				//this.SaveToFile("imgFile", "./static/uploadfile/"+h.Filename)
				this.SaveToFile("attachment", dir+"/"+newName)
				filepath = strings.Replace(dir, ".", "", 1) + "/" + newName
				Addfile(0, id, fileType, 0, filename, filepath)
			}
		}
	}
	var m MissionMy
	m.Detail = detail
	m.Feedback = feedback
	err1 := UpdateMissionMy(m, id)
	if err1 != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "success"}
		this.ServeJSON()
	}
}

//任务完善结构体
type MissionPerfectController struct {
	controllers.BaseController
}

//任务完善页面 Get
func (this *MissionPerfectController) Get() {

	id := this.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	mission, err := GetMission(id64)
	if err != nil {

	}
	this.Data["program"] = GetAllpro()
	this.Data["mission"] = mission
	this.TplName = "mission/mission-perfect.tpl"
	fmt.Println("Mission:", mission)
}

//任务完善上传文件
func (this *MissionPerfectController) Post() {
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	//headers,err:=this.GetFiles("attachment")
	headers, err := this.Ctx.Request.MultipartForm.File["attachment"]
	//fmt
	if len(headers) > 0 {
		var filepath string
		for _, h := range headers {
			now := time.Now()
			dir := "./static/uploadfile/" + strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
			err1 := os.MkdirAll(dir, 0755)
			if err1 != nil {
				this.Data["json"] = map[string]interface{}{"code": 1, "message": "目录权限不够"}
				this.ServeJSON()
				return
			}
			//生成新的文件名
			filename := h.Filename
			fileExt := path.Ext(filepath)
			var fileType int
			switch fileExt {
			case "jpg":
				fileType = 0
			case "png":
				fileType = 0
			case "gif":
				fileType = 0
			case "doc":
				fileType = 1
			case "docx":
				fileType = 1
			default:
				fileType = -1
			}
			//ext := utils.SubString(filename, strings.LastIndex(filename, "."), 5)
			//filename = utils.GetGuid() + ext

			if err != true {
				this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
				this.ServeJSON()
				return
			} else {
				//this.SaveToFile("imgFile", "./static/uploadfile/"+h.Filename)
				this.SaveToFile("attachment", dir+"/"+filename)
				filepath = strings.Replace(dir, ".", "", 1) + "/" + filename
				fmt.Println("id:", id)
				fmt.Println("fileType:", fileType)
				//Addfile(0,id,fileType,0,filepath)
			}
		}
	}
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "success"}
	this.ServeJSON()

}
