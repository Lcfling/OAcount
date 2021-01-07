package news

import (
	"github.com/Lcfling/OAcount/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

//消息列表
type NewsList struct {
	Id        int64
	Classid   int64
	Title     string
	Content   string
	Creatime  int64
	Classname string
}

//消息表
type News struct {
	Id       int64
	Classid  int64
	Title    string
	Content  string
	Creatime int64
}

//类型表
type Classic struct {
	Id        int64
	Classname string
	Parentid  int64
	Sorts     int64
}

func (this *News) TableName() string {
	return models.TableName("news")
}
func (this *Classic) TableName() string {
	return models.TableName("classic")
}

func init() {
	orm.RegisterModel(new(News), new(Classic))
}

//消息列表
func ListNews(page int, offset int) (error, []NewsList) {

	var News []NewsList
	o := orm.NewOrm()
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("p.id,p.classid,p.title,p.content,p.creatime, c.classname").From("pms_news AS p").
		InnerJoin("pms_classic AS c").On("p.classid = c.id").OrderBy("p.id").Desc().Limit(offset).Offset(start)

	sql := qb.String()
	_, err := o.Raw(sql).QueryRows(&News)
	return err, News
}

//消息数量
func CountNews() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("news"))
	qs = qs.RelatedSel()
	num, _ := qs.Count()
	return num

}

// 通过id查询消息
func GetNews(id int64) (error, []NewsList) {
	var News []NewsList
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("p.id,p.classid,p.title,p.content,p.creatime, c.classname").From("pms_news AS p").
		InnerJoin("pms_classic AS c").On("p.classid = c.id").Where("p.id=?")
	sql := qb.String()
	_, err := o.Raw(sql, id).QueryRows(&News)
	return err, News
}

//添加消息
func AddNews(Classid int64, Title, Content string) (int64, error) {
	o := orm.NewOrm()
	News1 := new(News)
	News1.Classid = Classid
	News1.Title = Title
	News1.Content = Content
	News1.Creatime = time.Now().Unix()
	id, err := o.Insert(News1)
	return id, err
}

//修改消息
func UpdNews(Newsid, Classid int64, Title, Content string) int {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE pms_news SET classid = ?,title=?,content=? WHERE id = ?", Classid, Title, Content, Newsid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return int(num)
	} else {
		return 0
	}
}

//删除消息
func DeleNews(Newsid int64) int {
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM pms_news WHERE id = ?", Newsid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return int(num)
	} else {
		return 0
	}
}

//获取消息类型列表
func GetClassic() []Classic {
	var Classic []Classic
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("c.id,c.classname,c.parentid,c.sorts").From("pms_classic AS c")
	sql := qb.String()
	_, _ = o.Raw(sql).QueryRows(&Classic)
	return Classic
}

//通过id获取消息类型
func GetClassicInfo(classicid int64) []Classic {
	var Classic []Classic
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("c.id,c.classname,c.parentid,c.sorts").From("pms_classic AS c").Where("c.id=?")
	sql := qb.String()
	_, _ = o.Raw(sql, classicid).QueryRows(&Classic)
	return Classic
}

//添加类型
func AddClassic(Classname string) (int64, error) {
	o := orm.NewOrm()
	Classic := new(Classic)
	Classic.Classname = Classname
	id, err := o.Insert(Classic)
	return id, err
}

//修改消息
func UpdClassic(Classid int64, Classname string) int {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE pms_classic SET classname = ? WHERE id = ?", Classname, Classid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return int(num)
	} else {
		return 0
	}
}

//删除类型
func DeleClassic(Newsid int64) int {
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM pms_classic WHERE id = ?", Newsid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return int(num)
	} else {
		return 0
	}
}
