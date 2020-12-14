package tags

import (
	"github.com/Lcfling/OAcount/controllers"
	. "github.com/Lcfling/OAcount/models/tags"
)

type TagsAllController struct {
	controllers.IndexController
}

func (this *TagsAllController) Get() {
	tags := GetTagsAll()

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "区域添加成功", "data": tags}
	this.ServeJSON()
}
