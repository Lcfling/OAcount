<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/css/table-responsive.css" rel="stylesheet">
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      <!--search start-->

      <!--search end-->
      {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h3> 我的任务 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">首页</a> </li>
        <li> <a href="/project/manage">我的任务</a> </li>
        <li class="active"> 列表 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 任务 / 总数：{{.countmissionmy}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span> </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="project-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>名称</th>
                        <th>区域</th>
                        <th>负责人</th>
                        <th>是否响应</th>
                        <th>状态</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>

                    {{range $k,$v := .missionmy}}
                    <tr>
                      <td><a href="/mission/detail/{{$v.Id}}">{{$v.Name}}</a></td>
                      <td>{{$v.Areaname}}</td>
                      <td><a data-id="{{$v.Userid}}">{{getRealname $v.Userid}}</a></td>
                      <td><a href="">{{if eq 0 $v.Check}}未响应{{else if eq 2 $v.Check}}已响应{{end}}</a></td>

                      <td>{{if eq 1 $v.Status}}进行中{{else if eq 2 $v.Status}}待反馈{{else if eq 3 $v.Status}}待解决{{else if eq 4 $v.Status}}已完成{{end}}</td>
                      <td><div class="btn-group">
                          <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                          <ul class="dropdown-menu">
                            <li><a href="/project/edit/{{$v.Id}}">编辑</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="javascript:;" class="js-mission-single" data-id="{{$v.Id}}" data-status="4">查看</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="/mymission/sub/{{$v.Id}}">提审</a></li>
                          </ul>
                        </div></td>
                    </tr>
                    {{end}}
                    </tbody>
                    
                  </table>
                </form>
                {{template "inc/page.tpl" .}}
				 </section>
            </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}

</body>
</html>
