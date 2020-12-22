package files

import (
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/files"
	"github.com/Lcfling/OAcount/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type FilesManageController struct {
	controllers.BaseController
}

func (this *FilesManageController) Get() {
	missionmyid, _ := this.GetInt64("missionmyid")
	aid, _ := this.GetInt64("aid")
	missionid, _ := this.GetInt64("aid")
	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}
	page, err := this.GetInt("p")
	if err != nil {
		page = 1
	}
	types, _ := this.GetInt("types")
	keywords := this.GetString("keywords")
	condArr := make(map[string]interface{})

	condArr["keywords"] = keywords
	condArr["types"] = types
	condArr["missionmyid"] = missionmyid
	condArr["aid"] = aid
	condArr["missionid"] = missionid

	countFiles := CountFiles(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, countFiles)
	_, _, files := ListFiles(condArr, page, offset)
	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["files"] = files
	this.Data["countFiles"] = countFiles

	this.TplName = "files/files.tpl"
}

//手机端上传
type FilesUploadController struct {
	controllers.UserBaseController
}
type fileslist struct {
	id       int64
	filetype int
	filename string
	filepath string
}

func (this *FilesUploadController) Post() {
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	//headers,err:=this.GetFiles("attachment")
	headers, err := this.Ctx.Request.MultipartForm.File["attachment"]
	//fmt

	var list []*fileslist
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
				f := new(fileslist)
				f.id = id
				f.filename = filename
				f.filetype = fileType
				f.filepath = filepath
				list = append(list, f)
				//Addfile(0, id, fileType, 0, filename, filepath)
			}
		}
	}
	//删掉原先的

	//全部存储
	for _, k := range list {
		_, err1 := Addfile(0, k.id, k.filetype, 0, k.filename, k.filepath)
		if err1 != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
			this.ServeJSON()
		}
	}

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "success"}
	this.ServeJSON()

}

//根据文件id删除文件
type DeleteController struct {
	controllers.BaseController
}

func (this *DeleteController) Post() {
	id, _ := this.GetInt64("id")
	err := DeleteFile(id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": err.Error(), "data": nil}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "success", "data": nil}
	}
	this.ServeJSON()

}

//根据文件id删除文件
type DeleteMobileController struct {
	controllers.UserBaseController
}

func (this *DeleteMobileController) Post() {
	id, _ := this.GetInt64("id")
	err := DeleteFile(id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": err.Error(), "data": nil}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "success", "data": nil}
	}
	this.ServeJSON()

}
