package program

import (
	"encoding/json"
	"fmt"
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego/orm"
)

type Subject struct {
	Id      int64
	Temps   string
	Pid     int64
	Modleid int64
	Sorts   int64
	Content string
}
type Subjects struct {
	Id      int64
	Temps   string
	Pid     int64
	Modleid int64
	Sorts   int64
	Content Content
}
type Content struct {
	Title   string
	Type    string
	Subject []CHoose
}
type CHoose struct {
	Content string //内容
	Check   int
	Count   int64
}

func (this *Subject) TableName() string {
	return models.TableName("subject")
}
func init() {
	orm.RegisterModel(new(Subject))
}
func GetList(pid int64) (error, []Subjects) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("subject"))
	cond := orm.NewCondition()

	cond = cond.And("pid", pid)
	qs = qs.SetCond(cond)

	qs = qs.RelatedSel()

	var subJect []Subject
	qs = qs.OrderBy("-sorts", "-id")
	_, err1 := qs.All(&subJect)
	if err1 != nil {
		return err1, nil
	}
	var subJects []Subjects
	for _, v := range subJect {
		//s:=new(Subjects)
		var s Subjects
		s.Id = v.Id
		s.Temps = v.Temps
		s.Sorts = v.Sorts
		s.Modleid = v.Modleid
		s.Pid = v.Pid
		var c Content
		json.Unmarshal([]byte(v.Content), &c)
		fmt.Println(c)
		s.Content = c
		subJects = append(subJects, s)
	}
	return err1, subJects
}
func AddSubject(pid int64, content string) (id int64, err error) {

	//var s Subject
	s := new(Subject)
	s.Pid = pid
	s.Content = content
	s.Modleid = 0
	s.Sorts = 0
	s.Temps = ""
	o := orm.NewOrm()
	return o.Insert(s)
}
