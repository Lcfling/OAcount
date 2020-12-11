package program

import (
	"fmt"
	"github.com/virteman/OAcount/controllers"
	. "github.com/virteman/OAcount/models/program"
	"strconv"
)

type SubjectAddController struct {
	controllers.BaseController
}
func (this *SubjectAddController)Post(){
	pidstr:=this.GetString("pid")
	content:=this.GetString("data")
	pid,_:=strconv.ParseInt(pidstr, 10, 64)
	id,err:=AddSubject(pid,content)
	if err!=nil{
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "添加题目失败","data":""}
		this.ServeJSON()
		return
	}else{
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "添加题目成功","id": fmt.Sprintf("%d", id)}
		this.ServeJSON()
		return
	}
}
type SubjectListController struct {
	controllers.BaseController
}

func (this *SubjectListController) Get() {
	pidstr:=this.GetString(":pid")
	pid,_:=strconv.ParseInt(pidstr, 10, 64)
	program,err:=GetProgram(pid)
	if err!=nil {
		this.Data["msg"] = "找不到测评"
	}
	_,s:=GetList(pid)
	fmt.Println(s)
	this.Data["subject"] = s
	this.Data["program"] = program
	this.TplName="program/subject.tpl"
}