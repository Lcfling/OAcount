package news

import (
	"fmt"
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/news"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

type NewsManageController struct {
	controllers.BaseController
}

//消息页面
func (this *NewsManageController) Get() {

	page, err1 := this.GetInt("p")
	if err1 != nil {
		page = 1
	}
	offset, err2 := beego.AppConfig.Int("pageoffset")
	if err2 != nil {
		offset = 15
	}
	// 统计消息数量
	countNews := CountNews()
	fmt.Println("countNews:", countNews)
	paginator := pagination.SetPaginator(this.Ctx, offset, countNews)
	//类型消息
	_, news := ListNews(page, offset)
	this.Data["paginator"] = paginator
	this.Data["news"] = news
	this.Data["countNews"] = countNews
	this.TplName = "news/news.tpl"
}

func (this *NewsManageController) Post() {

}

type NewsAddController struct {
	controllers.BaseController
}

//添加新消息页面
func (this *NewsAddController) Get() {
	//权限检测
	Classic := GetClassic()
	this.Data["classic"] = Classic
	this.TplName = "news/news-form.tpl"
}

//添加新消息
func (this *NewsAddController) Post() {
	//权限检测
	//if !strings.Contains(this.GetSession("userPermission").(string), "mission-add") {
	//	this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
	//	this.ServeJSON()
	//	return
	//}

	title := this.GetString("title")
	if "" == title {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写消息名称"}
		this.ServeJSON()
		return
	}
	classid := this.GetString("classid")
	if "" == classid {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写消息类型"}
		this.ServeJSON()
		return
	}
	desc := this.GetString("desc")
	if "" == desc {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写项目描述"}
		this.ServeJSON()
		return
	}

	classid64, _ := strconv.ParseInt(classid, 10, 64)
	_, err := AddNews(classid64, title, desc)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "项目信息添加成功", "id": fmt.Sprintf("%d", 1)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "项目信息添加失败"}
	}
	this.ServeJSON()
}

//消息类型控制器
type NewsClassicController struct {
	controllers.BaseController
}

//消息类型列表页
func (this *NewsClassicController) Get() {

	//page, err := this.GetInt("p")
	status := this.GetString("status")
	keywords := this.GetString("keywords")
	condArr := make(map[string]string)
	condArr["status"] = status
	condArr["keywords"] = keywords
	//类型消息
	classic := GetClassic()
	this.Data["classic"] = classic
	this.TplName = "news/classic.tpl"
}

//消息类型列表页post请求
func (this *NewsClassicController) Post() {

}

//消息类型+新类型页面
type NewsClassicAddController struct {
	controllers.BaseController
}

//消息类型+新类型post请求
func (this *NewsClassicAddController) Get() {
	//权限检测
	this.TplName = "news/classic-form.tpl"
}

//消息类型添加
func (this *NewsClassicAddController) Post() {
	//权限检测
	//if !strings.Contains(this.GetSession("userPermission").(string), "mission-add") {
	//	this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
	//	this.ServeJSON()
	//	return
	//}
	classname := this.GetString("classname")
	if "" == classname {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写消息类型"}
		this.ServeJSON()
		return
	}
	_, err := AddClassic(classname)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "项目信息添加成功", "id": fmt.Sprintf("%d", 1)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "项目信息添加失败"}
	}
	this.ServeJSON()
}

//消息类型删除
type NewsAjaxClassicDeleteController struct {
	controllers.BaseController
}

//消息类型删除 post请求
func (this *NewsAjaxClassicDeleteController) Post() {
	//任务id
	id, _ := this.GetInt64("id")
	fmt.Println("id:", id)
	//sql删除操作
	num := DeleNews(id)
	if num > 0 {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功", "id": fmt.Sprintf("%d", 1)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}

//消息编辑
type NewsEditsController struct {
	controllers.BaseController
}

//消息编辑
func (this *NewsEditsController) Get() {

	//任务id
	id := this.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	//类型消息
	_, news := GetNews(id64)
	fmt.Println("news:", news)

	//权限检测
	Classic := GetClassic()
	this.Data["news"] = news
	this.Data["classic"] = Classic
	this.TplName = "news/news-edits.tpl"
}

//消息编辑 post请求
func (this *NewsEditsController) Post() {
	//任务id
	id := this.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	title := this.GetString("title")
	if title == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写消息名称"}
		this.ServeJSON()
		return
	}
	classid := this.GetString("classid")
	if classid == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写消息类型"}
		this.ServeJSON()
		return
	}
	desc := this.GetString("desc")
	if desc == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写项目描述"}
		this.ServeJSON()
		return
	}

	classid64, _ := strconv.ParseInt(classid, 10, 64)
	//进行sql修改
	num := UpdNews(id64, classid64, title, desc)
	if num > 0 {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "消息管理编辑成功", "id": fmt.Sprintf("%d", 1)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "消息管理编辑失败"}
	}
	this.ServeJSON()
}

//消息删除
type NewsAjaxDeleteController struct {
	controllers.BaseController
}

//消息删除 post请求
func (this *NewsAjaxDeleteController) Post() {
	//任务id
	id, _ := this.GetInt64("id")
	//sql删除操作
	num := DeleNews(id)
	if num > 0 {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功", "id": fmt.Sprintf("%d", 1)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}

//类型编辑
type NewsClassicEditController struct {
	controllers.BaseController
}

//类型编辑
func (this *NewsClassicEditController) Get() {

	//任务id
	id := this.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	//类型消息
	classic := GetClassicInfo(id64)
	fmt.Println("news:", classic)
	this.Data["classic"] = classic
	this.TplName = "news/classic-edit.tpl"
}

//类型编辑 post请求
func (this *NewsClassicEditController) Post() {
	//任务id
	id := this.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	classname := this.GetString("classname")
	if classname == "" {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写类型名称"}
		this.ServeJSON()
		return
	}
	//进行sql修改
	num := UpdClassic(id64, classname)
	if num > 0 {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "消息类型编辑成功", "id": fmt.Sprintf("%d", 1)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "消息类型编辑失败"}
	}
	this.ServeJSON()
}
