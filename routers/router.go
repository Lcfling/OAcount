package routers

import (
	"github.com/Lcfling/OAcount/controllers"
	"github.com/Lcfling/OAcount/controllers/albums"
	"github.com/Lcfling/OAcount/controllers/area"
	"github.com/Lcfling/OAcount/controllers/businesstrips"
	"github.com/Lcfling/OAcount/controllers/checkworks"
	"github.com/Lcfling/OAcount/controllers/expenses"
	"github.com/Lcfling/OAcount/controllers/files"
	"github.com/Lcfling/OAcount/controllers/goouts"
	"github.com/Lcfling/OAcount/controllers/groups"
	"github.com/Lcfling/OAcount/controllers/knowledges"
	"github.com/Lcfling/OAcount/controllers/leaves"
	"github.com/Lcfling/OAcount/controllers/messages"
	"github.com/Lcfling/OAcount/controllers/mission"
	"github.com/Lcfling/OAcount/controllers/news"
	"github.com/Lcfling/OAcount/controllers/oagoods"
	"github.com/Lcfling/OAcount/controllers/overtimes"
	"github.com/Lcfling/OAcount/controllers/program"
	"github.com/Lcfling/OAcount/controllers/projects"
	"github.com/Lcfling/OAcount/controllers/resumes"
	"github.com/Lcfling/OAcount/controllers/tags"
	"github.com/Lcfling/OAcount/controllers/users"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &users.MainController{})

	//websocket
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

	//用户
	beego.Router("/user/manage", &users.ManageUserController{})
	beego.Router("/user/ajax/status", &users.AjaxStatusUserController{})
	beego.Router("/user/edit/:id", &users.EditUserController{})
	beego.Router("/user/add", &users.AddUserController{})
	beego.Router("/user/avatar", &users.AvatarUserController{})
	beego.Router("/user/ajax/search", &users.AjaxSearchUserController{}) //搜索用户名匹配
	beego.Router("/user/show/:id", &users.ShowUserController{})
	beego.Router("/my/manage", &users.ShowUserController{})
	beego.Router("/user/profile", &users.EditUserProfileController{})

	beego.Router("/user/password", &users.EditUserPasswordController{})

	beego.Router("/user/permission/:id", &users.PermissionController{})

	beego.Router("/login", &users.LoginUserController{})

	beego.Router("/mobile/user/login", &users.MobileLoginController{}) //手机用户信息完善

	//----------------------------------------    登陆
	beego.Router("/mobile/login", &users.LoginUserController{})

	beego.Router("/logout", &users.LogoutUserController{})
	beego.Router("/register", &users.RegisterController{})

	//部门
	beego.Router("/department/manage", &users.ManageDepartmentController{})
	beego.Router("/department/ajax/status", &users.AjaxStatusDepartmentController{})
	beego.Router("/department/edit/:id", &users.EditDepartmentController{})
	beego.Router("/department/add", &users.AddDepartmentController{})

	//职位
	beego.Router("/position/manage", &users.ManagePositionController{})
	beego.Router("/position/ajax/status", &users.AjaxStatusPositionController{})
	beego.Router("/position/edit/:id", &users.EditPositionController{})
	beego.Router("/position/add", &users.AddPositionController{})

	//公告
	beego.Router("/notice/manage", &users.ManageNoticeController{})
	beego.Router("/notice/ajax/status", &users.AjaxStatusNoticeController{})
	beego.Router("/notice/ajax/delete", &users.AjaxDeleteNoticeController{})
	beego.Router("/notice/edit/:id", &users.EditNoticeController{})
	beego.Router("/notice/add", &users.AddNoticeController{})

	//项目
	beego.Router("/project/manage", &projects.ManageProjectController{})
	beego.Router("/project/ajax/status", &projects.AjaxStatusProjectController{})
	beego.Router("/project/edit/:id", &projects.EditProjectController{})
	beego.Router("/project/add", &projects.AddProjectController{})
	beego.Router("/project/:id", &projects.ShowProjectController{})

	beego.Router("/my/project", &projects.MyProjectController{})
	beego.Router("/project/chart/:id", &projects.ChartProjectController{})

	//项目成员
	beego.Router("/project/team/:id", &projects.TeamProjectController{})
	beego.Router("/team/ajax/delete", &projects.AjaxDeleteTeamProjectController{})
	beego.Router("/team/add/:id", &projects.AddTeamProjectController{})

	//项目需求
	beego.Router("/project/need/:id", &projects.NeedsProjectController{})
	beego.Router("/need/edit/:id", &projects.EditNeedsProjectController{})
	beego.Router("/need/add/:id", &projects.AddNeedsProjectController{})
	beego.Router("/need/show/:id", &projects.ShowNeedsProjectController{})
	beego.Router("/need/ajax/status", &projects.AjaxStatusNeedProjectController{})

	beego.Router("/my/need", &projects.MyNeedProjectController{})

	//项目任务
	beego.Router("/project/task/:id", &projects.TaskProjectController{})
	beego.Router("/task/edit/:id", &projects.EditTaskProjectController{})
	beego.Router("/task/add/:id", &projects.AddTaskProjectController{})
	beego.Router("/task/ajax/accept", &projects.AjaxAcceptTaskController{})
	beego.Router("/task/ajax/status", &projects.AjaxStatusTaskController{})
	beego.Router("/task/ajax/delete", &projects.DeleteTaskProjectController{})
	beego.Router("/task/show/:id", &projects.ShowTaskProjectController{})
	beego.Router("/project/taskbatch/:id", &projects.TaskBatchProjectController{})
	beego.Router("/task/clone/:id", &projects.CloneTaskProjectController{})

	beego.Router("/my/task", &projects.MyTaskProjectController{})

	//项目测试Bug
	beego.Router("/project/test/:id", &projects.TestProjectController{})
	beego.Router("/test/edit/:id", &projects.EditTestProjectController{})
	beego.Router("/test/add/:id", &projects.AddTestProjectController{})
	beego.Router("/test/ajax/accept", &projects.AjaxAcceptTestController{})
	beego.Router("/test/ajax/status", &projects.AjaxStatusTestController{})
	beego.Router("/test/ajax/delete", &projects.DeleteTestProjectController{})
	beego.Router("/test/show/:id", &projects.ShowTestProjectController{})

	beego.Router("/my/test", &projects.MyTestProjectController{})

	//项目文档
	beego.Router("/project/doc/:id", &projects.DocProjectController{})
	beego.Router("/doc/ajax/delete", &projects.AjaxDeleteDocPorjectController{})
	beego.Router("/doc/add/:id", &projects.FormDocProjectController{})
	beego.Router("/doc/edit/:id", &projects.FormDocProjectController{})
	beego.Router("/doc/show/:id", &projects.ShowDocProjectController{})

	//项目版本
	beego.Router("/project/version/:id", &projects.VersionProjectController{})
	beego.Router("/version/ajax/delete", &projects.AjaxDeleteVersionPorjectController{})
	beego.Router("/version/add/:id", &projects.FormVersionProjectController{})
	beego.Router("/version/edit/:id", &projects.FormVersionProjectController{})
	beego.Router("/version/show/:id", &projects.ShowVersionProjectController{})

	//知识分享
	beego.Router("/knowledge/manage", &knowledges.ManageKnowledgeController{})
	beego.Router("/knowledge/add", &knowledges.AddKnowledgeController{})
	beego.Router("/knowledge/edit/:id", &knowledges.EditKnowledgeController{})
	beego.Router("/knowledge/:id", &knowledges.ShowKnowledgeController{})
	beego.Router("/knowledge/comment/add", &knowledges.AddCommentController{})
	beego.Router("/knowledge/ajax/laud", &knowledges.AjaxLaudController{})
	beego.Router("/knowledge/ajax/delete", &knowledges.AjaxDeleteKnowledgeController{})

	//beego.Router("/task/ajax/status", &projects.AjaxAcceptTaskController{}, "*:AddPost")

	//相片
	beego.Router("/album/manage", &albums.ListAlbumController{})
	beego.Router("/album/upload", &albums.UploadAlbumController{})
	beego.Router("/album/edit", &albums.EditAlbumController{})
	beego.Router("/uploadmulti", &albums.UploadMultiController{})
	beego.Router("/album/:id", &albums.ShowAlbumController{})
	beego.Router("/album/comment/add", &albums.AddCommentController{})
	beego.Router("/album/ajax/laud", &albums.AjaxLaudController{})
	beego.Router("/album/ajax/delete", &albums.AjaxDeleteAlbumController{})

	//简历
	beego.Router("/resume/manage", &resumes.ManageResumeController{})
	beego.Router("/resume/add", &resumes.AddResumeController{})
	beego.Router("/resume/edit/:id", &resumes.EditResumeController{})
	beego.Router("/resume/ajax/status", &resumes.AjaxStatusResumeController{})
	beego.Router("/resume/ajax/delete", &resumes.AjaxDeleteResumeController{})

	beego.Router("/kindeditor/upload", &albums.UploadKindController{})

	beego.Router("/approval/manage", &leaves.ManageApprovalController{})
	//请假
	beego.Router("/leave/manage", &leaves.ManageLeaveController{})
	beego.Router("/leave/approval", &leaves.ApprovalLeaveController{})
	beego.Router("/leave/approval/:id", &leaves.ShowLeaveController{})
	beego.Router("/leave/edit/:id", &leaves.EditLeaveController{})
	beego.Router("/leave/add", &leaves.AddLeaveController{})
	beego.Router("/leave/ajax/status", &leaves.AjaxLeaveStatusController{})
	beego.Router("/leave/ajax/delete", &leaves.AjaxLeaveDeleteController{})

	//报销
	beego.Router("/expense/manage", &expenses.ManageExpenseController{})
	beego.Router("/expense/approval", &expenses.ApprovalExpenseController{})
	beego.Router("/expense/approval/:id", &expenses.ShowExpenseController{})
	beego.Router("/expense/edit/:id", &expenses.EditExpenseController{})
	beego.Router("/expense/add", &expenses.AddExpenseController{})
	beego.Router("/expense/ajax/status", &expenses.AjaxExpenseStatusController{})
	beego.Router("/expense/ajax/delete", &expenses.AjaxExpenseDeleteController{})

	//出差
	beego.Router("/businesstrip/manage", &businesstrips.ManageBusinesstripController{})
	beego.Router("/businesstrip/approval", &businesstrips.ApprovalBusinesstripController{})
	beego.Router("/businesstrip/approval/:id", &businesstrips.ShowBusinesstripController{})
	beego.Router("/businesstrip/edit/:id", &businesstrips.EditBusinesstripController{})
	beego.Router("/businesstrip/add", &businesstrips.AddBusinesstripController{})
	beego.Router("/businesstrip/ajax/status", &businesstrips.AjaxBusinesstripStatusController{})
	beego.Router("/businesstrip/ajax/delete", &businesstrips.AjaxBusinesstripDeleteController{})

	//外出
	beego.Router("/goout/manage", &goouts.ManageGooutController{})
	beego.Router("/goout/approval", &goouts.ApprovalGooutController{})
	beego.Router("/goout/approval/:id", &goouts.ShowGooutController{})
	beego.Router("/goout/edit/:id", &goouts.EditGooutController{})
	beego.Router("/goout/add", &goouts.AddGooutController{})
	beego.Router("/goout/ajax/status", &goouts.AjaxGooutStatusController{})
	beego.Router("/goout/ajax/delete", &goouts.AjaxGooutDeleteController{})

	//物品领用
	beego.Router("/oagood/manage", &oagoods.ManageOagoodController{})
	beego.Router("/oagood/approval", &oagoods.ApprovalOagoodController{})
	beego.Router("/oagood/approval/:id", &oagoods.ShowOagoodController{})
	beego.Router("/oagood/edit/:id", &oagoods.EditOagoodController{})
	beego.Router("/oagood/add", &oagoods.AddOagoodController{})
	beego.Router("/oagood/ajax/status", &oagoods.AjaxOagoodStatusController{})
	beego.Router("/oagood/ajax/delete", &oagoods.AjaxOagoodDeleteController{})

	//加班
	beego.Router("/overtime/manage", &overtimes.ManageOvertimeController{})
	beego.Router("/overtime/approval", &overtimes.ApprovalOvertimeController{})
	beego.Router("/overtime/approval/:id", &overtimes.ShowOvertimeController{})
	beego.Router("/overtime/edit/:id", &overtimes.EditOvertimeController{})
	beego.Router("/overtime/add", &overtimes.AddOvertimeController{})
	beego.Router("/overtime/ajax/status", &overtimes.AjaxOvertimeStatusController{})
	beego.Router("/overtime/ajax/delete", &overtimes.AjaxOvertimeDeleteController{})

	//考勤打卡
	beego.Router("/checkwork/manage", &checkworks.ManageCheckworkController{})
	beego.Router("/checkwork/all", &checkworks.ManageCheckworkAllController{})
	beego.Router("/checkwork/ajax/clock", &checkworks.AjaxClockUserController{})

	//消息
	beego.Router("/message/manage", &messages.ManageMessageController{})
	beego.Router("/message/ajax/delete", &messages.AjaxDeleteMessageController{})
	beego.Router("/message/ajax/status", &messages.AjaxStatusMessageController{})

	//组
	beego.Router("/group/manage", &groups.ManageGroupController{})
	beego.Router("/group/ajax/delete", &groups.AjaxDeleteGroupController{})
	beego.Router("/group/add", &groups.FormGroupController{})
	beego.Router("/group/edit/:id", &groups.FormGroupController{})
	//组成员
	beego.Router("/group/user/:id", &groups.ManageGroupUserController{})
	beego.Router("/group/user/add/:id", &groups.FormGroupUserController{})
	beego.Router("/group/user/ajax/delete", &groups.AjaxDeleteGroupUserController{})
	//组权限
	beego.Router("/group/permission/:id", &groups.ManageGroupPermissionController{})
	//beego.Router("/group/permission/add/:id", &groups.FormGroupPermissionController{})
	beego.Router("/group/permission/ajax/delete", &groups.AjaxDeleteGroupPermissionController{})

	//权限
	beego.Router("/permission/manage", &groups.ManagePermissionController{})
	beego.Router("/permission/ajax/delete", &groups.AjaxDeletePermissionController{})
	beego.Router("/permission/add", &groups.FormPermissionController{})
	beego.Router("/permission/edit/:id", &groups.FormPermissionController{})

	//区域操作
	beego.Router("/area/manage", &area.AreaMangerController{})
	beego.Router("/area/add/:id", &area.AreaAddController{})
	beego.Router("/area/edit/:id", &area.AreaEditController{})
	beego.Router("/area/delete/:id", &area.AreaDeleteController{})
	beego.Router("/area/getall/", &area.AllAreaController{})
	beego.Router("/area/getallarea/", &area.GetAllAreaController{})
	beego.Router("/ding", &area.DingController{})

	//项目mission
	beego.Router("/mission/manage", &mission.MissionManageController{})
	beego.Router("/mission/list", &mission.MissionListController{})
	beego.Router("/mission/add", &mission.MissionAddController{})
	beego.Router("/mission/detail/:id", &mission.MissionDetailController{})
	beego.Router("/mission/sendtask/:id", &mission.SendTaskController{})
	beego.Router("/mission/perfect/:id", &mission.MissionPerfectController{}) //todo
	beego.Router("/mission/sendtaskuni", &mission.SendTaskuniController{})
	//my mission
	beego.Router("/mymission/manage", &mission.MissionMyController{})
	beego.Router("/mymission/sub/:id", &mission.MissionMySubController{})
	beego.Router("/mobile/mymission/sub", &mission.MissionMySubMobileController{})
	beego.Router("/mymission/arraignment/:types", &mission.ArraignmentController{})
	beego.Router("/mymission/arraignmentsub/:id", &mission.ArraignmentSubController{})

	//档案资料
	beego.Router("/document/manage", &mission.DocumentManageController{})
	beego.Router("/document/add", &mission.DocumentAddController{})
	beego.Router("/document/detail", &mission.DocumentDetail{})
	beego.Router("/document/subtree", &mission.DocumentSubtree{})
	beego.Router("/documentmy/manage", &mission.DocumentSubview{})
	beego.Router("/document/audit", &mission.DocAudit{})
	beego.Router("/document/audit/sub", &mission.DocumentArraignmentController{})
	beego.Router("/document/sub", &mission.DocumentSub{})
	beego.Router("/docclass/manage", &mission.DocClassManageController{})
	beego.Router("/docclass/add", &mission.DocClassAddController{})
	beego.Router("/docclass/list", &mission.DocclassTree{})

	//实地考察
	beego.Router("/examination/manage", &mission.ExaminationManageController{})
	beego.Router("/examination/add", &mission.ExaminationAddController{})
	beego.Router("/examination/list", &mission.ExaminationList{})
	beego.Router("/exaclass/manage", &mission.ExaClassManageController{})
	beego.Router("/exaclass/add", &mission.ExaClassAddController{})
	beego.Router("/exaclass/list", &mission.ExaclassTree{})

	//测评管理
	beego.Router("/program/manage", &program.ProgramController{})
	beego.Router("/program/add", &program.ProgramAddController{})
	beego.Router("/program/edit/:id", &program.ProgramEditController{})
	beego.Router("/program/index", &program.ProgramIndexController{})
	beego.Router("/api/answer/list", &program.AnwserList{})
	beego.Router("/program/subjectlist/:pid", &program.SubjectListController{})
	beego.Router("/addsubject", &program.SubjectAddController{})
	beego.Router("/subject/edit", &program.SubjectManageController{})
	beego.Router("/mobile/program/list", &program.ListController{})
	beego.Router("/mobile/program/share", &program.ShareController{})
	beego.Router("/mobile/api/answer", &program.AnswerController{})

	//文件管理
	beego.Router("/files/manage", &files.FilesManageController{})
	beego.Router("/files/delete", &files.DeleteController{})
	beego.Router("/mobile/files/delete", &files.DeleteMobileController{})
	//标签管理
	beego.Router("/tags", &tags.TagsAllController{})

	//通知管理 消息管理
	beego.Router("/news/manage", &news.NewsManageController{})
	beego.Router("/news/add", &news.NewsAddController{})
	beego.Router("/news/classic", &news.NewsClassicController{})
	beego.Router("/news/classicedit/:id", &news.NewsClassicEditController{})
	beego.Router("/news/classicadd", &news.NewsClassicAddController{})
	beego.Router("/news/edits/:id", &news.NewsEditsController{})
	beego.Router("/news/ajax/delete", &news.NewsAjaxDeleteController{})
	beego.Router("/news/ajax/classicdelete", &news.NewsAjaxClassicDeleteController{})
	// API接口---------------------
	//任务管理 单位列表
	beego.Router("/mobile/news/areaList", &news.ApiAreaController{})
	//任务管理 任务列表
	beego.Router("/mobile/news/missionList", &news.ApiMissionController{})
	//任务管理 任务详情
	beego.Router("/mobile/news/missionInfo", &news.ApiMissionInfoController{})

	//消息列表
	beego.Router("/mobile/news/newsList", &news.ApiNewsController{})
	//消息详情
	beego.Router("/mobile/news/newsinfo", &news.ApiNewsInfoController{})
	//消息类型
	beego.Router("/mobile/news/classicList", &news.ApiNewsClassicController{})
	//我的任务
	beego.Router("/mobile/mission/missionMy", &mission.ApiMissionMyController{})
	//任务详情
	beego.Router("/mobile/mission/missionInfo", &mission.ApiMissionInfoController{})
	//点位信息
	beego.Router("/mobile/mission/areaInfo", &mission.ApiAreaInfoController{})
	beego.Router("/mobile/user/profile", &users.UserProfileController{}) //手机用户信息完善

	//成功率
	beego.Router("/mobile/area/doneRate", &area.ApiDoneRateController{})
	beego.Router("/mobile/checkworks", &checkworks.MobileClockUserController{})
	beego.Router("/mobile/user/info", &users.UserInfo{})
	//点位打卡
	beego.Router("/area/daka", &area.DakaController{})
	//达标率
	beego.Router("/area/passRate", &area.ApiPassRateController{})
	//问卷调查
	beego.Router("/area/question", &area.ApiQuestionController{})

}
