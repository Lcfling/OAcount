package mission

import (
	"encoding/json"
	"fmt"
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/area"
	. "github.com/Lcfling/OAcount/models/files"
	. "github.com/Lcfling/OAcount/models/mission"
	. "github.com/Lcfling/OAcount/models/tags"
	"strconv"
)

type SendTaskController struct {
	controllers.BaseController
}

func (this *SendTaskController) Get() {
	tags := GetTagsAll()
	tjson, _ := json.Marshal(tags)

	s := GetChild(0, "")
	json, _ := json.Marshal(s)

	this.Data["area"] = string(json)
	this.Data["tage"] = string(tjson)

	this.TplName = "mission/sendtask.tpl"
}

//任务下发页面post
func (this *SendTaskController) Post() {

	//任务id
	id := this.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)

	//全部
	all := this.GetString("all")
	//点位
	checkareas := this.GetStrings("checkareas")
	//类型社区
	tagss := this.GetStrings("tags")

	//获取任务信息
	var mission Mission
	mission, _ = GetMission(id64)
	//---------------------------------------------------------------------------------
	//判断是否全部下发
	if !(all == "") {
		//全部下发   查询所有社区人员
		err, AllArea := AllArea()
		if err == nil {
			//对社区全部人员进行任务下发
			for _, value := range AllArea {
				if value.Owner > int64(0) {
					//fmt.Println("社区人员ower:", value.Owner)
					//插入任务
					go AddMyMission(id64, value.Owner, value.Id, mission.Types)
					//插入我的消息
				}
			}
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "success"}
			this.ServeJSON()
		}

	}
	//---------------------------------------------------------------------------------

	//根据类型社区进行任务下发
	if !(len(tagss) == 0) {
		for _, v := range tagss {
			v64, _ := strconv.ParseInt(v, 10, 64)
			//查询任务相关的任务文件  进行绑定
			go TageFile(v64, id64)
			//查询对应的社区
			err, TagsArea := TagsArea(v64)
			if err == nil {
				//插入任务
				for _, value := range TagsArea {
					//查询对应的社区信息
					AreaUser, _ := GetArea(value.Aid)
					if AreaUser.Owner > int64(0) {
						//插入任务
						go AddMyMission(id64, AreaUser.Owner, value.Id, mission.Types)
						//插入我的消息
					}
				}
			}

		}
	}
	//---------------------------------------------------------------------------------
	// 单个人员进行下发
	if !(len(checkareas) == 0) {
		for _, value := range checkareas {
			value64, _ := strconv.ParseInt(value, 10, 64)
			Area, err := GetArea(value64)
			if err == nil {
				//插入任务
				if Area.Owner > int64(0) {
					//插入任务
					go AddMyMission(id64, Area.Owner, Area.Id, mission.Types)

				}
			}

		}
	}

	//---------------------------------------------------------------------------------
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "success"}
	this.ServeJSON()
}

//档案 实地 任务下发
type SendTaskuniController struct {
	controllers.BaseController
}

func (this *SendTaskuniController) Get() {
	tags := GetTagsAll()
	tjson, _ := json.Marshal(tags)

	s := GetChild(0, "")
	json, _ := json.Marshal(s)

	this.Data["area"] = string(json)
	this.Data["tage"] = string(tjson)

	this.TplName = "mission/sendtask.tpl"
}

//任务下发页面post
func (this *SendTaskuniController) Post() {

	//任务id
	id := this.GetString("id")
	id64, _ := strconv.ParseInt(id, 10, 64)

	//全部
	all := this.GetString("all")
	//点位
	checkareas := this.GetStrings("checkareas")
	//类型社区
	tagss := this.GetStrings("tags")

	Types, _ := this.GetInt("types")
	//获取任务信息

	//---------------------------------------------------------------------------------
	//判断是否全部下发
	if !(all == "") {
		//全部下发   查询所有社区人员
		err, AllArea := AllArea()
		if err == nil {
			//对社区全部人员进行任务下发
			for _, value := range AllArea {
				if value.Owner > int64(0) {
					//fmt.Println("社区人员ower:", value.Owner)
					//插入任务
					go AddMyMission(id64, value.Owner, value.Id, Types)
					//插入我的消息
				}
			}
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "success"}
			this.ServeJSON()
		}

	}
	//---------------------------------------------------------------------------------

	//根据类型社区进行任务下发
	if !(len(tagss) == 0) {
		for _, v := range tagss {
			v64, _ := strconv.ParseInt(v, 10, 64)
			//查询任务相关的任务文件  进行绑定
			//go TageFile(v64, id64)
			//查询对应的社区
			err, TagsArea := TagsArea(v64)
			if err == nil {
				//插入任务
				for _, value := range TagsArea {
					//查询对应的社区信息
					AreaUser, _ := GetArea(value.Aid)
					if AreaUser.Owner > int64(0) {
						//插入任务
						fmt.Println("+++++++++++", id64, AreaUser.Owner, value.Aid, Types)
						go AddMyMission(id64, AreaUser.Owner, value.Aid, Types)
						//插入我的消息
					}
				}
			}

		}
	}
	//---------------------------------------------------------------------------------
	// 单个人员进行下发
	if !(len(checkareas) == 0) {
		for _, value := range checkareas {
			value64, _ := strconv.ParseInt(value, 10, 64)
			Area, err := GetArea(value64)
			if err == nil {
				//插入任务
				if Area.Owner > int64(0) {
					//插入任务
					go AddMyMission(id64, Area.Owner, Area.Id, Types)

				}
			}

		}
	}

	//---------------------------------------------------------------------------------
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "success"}
	this.ServeJSON()
}
