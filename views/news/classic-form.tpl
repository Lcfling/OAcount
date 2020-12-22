<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/js/bootstrap-datepicker/css/datepicker-custom.css" rel="stylesheet" />
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a> {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h3> 项目管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">OPMS</a> </li>
        <li> <a href="/news/manage">通知管理</a> </li>
        <li class="active"> +新消息 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="classic-form">
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>消息类型</label>
                  <div class="col-sm-10">
                    <input type="text" maxlength="1" name="classname" value=""  placeholder="请填写消息类型">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" name="id" value="">
                    <button type="submit" class="btn btn-primary">提 交</button>
                  </div>
                </div>
              </form>
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
<div aria-hidden="true" aria-labelledby="projectModalLabel" role="dialog" tabindex="-1" id="projectModal" class="modal fade">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
            <h4 class="modal-title">新建项目成功，请先按项目流程设置</h4>
          </div>
          <div class="modal-body">



          </div>
          <div class="modal-footer">
            <a href="/mission/manage" class="btn btn-primary">去设置管理</a>
          </div>
        </div>
      </div>
    </div>
{{template "inc/foot.tpl" .}}
<script src="/static/js/bootstrap-datepicker/js/bootstrap-datepicker.js"></script>
<script src="/static/keditor/kindeditor-min.js"></script>
<script>
$(function(){
    $("#types").on("input",function () {
      //alert("dddd")
      console.log($(this))
      var val=$(this).val()
      if (val==2){
        $("#pid").show()
      } else{
        $("#pid").hide()
      }
    })
	var editor = KindEditor.create('textarea[name="desc"]', {
	    uploadJson: "/kindeditor/upload",
	    allowFileManager: true,
	    filterMode : false,
	    afterBlur: function(){this.sync();}
	});

	var nowTemp = new Date();
    var now = new Date(nowTemp.getFullYear(), nowTemp.getMonth(), nowTemp.getDate(), 0, 0, 0, 0);

    var checkin = $('.dpd1').datepicker({
		 format: 'yyyy-mm-dd',
        onRender: function(date) {
            return date.valueOf() < now.valueOf() ? 'disabled' : '';
        }
    }).on('changeDate', function(ev) {
            if (ev.date.valueOf() > checkout.date.valueOf()) {
                var newDate = new Date(ev.date)
                newDate.setDate(newDate.getDate() + 1);
                checkout.setValue(newDate);
            }
            checkin.hide();
            $('.dpd2')[0].focus();
        }).data('datepicker');
    var checkout = $('.dpd2').datepicker({
		 format: 'yyyy-mm-dd',
        onRender: function(date) {
            return date.valueOf() <= checkin.date.valueOf() ? 'disabled' : '';
        }
    }).on('changeDate', function(ev) {
            checkout.hide();
        }).data('datepicker');
})
</script>
</body>
</html>
