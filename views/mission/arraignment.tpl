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
      <h3> 我的任务 <span><a href="/mymission/arraignment/0">未审核</a><a href="/mymission/arraignment/2">已经审核</a><a href="/mymission/arraignment/1">驳回</a></span></h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">首页</a> </li>
        <li> <a href="/project/manage">任务管理</a> </li>
        <li class="active"> 任务审核 </li>
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
                        <th>完成状态</th>
                        <th>审核状态</th>
                        <th>反馈信息</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>

                    {{range $k,$v := .mymissions}}
                    <tr>
                      <td>{{$v.Name}}</td>
                      <td>{{$v.Areaname}}</td>
                      <td><a data-id="{{$v.Userid}}">{{getRealname $v.Userid}}</a></td>
                      <td><a href="">{{if eq 0 $v.Status}}未完成{{else if eq 1 $v.Status}}已完成{{end}}</a></td>

                      <td>{{if eq 0 $v.Arraignment}}未审核{{else if eq 1 $v.Arraignment}}驳回{{else if eq 2 $v.Arraignment}}审核通过{{end}}</td>
                      <td>{{$v.Feedback}}</td>
                      <td><div class="btn-group">
                          <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                          <ul class="dropdown-menu">
                            <li><a href="/project/edit/{{$v.Id}}">编辑</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="/mymission/arraignmentsub/{{$v.Id}}">审核</a></li>
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
