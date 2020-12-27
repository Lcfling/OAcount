package area

import (
	"fmt"
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/area"
	. "github.com/Lcfling/OAcount/models/users"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
	"strings"
)

type AreaMangerController struct {
	controllers.BaseController
}

func (this *AreaMangerController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "area-manage") {
		this.Abort("401")
	}
	page, err := this.GetInt("p")
	if err != nil {
		page = 1
	}
	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}
	keywords := this.GetString("keywords")
	condArr := make(map[string]string)
	condArr["keywords"] = keywords
	parentid := this.GetString("parentid")
	if parentid == "" {
		parentid = "0"
	}
	condArr["parentid"] = parentid
	CountArea := CountArea(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, CountArea)
	_, _, areas := ListArea(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["areas"] = areas
	this.Data["countArea"] = CountArea
	if parentid == "" {
		this.Data["parentid"] = "0"
	} else {
		this.Data["parentid"] = parentid
	}

	//一级栏目
	condArrP := make(map[string]string)
	condArrP["parentid"] = "0"
	_, _, parentareas := ListArea(condArrP, 0, 100)
	this.Data["parentareas"] = parentareas

	fmt.Println(this.Data["myuser"], "dasdasdasdasdsad")
	this.TplName = "area/area.tpl"
}
func (this *AreaMangerController) Post() {

}

type AreaListController struct {
	controllers.BaseController
}

func (this *AreaListController) Get() {

}

//添加区域
type AreaAddController struct {
	controllers.BaseController
}

func (this *AreaAddController) Get() {
	var users []UsersFind
	_, users = ListMyUser(this.UserUserId)
	var areaInfo Area
	areaInfo.Owner = 0
	areaInfo.Jstatus = 0
	this.Data["area"] = areaInfo
	this.Data["myuser"] = users
	fmt.Println(users)
	this.TplName = "area/area-form.tpl"
}
func (this *AreaAddController) Post() {
	//todo 权限检查
	id := this.Ctx.Input.Param(":id")
	tagstr := this.GetString("tags")
	jstatus, _ := this.GetInt64("jstatus")
	parentid, _ := strconv.ParseInt(id, 10, 64)
	locations := this.GetString("locations")
	name := this.GetString("name")
	owner := this.GetString("owner")
	ownerid, _ := strconv.ParseInt(owner, 10, 64)
	coler := this.GetString("coler")
	var areaInfo Area
	areaInfo.Parentid = parentid
	areaInfo.Name = name
	areaInfo.Jstatus = jstatus
	areaInfo.Tags = tagstr
	areaInfo.Locations = locations
	areaInfo.Owner = ownerid
	areaInfo.Coler = coler
	_, err := AddArea(areaInfo)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "区域添加成功", "id": fmt.Sprintf("%d", id), "pid": fmt.Sprintf("%d", parentid)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "区域添加失败;" + err.Error()}
	}
	this.ServeJSON()
}

//编辑区域
type AreaEditController struct {
	controllers.BaseController
}

func (this *AreaEditController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "area-edit") {
		this.Abort("401")
	}
	id := this.Ctx.Input.Param(":id")
	idp, _ := strconv.ParseInt(id, 10, 64)

	var area Area
	area, _ = GetArea(idp)
	this.Data["area"] = area
	var users []UsersFind
	_, users = ListMyUser(this.UserUserId)
	this.Data["myuser"] = users
	this.TplName = "area/area-form.tpl"
}
func (this *AreaEditController) Post() {
	id := this.Ctx.Input.Param(":id")
	idp, _ := strconv.ParseInt(id, 10, 64)
	locations := this.GetString("locations")
	name := this.GetString("name")
	owner := this.GetString("owner")
	coler := this.GetString("coler")
	if name == "" {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "名称不能为空", "id": fmt.Sprintf("%d", id)}
	}
	if locations == "" {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "位置信息不能为空", "id": fmt.Sprintf("%d", id)}
	}
	if owner == "" {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "负责人不能为空", "id": fmt.Sprintf("%d", id)}
	}
	ownerid, _ := strconv.ParseInt(owner, 10, 64)
	var area Area
	area.Name = name
	area.Locations = locations
	area.Owner = ownerid
	area.Coler = coler
	err := UpdateArea(idp, area)
	fmt.Println(err)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功", "id": fmt.Sprintf("%d", idp)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	}
	this.ServeJSON()

}

//删除区域
type AreaDeleteController struct {
	controllers.BaseController
}

func (this *AreaDeleteController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "area-delete") {
		this.Abort("401")
	}
	id, err := this.GetInt64("id")
	if 0 == id || err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择要删除的区域"}
		this.ServeJSON()
		return
	}
	err = DeleteArea(id)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}

type AllAreaController struct {
	controllers.IndexController
}
type AreaShow struct {
	Id        int64
	Name      string
	Locations string
	Owner     int64
	Coler     string
}

func (this *AllAreaController) Post() {

	s := GetChild(0, "0")

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "系统错误", "data": s}

	this.ServeJSON()
}

type GetAllAreaController struct {
	controllers.IndexController
}

func (this *GetAllAreaController) Get() {

	pid, _ := this.GetInt64("pid")
	s := GetChild(pid, "")
	arr := GetAllAreaIdByPid(s)
	_, _, data := GetAllByArray(arr)

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "success", "data": data}

	this.ServeJSON()
}

type DingController struct {
	controllers.IndexController
}

func (this *DingController) Get() {

	InsertNew()
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "打卡成功", "data": ""}
	this.ServeJSON()
}
func (this *DingController) Post() {

	d, _ := GetNew()
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "打卡数据", "data": d}
	this.ServeJSON()
}
