<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/css/table-responsive.css" rel="stylesheet">
<style>
.border { border-bottom: 1px solid #ddd;margin-bottom:20px;}
.pul{}
.pul li {display: block;float: left; width: 100%;}
.cul {display:block}
.cul li { float: left; list-unstyled: none; width: 130px;}
</style>
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
      <h3> 组织管理 {{template "users/nav.tpl" .}}</h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">OPMS</a> </li>
        <li> <a href="/user/manage">用户管理</a> </li>
        <li class="active"> 权限 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 任务下发 <span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span> </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="sendtask-form">
                  <ul class="list-unstyled pul" id="js-permission">
					<li class="text-center">
                      <input type="hidden" id="groupid" value="{{.group.Id}}">
                      <button type="button" id="permission-btn-new" class="btn btn-success">任务下发</button>
                    </li>
                  </ul>
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
<script>
$(function(){
	$('body').delegate('input[name="checkpermission"]', 'click',function(){
		var obj = $(this);
		obj.parent().next().find('input').each(function(){
			if (obj.is(':checked')) {
				$(this).prop('checked', true)
			} else {
				$(this).prop('checked', false)
			}		
		});
	});


	var pstring = "{{.area}}";
	var farr = pstring.substring(0,pstring.length-1).split(',');
	
	var cstring = "{{.cstring}}";
	var carr = cstring.substring(0,cstring.length-1).split(',');
	
	var html = '';
	var lefthtml = '';

	var jsonstr='{{.area}}'

    var jsonOb=JSON.parse(jsonstr);
	console.log(jsonOb)
	
	for(var i=0;i<jsonOb.length;i++) {
		html +='<li data-id="'+jsonOb[i].Id+'" class="border">';
		html += '<h4>'+jsonOb[i].Name+' <input type="checkbox" name="checkarea" ></h4>';
		
		html+='<ul class="cul">'
        if (jsonOb[i].Child!=null) {
          for (var j = 0; j < jsonOb[i].Child.length; j++) {
            var arr2=jsonOb[i].Child
            html += '<li>';
            html += '<div class="form-group" data-id="' +arr2[j].Name+ '"> ';
            html += '<label class="checkbox-inline">';
            html += '<input type="checkbox" name="checkareas[]" data-ename="' + arr2[j].Id + '">';
            html += arr2[j].Name;
            html += '</label>';
            html += '</div>';
            html += '</li>';

            if (arr2[j].Child!=null){

                for (var k = 0; k < arr2[j].Child.length; k++) {
                  var arr3=arr2[j].Child
                  html += '<li>';
                  html += '<div class="form-group" data-id="' +arr3[k].Name+ '"> ';
                  html += '<label class="checkbox-inline">';
                  html += '<input type="checkbox" name="checkareas[]" data-ename="' + arr3[k].Id + '">';
                  html += arr3[k].Name;
                  html += '</label>';
                  html += '</div>';
                  html += '</li>';

                }

            }
          }
        }

		html += '</ul>';			
		html +='</li>';
	}

	$('#js-permission').prepend(html);
	//$('.js-left-nav').append(lefthtml);
	
	var per = '{{.groupspermissions}}';
	var val = '';
	$('input[name="permission[]"]').each(function(){
		val =  $(this).val();
		//console.log(val)
		if (per.indexOf(val) > -1) {
			$(this).attr('checked', true)
		}
	});
	
	var leftnav = '';
	var lefthtml = '';
})
</script>
</body>
</html>