package news

import (
	"fmt"
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/news"
	"strconv"
)

type NewsManageController struct {
	controllers.BaseController
}

func (this *NewsManageController) Get() {

	//page, err := this.GetInt("p")
	status := this.GetString("status")
	keywords := this.GetString("keywords")
	condArr := make(map[string]string)
	condArr["status"] = status
	condArr["keywords"] = keywords
	//类型消息
	_, news := ListNews()
	fmt.Println("news:", news)
	this.Data["news"] = news
	this.TplName = "news/news.tpl"
}
func (this *NewsManageController) Post() {

}

type NewsAddController struct {
	controllers.BaseController
}

func (this *NewsAddController) Get() {
	//权限检测
	Classic := GetClassic()
	fmt.Println("Classic:", Classic)
	this.Data["classic"] = Classic
	this.TplName = "news/news-form.tpl"
}

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
	content := this.GetString("content")
	if "" == content {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写项目描述"}
		this.ServeJSON()
		return
	}
	classid64, _ := strconv.ParseInt(classid, 10, 64)
	_, err := AddNews(classid64, title, content)
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
	fmt.Println("classic:", classic)
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
