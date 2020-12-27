package area

import (
	"fmt"
	"github.com/Lcfling/OAcount/models"
	. "github.com/Lcfling/OAcount/models/program"
	"github.com/astaxie/beego/orm"
	"strconv"
)

//根据aid返回区域问卷评价分数据
func GetAreaQuestionAver(aid int) []Ques {
	area := getAreaListByAid(aid)
	for k, v := range area {
		aids := getAidsByaid(v.Id)
		if aids != nil {
			question := getAverage(aids)
			area[k].Question = question
		}
	}
	return area
}

//根据aid获取区域数据
func getAreaListByAid(aid int) []Ques {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("area"))
	cond := orm.NewCondition()
	cond = cond.And("parentid", aid)
	qs = qs.SetCond(cond)
	var areas []Area
	num, _ := qs.All(&areas)
	if num > 0 {
		var areaLists []Ques
		for _, v := range areas {
			var area Ques
			area.Parentid = v.Parentid
			area.Tags = v.Tags
			area.Jstatus = v.Jstatus
			area.Id = v.Id
			area.Name = v.Name
			area.Creatime = v.Creatime
			area.Owner = v.Owner
			area.Coler = v.Coler
			area.Locations = v.Locations
			area.Question = 0
			areaLists = append(areaLists, area)
		}
		return areaLists
	} else {
		return nil
	}
}

//问卷平均分
func getAverage(aids []string) float64 {
	count := scoreCount(aids)
	sum := scoreSum(aids)
	if count == 0 {
		return 0
	}
	doneRate := float64(sum) / float64(count)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", doneRate), 64)

	return value
}

//根据区域获取area id数组
func getAidsByaid(aid int64) []string {
	AreaList := GetChild(aid, "")
	aids := GetAllAreaIdByPid(AreaList)
	return aids
}

//根据aids计算问卷调查的数量
func scoreCount(aids []string) float64 {
	var all int64
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("sub"))
	cond := orm.NewCondition()
	cond = cond.And("community__in", aids) //aids数组
	all, _ = qs.SetCond(cond).Count()
	return float64(all)
}
func scoreSum(aids []string) float64 {
	questionList := questionList(aids)
	var sum float64
	for _, v := range questionList {
		sum = sum + v.Score
	}
	return sum
}

//根据aids计算问卷调查的评价数据的总数
func questionList(aids []string) []Question {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("sub"))
	cond := orm.NewCondition()
	cond = cond.And("community__in", aids)
	qs = qs.SetCond(cond)
	var questions []Sub
	num, _ := qs.All(&questions, "Score")
	if num > 0 {
		var questionLists []Question
		for _, v := range questions {
			var question Question
			question.Score = float64(v.Score)
			questionLists = append(questionLists, question)
		}
		return questionLists
	} else {
		return nil
	}
}
