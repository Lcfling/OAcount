package mission

import (
	"fmt"
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/files"
	. "github.com/Lcfling/OAcount/models/mission"
	"github.com/Lcfling/OAcount/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type DocClassManageController struct {
	controllers.BaseController
}

func (this *DocClassManageController) Get() {
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
	CountClass := CountDocclass(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, CountClass)
	_, _, class := ListClass(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["class"] = class
	this.Data["CountClass"] = CountClass

	this.Data["parentid"] = parentid
	this.TplName = "mission/docclass.tpl"
}

type DocClassAddController struct {
	controllers.BaseController
}

func (this *DocClassAddController) Get() {
	pid, _ := this.GetInt64("pid")
	this.Data["pid"] = pid
	this.TplName = "mission/docclass-form.tpl"
}
func (this *DocClassAddController) Post() {

	pid, _ := this.GetInt64("pid")
	title := this.GetString("title")

	var pro Docclass
	pro.Pid = pid
	pro.Title = title
	pro.Creatime = time.Now().Unix()
	id, err := AddClass(pro)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败"}

	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "添加成功", "id": fmt.Sprintf("%d", id)}
	}
	this.ServeJSON()

}

type DocumentManageController struct {
	controllers.BaseController
}

func (this *DocumentManageController) Get() {
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
	CountDoc := CountDocument(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, CountDoc)
	_, _, doc := ListDocument(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["doc"] = doc
	this.Data["countdoc"] = CountDoc

	this.Data["parentid"] = pid
	this.TplName = "mission/document.tpl"
}

type DocumentAddController struct {
	controllers.BaseController
}

func (this *DocumentAddController) Get() {
	pid, _ := this.GetInt64("pid")
	this.Data["pid"] = pid

	this.TplName = "mission/document-form.tpl"
}
func (this *DocumentAddController) Post() {

	pid, _ := this.GetInt64("pid")
	name := this.GetString("name")
	content := this.GetString("content")
	need := this.GetString("need")

	var pro Document
	pro.Pid = pid
	pro.Name = name
	pro.Content = content
	pro.Need = need
	pro.Creatime = time.Now().Unix()
	id, err := AddDocument(pro)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "添加成功", "id": fmt.Sprintf("%d", id)}
	}
	this.ServeJSON()

}

type DocumentlistController struct {
	controllers.IndexController
}

func (this *DocumentlistController) Get() {
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
	CountDoc := CountDocument(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, CountDoc)
	_, _, doc := ListDocument(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["doc"] = doc
	this.Data["countdoc"] = CountDoc

	this.Data["parentid"] = pid
	this.TplName = "mission/document.tpl"
}

type DocclassTree struct {
	controllers.IndexController
}

func (this *DocclassTree) Get() {
	pid, _ := this.GetInt64("pid")
	tree := GetChildTree(pid, "")
	if len(tree) > 0 {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "success", "data": tree}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "feild"}
	}
	this.ServeJSON()
}

type DocumentDetail struct {
	controllers.IndexController
}

func (this *DocumentDetail) Get() {
	id, _ := this.GetInt64("id")
	doc, err := GetDocument(id)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "success", "data": doc}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "feild"}
	}
	this.ServeJSON()
}

//我的档案资料
type DocumentMy struct {
	controllers.BaseController
}

func (this *DocumentMy) Get() {

}

type DocumentSubtree struct {
	controllers.BaseController
}

func (this *DocumentSubtree) Get() {
	//id,_:=this.GetInt64("id")
	pid, _ := this.GetInt64("pid")

	fmt.Println("sdasdasdas:::::", pid, this.UserUserId, 0)
	tree := GetMyChildTree(pid, this.UserUserId, 0)

	fmt.Println("tree", tree)

	newTree := TreeFlater(tree)
	if newTree != nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "success", "data": newTree}

	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "faild"}
	}
	this.ServeJSON()
}

type DocumentSubview struct {
	controllers.BaseController
}

func (this *DocumentSubview) Get() {
	this.TplName = "mission/documentsub-view.tpl"
}

type DocumentSub struct {
	controllers.BaseController
}

