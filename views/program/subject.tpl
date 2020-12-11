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
      <h3> 测评管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/">首页</a> </li>
        <li> <a href="/project/manage">测评管理</a> </li>
        <li class="active"> {{.program.Title}} </li>
      </ul>

    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 答题总人数 ：{{.program.Counts}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span> </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="project-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>测评题目</th>
						<th>类型</th>
                        <th>选项</th>

                      </tr>
                    </thead>
                    <tbody>

                    {{$q:=.program.Counts}}

                    {{range $k,$v := .subject}}
                    <tr>
                      <td>{{$v.Content.Title}}</td>
                      <td>{{if eq "1" $v.Content.Type}}单选{{else if eq "2" $v.Content.Type}}多选{{else}}未知{{end}}</td>
					  <td>
                        {{range $s,$a := $v.Content.Subject}}
                        <p>{{$a.Content}} <span>人数：{{$a.Count}} 比例：%{{Ratio $a.Count $q}} </span></p>
                        {{end}}
                      </td>
                    </tr>
                    {{end}}
                    </tbody>
                    
                  </table>
                </form>
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
