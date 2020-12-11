package initial

import (
	"github.com/virteman/OAcount/models/mission"
	"github.com/virteman/OAcount/models/area"
	//"fmt"
	"github.com/virteman/OAcount/models/businesstrips"
	"github.com/virteman/OAcount/models/expenses"
	"github.com/virteman/OAcount/models/goouts"
	"github.com/virteman/OAcount/models/groups"
	"github.com/virteman/OAcount/models/leaves"
	"github.com/virteman/OAcount/models/oagoods"
	"github.com/virteman/OAcount/models/overtimes"
	"github.com/virteman/OAcount/models/projects"
	"github.com/virteman/OAcount/models/users"
	"github.com/virteman/OAcount/utils"
	//"time"

	"github.com/astaxie/beego"
)

func InitTemplate() {
	beego.AddFuncMap("getRealname", users.GetRealname)
	beego.AddFuncMap("getNeedsname", projects.GetProjectNeedsName)
	beego.AddFuncMap("getProjectname", projects.GetProjectName)
	beego.AddFuncMap("getPermissionname", groups.GetPermissiontName)
	beego.AddFuncMap("getAreaname", area.GetAreaName)
	beego.AddFuncMap("getMissionname", mission.GetMissionName)
	beego.AddFuncMap("Ratio", utils.Ratio)
	beego.AddFuncMap("getLeaveProcess", leaves.ListLeaveApproverProcessHtml)
	beego.AddFuncMap("getExpenseProcess", expenses.ListExpenseApproverProcessHtml)
	beego.AddFuncMap("getBusinesstripProcess", businesstrips.ListBusinesstripApproverProcessHtml)
	beego.AddFuncMap("getGooutProcess", goouts.ListGooutApproverProcessHtml)
	beego.AddFuncMap("getOagoodProcess", oagoods.ListOagoodApproverProcessHtml)
	beego.AddFuncMap("getOvertimeProcess", overtimes.ListOvertimeApproverProcessHtml)

	beego.AddFuncMap("getDate", utils.GetDate)
	beego.AddFuncMap("getDateMH", utils.GetDateMH)
	beego.AddFuncMap("getNeedsStatus", utils.GetNeedsStatus)
	beego.AddFuncMap("getNeedsSource", utils.GetNeedsSource)
	beego.AddFuncMap("getNeedsStage", utils.GetNeedsStage)
	beego.AddFuncMap("getTaskStatus", utils.GetTaskStatus)
	beego.AddFuncMap("getTaskType", utils.GetTaskType)
	beego.AddFuncMap("getTestStatus", utils.GetTestStatus)
	beego.AddFuncMap("getLeaveType", utils.GetLeaveType)

	beego.AddFuncMap("getOs", utils.GetOs)
	beego.AddFuncMap("getBrowser", utils.GetBrowser)
	beego.AddFuncMap("getAvatarSource", utils.GetAvatarSource)
	beego.AddFuncMap("getAvatar", utils.GetAvatar)
	beego.AddFuncMap("getAvatarUserid", users.GetAvatarUserid)
	beego.AddFuncMap("getPositionsName", users.GetPositionsNameForUserid)
	beego.AddFuncMap("getDepartmentsName", users.GetDepartmentsNameForUserid)

	beego.AddFuncMap("getEdu", utils.GetEdu)
	beego.AddFuncMap("getWorkYear", utils.GetWorkYear)
	beego.AddFuncMap("getResumeStatus", utils.GetResumeStatus)

	beego.AddFuncMap("getCheckworkType", utils.GetCheckworkType)

	beego.AddFuncMap("getMessageType", utils.GetMessageType)
	beego.AddFuncMap("getMessageSubtype", utils.GetMessageSubtype)

}
