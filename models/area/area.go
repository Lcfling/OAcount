package area

import (
	"fmt"
	"github.com/Lcfling/OAcount/models"
	. "github.com/Lcfling/OAcount/models/tags"
	"github.com/Lcfling/OAcount/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

type Area struct {
	Id        int64
	Parentid  int64
	Jstatus   int64
	Name      string
	Tags      string
	Locations string
	Owner     int64
	Coler     string
	Creatime  int64
	Imgurl    string
}

type AreaList struct {
	Id        int64
	Parentid  int64
	Jstatus   int64
	Name      string
	Tags      string
	Locations string
	Owner     int64
	Coler     string
	Creatime  int64
	Child     []AreaList
}

//仅为打卡用
type Daka struct {
	Id        int64
	Parentid  int64
	Jstatus   int64
	Name      string
	Tags      string
	Locations string
	Owner     int64
	Coler     string
	Creatime  int64
	Daka      int
}

//仅为达标率用
type Pass struct {
	Id        int64
	Parentid  int64
	Jstatus   int64
	Name      string
	Tags      string
	Locations string
	Owner     int64
	Coler     string
	Creatime  int64
	PassRate  float64
}

//仅为问卷调查用
type Ques struct {
	Id        int64
	Parentid  int64
	Jstatus   int64
	Name      string
	Tags      string
	Locations string
	Owner     int64
	Coler     string
	Creatime  int64
	Question  float64
}

type TagsA struct {
	Id  int64
	Aid int64
}

func (this *Area) TableName() string {
	return models.TableName("area")
}
func init() {
	orm.RegisterModel(new(Area))
}
func (this *Area) GetArea(Id int64) {

}
func AddArea(a Area) (int64, error) {
	o := orm.NewOrm()
	areaInfo := new(Area)
	areaInfo.Parentid = a.Parentid
	areaInfo.Jstatus = a.Jstatus
	areaInfo.Locations = a.Locations
	areaInfo.Owner = a.Owner
	areaInfo.Tags = a.Tags
	areaInfo.Coler = a.Coler
	areaInfo.Name = a.Name
	areaInfo.Creatime = time.Now().Unix()
	aid, err := o.Insert(areaInfo)
	if aid > 0 {
		if a.Tags != "" {
			taglist := strings.Split(a.Tags, ",")

			for _, v := range taglist {
				if v != "" {
					tid, err := AddTags(v)
					if err != nil {
						continue
					}
					go AddTagsarea(aid, tid)
				}
			}
		}
	}
	return aid, err
}
func CountArea(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("area"))
	cond := orm.NewCondition()

	if condArr["keywords"] != "" {
		cond = cond.And("name__icontains", condArr["keywords"])
	}
	if condArr["parentid"] != "" {
		cond = cond.And("parentid", condArr["parentid"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}
func ListArea(condArr map[string]string, page int, offset int) (num int64, err error, ops []Area) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("area"))
	cond := orm.NewCondition()

	if condArr["keywords"] != "" {
		cond = cond.And("name__icontains", condArr["keywords"])
	}
	if condArr["parentid"] != "" {
		cond = cond.And("parentid", condArr["parentid"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	qs = qs.OrderBy("id")
	var areas []Area
	num, errs := qs.Limit(offset, start).All(&areas)
	return num, errs, areas
}
func GetAreaName(id int64) string {
	var err error
	var name string

	err = utils.GetCache("GetAreaName.id."+fmt.Sprintf("%d", id), &name)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var areas Area
		o := orm.NewOrm()
		o.QueryTable(models.TableName("area")).Filter("id", id).One(&areas, "name")
		name = areas.Name
		if "" == name {
			name = "一级区域"
		}
		utils.SetCache("GetAreaName.id."+fmt.Sprintf("%d", id), name, cache_expire)
	}
	return name
}
func GetArea(id int64) (Area, error) {
	var area Area
	var err error

	err = utils.GetCache("GetArea.id."+fmt.Sprintf("%d", id), &area)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		o := orm.NewOrm()
		area = Area{Id: id}
		err = o.Read(&area)
		utils.SetCache("GetArea.id."+fmt.Sprintf("%d", id), area, cache_expire)
	}
	return area, err
}

func GetUsersArea(userid int64) []Area {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("area"))
	cond := orm.NewCondition()
	cond = cond.And("userid", userid)
	var areas []Area
	_, errs := qs.All(&areas)
	if errs != nil {
		return nil
	} else {
		return areas
	}
}
func UpdateArea(id int64, area Area) error {
	var areaold Area
	o := orm.NewOrm()
	areaold = Area{Id: id}
	o.Read(&areaold)
	areaold.Name = area.Name
	areaold.Jstatus = area.Jstatus
	areaold.Locations = area.Locations
	areaold.Owner = area.Owner
	areaold.Coler = area.Coler
	areaold.Imgurl = area.Imgurl
	var err error
	_, err = o.Update(&areaold)
	DeleteTagArea(id)

	if id > 0 {
		if area.Tags != "" {
			taglist := strings.Split(area.Tags, ",")

			for _, v := range taglist {
				if v != "" {
					tid, err := AddTags(v)
					if err != nil {
						continue
					}
					go AddTagsarea(id, tid)
				}
			}
		}
	}

	return err
}
func DeleteArea(id int64) error {
	o := orm.NewOrm()
	ids := strconv.FormatInt(id, 10)
	_, err := o.Raw("DELETE FROM " + models.TableName("area") + " WHERE id =" + ids + "").Exec()
	return err
}
func DeleteTagArea(id int64) error {
	o := orm.NewOrm()
	ids := strconv.FormatInt(id, 10)
	_, err := o.Raw("DELETE FROM " + models.TableName("tags_area") + " WHERE aid =" + ids + "").Exec()
	return err
}

func GetChild(pid int64, types string) []AreaList {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("area"))
	cond := orm.NewCondition()
	cond = cond.And("parentid", pid)
	if types != "" {
		cond = cond.And("jstatus", types)
	}
	qs = qs.SetCond(cond)
	var areas []Area
	num, _ := qs.All(&areas)
	if num > 0 {
		var areaLists []AreaList
		for _, v := range areas {
			var areaList AreaList
			areaList.Parentid = v.Parentid
			areaList.Tags = v.Tags
			areaList.Jstatus = v.Jstatus
			areaList.Id = v.Id
			areaList.Name = v.Name
			areaList.Creatime = v.Creatime
			areaList.Owner = v.Owner
			areaList.Coler = v.Coler
			areaList.Locations = v.Locations
			areaList.Child = GetChild(v.Id, types)
			areaLists = append(areaLists, areaList)
		}
		return areaLists
	} else {
		return nil
	}

}

//通过pid 获取下面所有地区id数组
func GetAllAreaIdByPid(areaList []AreaList) (arr []string) {
	for _, v := range areaList {
		arr = append(arr, strconv.FormatInt(v.Id, 10))
		if v.Child != nil {
			arr = append(arr, GetAllAreaIdByPid(v.Child)...)
		}

	}
	return
}

func GetAllByArray(arr []string) (nums int64, err error, ar []Area) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("area"))
	cond := orm.NewCondition()

	cond = cond.And("id__in", arr)

	qs = qs.SetCond(cond)

	qs = qs.OrderBy("id")
	var areas []Area
	num, errs := qs.All(&areas)
	return num, errs, areas
}

//查询社区所有人员
func AllArea() (error, []Area) {

	var Area []Area
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("p.id,p.owner").From("pms_area AS p")
	sql := qb.String()
	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&Area)
	return err, Area
}

//查询类型对应的社区
func TagsArea(tage int64) (error, []TagsA) {

	var Tags []TagsA
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("p.aid").From("pms_tags_area AS p").
		Where("p.tid=?")
	sql := qb.String()
	o := orm.NewOrm()
	_, err := o.Raw(sql, tage).QueryRows(&Tags)
	return err, Tags

}
