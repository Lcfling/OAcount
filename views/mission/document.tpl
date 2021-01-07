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
      <h3> 档案管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/">首页</a> </li>
        <li> <a href="/project/manage">档案管理</a> </li>
        <li class="active"> 档案 </li>
      </ul>
      <div class="pull-right"><a href="/document/add?pid={{.parentid}}" class="btn btn-success">+新内容</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 项目 / 总数：{{.countMission}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span> </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="project-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                    <tr>
                      <th>所属类目</th>
                      <th>名称</th>
                      <th>内容</th>
                      <th>要求</th>
                      <th width="200">操作</th>
                    </tr>
                    </thead>
                    <tbody>

                    {{range $k,$v := .doc}}
                    <tr>
                      <td><a href="/mission/detail/{{$v.Id}}">{{GetDocclassName $v.Pid}}</a></td>

                      <td><a href="">{{$v.Name}}</a></td>
                      <td>{{$v.Content}}</td>
                      <td>{{$v.Need}}</td>
                      <td><div class="btn-group">
                        <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                        <ul class="dropdown-menu">
                          <li><a href="/mission/edit/{{$v.Id}}">编辑</a></li>
                          <li role="separator" class="divider"></li>
                          <li><a href="/mission/sendtaskuni?types=0&id={{$v.Id}}">任务下发</a></li>

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
