<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head>
<body class="sticky-header">
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
      <h3> 档案分类
        <span>
          {{if eq .status 1}}待审核{{else}}<a href="/document/audit?status=1">待审核</a>{{end}}
          {{if eq .status 2}}驳回{{else}}<a href="/document/audit?status=2">驳回</a>{{end}}
          {{if eq .status 3}}初验合格{{else}}<a href="/document/audit?status=3">初验合格</a>{{end}}
          {{if eq .status 4}}合格{{else}}<a href="/document/audit?status=4">合格</a>{{end}}
        </span>
      </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/">首页</a> </li>
        <li> <a href="/permission/manage">档案分类</a> </li>
        <li class="active"> {{if eq .status 1}}待审核{{end}}
          {{if eq .status 2}}驳回{{else}}{{end}}
          {{if eq .status 3}}初验合格{{else}}{{end}}
          {{if eq .status 4}}合格{{else}} {{end}}</li>
      </ul>
      <div class="pull-right"> <a href="/docclass/add?pid={{.parentid}}" class="btn btn-success">+新增分类</a> </div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 权限管理 / 总数：{{.countArea}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a> </span> </header>
            <div class="panel-body">
              <table class="table table-bordered table-striped table-condensed">
                <thead>
                  <tr>
					<th>档案名称</th>
                    <th>区域</th>
                    <th class="hidden-phone hidden-xs">负责人</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody>
                
                {{range $k,$v := .missions}}
                <tr>
				  <td>{{$v.Name}}</td>
                  <td><a href="/docclass/manage?pid={{$v.Id}}">{{$v.Areaname}}</a> </td>
                  <td class="hidden-phone hidden-xs">{{getRealname $v.Userid}}</td>
                  <td><div class="btn-group">
                      <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                      <ul class="dropdown-menu">
                        <li><a href="/document/audit/sub?id={{$v.Id}}">审核</a></li>
                      </ul>
                    </div></td>
                </tr>
                {{else}}
                <tr>
                  <td colspan="7">你还没有添加响应区域</td>
                </tr>
                {{end}}
                </tbody>
                
              </table>
              {{template "inc/page.tpl" .}} </div>
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