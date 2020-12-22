package mission

import (
	"encoding/json"
	"fmt"
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/area"
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

	all := this.GetString("all")
	fmt.Println("全部:", all)

	checkareas := this.GetString("checkareas")
	fmt.Println("人员:", checkareas)

	tagss := this.GetStrings("tags")
	fmt.Println("社区:", tagss)

	for _, value := range tagss {
		fmt.Println("遍历社区:", value)
	}

	//---------------------------------------------------------------------------------
	//全部下发   查询所有社区人员
	err, AllArea := AllArea()
	if err != nil {
		fmt.Println("查询社区所有人员错误err:", err)
	}
	//对社区人员进行任务下发
	for _, value := range AllArea {
		if value.Owner > int64(0) {
			//fmt.Println("社区人员ower:", value.Owner)
			//插入任务
			go AddMyMission(id64, value.Owner, value.Id)
			//插入我的消息
		}
	}

	//---------------------------------------------------------------------------------
	//根据类型社区进行任务下发
	tags := []string{"卫生站", "教育", "社区"}
	//根据类型社会进行任务下发
	for _, value := range tags {
		err, TagsArea := TagsArea(value)
		if err != nil {
			fmt.Println("没有此类型社区")
		}
		//插入任务
		for _, value := range TagsArea {
			if value.Owner > int64(0) {
				//对社区人员进行任务下发
				//	fmt.Println("社区人员ower:", value.Owner)
				//插入任务
				go AddMyMission(id64, value.Owner, value.Id)
				//插入我的消息
			}

		}
	}

	//---------------------------------------------------------------------------------
	// 单个社区进行下发
	Area64 := []int64{1, 2, 3}
	//根据类型社会进行任务下发
	for _, value := range Area64 {
		Area, err := GetArea(value)
		if err != nil {
			fmt.Println("没有此类型社区管理员")
		} else {
			//插入任务
			if Area.Owner > int64(0) {
				//对社区人员进行任务下发
				//fmt.Println("社区人员ower:", Area.Owner)
				//插入任务
				go AddMyMission(id64, Area.Owner, Area.Id)
				//插入我的消息
			}
		}

	}
	//---------------------------------------------------------------------------------
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "success"}
	this.ServeJSON()
}
