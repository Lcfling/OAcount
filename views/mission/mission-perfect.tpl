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
        <li> <a href="/mission/manage">项目管理</a> </li>
        <li class="active"> 项目 </li>
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
              <form class="form-horizontal adminex-form" id="mission-perfect">
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>项目名称</label>
                  <div class="col-sm-10">
                    <input type="text"  disabled="disabled" name="name" value="{{.mission.Name}}" class="form-control" placeholder="请填写名称">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">任务类型</label>
                  <div class="col-sm-10">
                    <select name="types" class="form-control" id="types" disabled="disabled">
                      <option value="0" {{if eq 0 .mission.Types}}selected{{end}}>档案资料</option>
                      <option value="1" {{if eq 1 .mission.Types}}selected{{end}}>实地考察</option>
                      <option value="2" {{if eq 2 .mission.Types}}selected{{end}}>测评调查</option>
                      <option value="3" {{if eq 3 .mission.Types}}selected{{end}}>岗位打卡</option>

                    </select>
                  </div>
                </div>



      <div class="form-group">
        <label class="col-sm-2 col-sm-2 control-label">附件</label>
        <div class="col-sm-10">
          {{range $k,$v:=.files}}

          {{if eq $v.Types 1}}
          <p><a target="_blank" href="{{$v.Path}}"><img width="225" src="{{$v.Path}}"/></a> <span data-id="{{$v.Id}}" class="delete-file fa-2x"> <i style="color: red" class="fa fa-trash-o"></i> </span></p>
          {{else}}
          <p><a target="_blank" href="{{$v.Path}}">{{$v.Oldname}}</a><span data-id="{{$v.Id}}" class="delete-file fa-2x"> <i style="color: red" class="fa fa-trash-o"></i> </span></p>
          {{end}}
          {{end}}


          <p><input type="file" name="attachment" multiple="multiple">
            <span> - </span><br/></p>
          <p><input type="file" name="attachment" multiple="multiple">
            <span> - </span><br/></p>
          <p><input type="file" name="attachment" multiple="multiple">
            <span> - </span><br/></p>
        </div>

<!--                  <label class="col-sm-2 col-sm-2 control-label">开始和结束日期</label>-->
<!--                  <div class="col-sm-10">-->
<!--                    <div class="input-group input-large custom-date-range" data-date="2016-07-07" data-date-format="yyyy-mm-dd">-->
<!--                      <input type="text" class="form-control dpd1" name="started" placeholder="开始日期" value="{{getDate .mission.Started}}">-->
<!--                      <span class="input-group-addon">To</span>-->
<!--                      <input type="text" class="form-control dpd2" name="ended"  placeholder="结束日期" value="{{getDate .mission.Ended}}">-->
<!--                    </div>-->
<!--                  </div>-->
<!--                </div>-->
<!--                <div class="form-group" id="pid" style="display: none">-->
<!--                  <label class="col-sm-2 col-sm-2 control-label">任务类型</label>-->
<!--                  <div class="col-sm-10">-->
<!--                    {{$checkp:= .mission.Mid}}-->
<!--                    <select name="types" class="form-control">-->

<!--                      {{range $k,$v:=.program}}-->
<!--                      <option value="{{$v.Id}}" {{if eq $checkp $v.Id}}selected{{end}}>{{$v.Title}}</option>-->
<!--                      {{end}}-->
<!--                    </select>-->
<!--                  </div>-->
<!--                </div>-->
<!--                <div class="form-group">-->
<!--                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>描述</label>-->
<!--                  <div class="col-sm-10">-->
<!--                    <textarea name="desc" placeholder="请填写描述" style="height:300px;" class="form-control">{{.mission.Desc}}</textarea>-->
<!--                  </div>-->
<!--                </div>-->

                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" name="id" value="{{.mission.Id}}">
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
