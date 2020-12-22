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
      <form class="searchform" action="/project/manage" method="get">
        <select name="status" class="form-control">
          <option value="">测评状态</option>
          <option value="0" {{if eq "0" .condArr.status}}selected{{end}}>未发布</option>
          <option value="1" {{if eq "1" .condArr.status}}selected{{end}}>已发布</option>
        </select>
        <input type="text" class="form-control" name="keywords" placeholder="请输入名称" value="{{.condArr.keywords}}"/>
        <button type="submit" class="btn btn-primary">搜索</button>
      </form>
      <!--search end-->
      {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h3> 测评管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/">首页</a> </li>
        <li> <a href="/project/manage">测评管理</a> </li>
        <li class="active"> 测评管理 </li>
      </ul>
      <div class="pull-right"><a href="/program/add" class="btn btn-success">+新建测评</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 测评 / 总数：{{.countProject}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span> </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="project-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>测评名称</th>
						<th>答题人数</th>
                        <th>创建时间</th>
                        <th>状态</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                    
                    {{range $k,$v := .program}}
                    <tr>
                      <td><a href="/subject/edit?pid={{$v.Id}}">{{$v.Title}}</a></td>
                      <td>{{$v.Counts}}</td>
					  <td>{{getDate $v.Creatime}}</td>
                      <td>{{if eq 1 $v.Pstatus}}发布{{else if eq 0 $v.Pstatus}}未发布{{end}}</td>
                      <td><div class="btn-group">
                          <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                          <ul class="dropdown-menu">
                            <li><a href="/program/edit/{{$v.Id}}">编辑</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="/subject/edit?pid={{$v.Id}}" data-status="1">查看题目</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="javascript:;" class="js-project-single" data-id="{{$v.Id}}" data-status="3">进行</a></li>
                            <li><a href="javascript:;" class="js-project-single" data-id="{{$v.Id}}" data-status="4">结束</a></li>
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