func (this *DocumentSub) Get() {
	pid, _ := this.GetInt64("pid")

	id, _ := this.GetInt64("id")

	if pid == 0 || id == 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "faild"}
		this.ServeJSON()
	}

	document, _ := GetDocument(pid)

	condArr2 := make(map[string]interface{})
	condArr2["keywords"] = ""
	condArr2["missionid"] = int64(0)
	condArr2["missionmyid"] = id
	condArr2["aid"] = int64(0)
	condArr2["types"] = 0
	_, _, files := ListFiles(condArr2, 1, 100)
	missionmy := GetMissionMy(id)
	this.Data["doc"] = document
	this.Data["files"] = files
	this.Data["missionmy"] = missionmy

	//this.TplName="mission/documentsub-form.tpl"
	this.TplName = "mission/documentsub-form.tpl"

}
func (this *DocumentSub) Post() {
	id, _ := this.GetInt64("id")
	fmt.Println("id:", id)
	feedback := this.GetString("feedback")
	detail := this.GetString("detail")
	if !(id > int64(0)) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "任务选择错误"}
		this.ServeJSON()
	}

	missionmy := GetMissionMy(id)

	if missionmy.Id != id {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无法提交他人任务"}
		this.ServeJSON()
	}

	headers, err := this.Ctx.Request.MultipartForm.File["attachment"]
	if len(headers) > 0 {

		for _, h := range headers {
			tmp, err := h.Open()
			fmt.Println("tmp:", tmp)
			var filepath string
			now := time.Now()
			dir := "./static/uploadfile/" + strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
			err1 := os.MkdirAll(dir, 0755)
			if err1 != nil {
				this.Data["json"] = map[string]interface{}{"code": 0, "message": "目录权限不够"}
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
			if err != nil {
				this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
				this.ServeJSON()
				return
			} else {
				//this.SaveToFile("imgFile", "./static/uploadfile/"+h.Filename)
				//this.SaveToFile("attachment", dir+"/"+newName)
				f, err := os.OpenFile(dir+"/"+newName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
				if err != nil {
					this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
					this.ServeJSON()
					return
				}
				defer f.Close()
				io.Copy(f, tmp)
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

//档案审核

type DocAudit struct {
	controllers.BaseController
}

func (this *DocAudit) Get() {
	//权限检测
	/*if !strings.Contains(this.GetSession("userPermission").(string), "mission-manage") {
		this.Redirect("/", 302)
		return
		//this.Abort("401")
	}*/
	page, err := this.GetInt("p")
	status, _ := this.GetInt64("status")
	keywords := this.GetString("keywords")
	if err != nil {
		page = 1
	}

	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	condArr := make(map[string]string)
	condArr["keywords"] = keywords

	countMission := CountMissionDocAudit(status, condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, countMission)
	_, _, missions := GetMissionDocAudit(status, page, offset, condArr)

	fmt.Println("missions:", missions)
	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["missions"] = missions
	this.Data["status"] = status

	this.Data["countMission"] = countMission

	this.TplName = "mission/doc-audit.tpl"

}

//档案审核提交

type DocumentArraignmentController struct {
	controllers.BaseController
}

func (this *DocumentArraignmentController) Get() {
	missionmyidstr := this.GetString("id")
	missionmyid, _ := strconv.ParseInt(missionmyidstr, 10, 64)
	missionmy := GetMissionMy(missionmyid)

	var condArr map[string]interface{}

	condArr = make(map[string]interface{})
	condArr["missionmyid"] = missionmyid
	condArr["aid"] = int64(0)
	condArr["missionid"] = int64(0)
	condArr["types"] = 0
	condArr["keywords"] = ""

	fnums, _, files := ListFiles(condArr, 1, 100)
	document, _ := GetDocument(missionmy.Missionid)
	this.Data["files"] = files
	this.Data["fnums"] = fnums
	this.Data["missionmy"] = missionmy
	this.Data["document"] = document

	this.TplName = "mission/docarraignment-form.tpl"

}
func (this *DocumentArraignmentController) Post() {
	missionmyidstr := this.GetString("id")
	missionmyid, _ := strconv.ParseInt(missionmyidstr, 10, 64)
	arraignment, _ := this.GetInt64("arraignment")
	if !(arraignment > 0) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请更改审核状态"}

	} else if !(arraignment == 2 || arraignment == 3 || arraignment == 4) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请更改审核状态"}
	} else {
		err := ChangeArraignment(missionmyid, arraignment)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
		} else {
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "审核成功"}
		}
	}
	this.ServeJSON()
}
