package files

import (
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/files"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
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
