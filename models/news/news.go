package news

import (
	"github.com/Lcfling/OAcount/models"
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

func init() {
	orm.RegisterModel(new(News), new(Classic))
}

//消息列表
func ListNews() (error, []NewsList) {

	var News []NewsList
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("p.id,p.classid,p.title,p.content,p.creatime, c.classname").From("pms_news AS p").
		InnerJoin("pms_classic AS c").On("p.classid = c.id")
	sql := qb.String()
	_, err := o.Raw(sql).QueryRows(&News)
	return err, News
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

//添加消息
func AddClassic(Classname string) (int64, error) {

	o := orm.NewOrm()
	Classic := new(Classic)
	Classic.Classname = Classname
	id, err := o.Insert(Classic)
	return id, err
}
