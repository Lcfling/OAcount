<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
  <style type="text/css">
    #allmap {width: 80%;height: 80%;overflow: hidden;margin:0;font-family:"微软雅黑";}
  </style>
</head><body class="sticky-header">
<style type="text/css">#allmap {overflow: hidden;margin:0;font-family:"微软雅黑";}</style>
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
      <h3> 添加、编辑区域 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">OPMS</a> </li>
        <li> <a href="/department/manage">区域管理</a> </li>
        <li class="active"> 区域 </li>
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
              <form class="form-horizontal adminex-form" id="area-form">
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>名称</label>
                  <div class="col-sm-10">
                    <input type="text" name="name" value="{{.area.Name}}" class="form-control" placeholder="请填写名称">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">负责人</label>
                  <div class="col-sm-10">
                    <select name="owner" class="form-control">
                      <option value="0" {{if eq 0 $.area.Owner}}selected{{end}}>无</option>
                      {{range .myuser}}
                      <option value="{{.Userid}}" {{if eq .Userid $.area.Owner}}selected{{end}}>{{.Realname}}</option>
                      {{end}}
                    </select>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">是否局委</label>
                  <div class="col-sm-10">
                    <select name="jstatus" class="form-control">
                      <option value="0" {{if eq 0 .area.Jstatus}}selected{{end}}>区域</option>
                    <option value="1" {{if eq 1 .area.Jstatus}}selected{{end}}>局委</option>

                </select>
              </div>
            </div>

            <div class="form-group">
              <label class="col-sm-2 col-sm-2 control-label"><span>*</span>标签</label>
              <div class="col-sm-10">
                <input type="text" name="tags" value="{{.area.Tags}}" class="form-control" placeholder="标签">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-2 col-sm-2 control-label"><span>*</span>坐标信息</label>
              <div class="col-sm-10">
                <input type="text" id="location" name="locations" value="{{.area.Locations}}" class="form-control" placeholder="请填写坐标信息">
              </div>
            </div>
                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" name="id" value="{{.area.Id}}">
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
<div aria-hidden="true" aria-labelledby="projectModalLabel" role="dialog" tabindex="-1" id="projectModal" class="modal fade">
<div class="modal-dialog">
  <div class="modal-content">
    <div class="modal-header">
      <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
      <h4 class="modal-title">请选择坐标</h4>
    </div>
    <div class="modal-body" style="height:600px">
      <div id="allmap">
      </div>


    </div>

    <div class="modal-footer">
      <a href="/project/manage" class="btn btn-primary">去设置管理</a>
    </div>
  </div>
</div>
</div>
<script type="text/javascript" src="http://api.map.baidu.com/api?v=2.0&ak=Z1bz6BGmrthCGyz95vQEpITcy7VSztDx"></script>
<script>


var markergg;
var map = new BMap.Map("allmap");
var point = new BMap.Point(116.331398,39.897445);//地图初始位置
map.centerAndZoom(point,12);//默认地图级别12级
map.setDefaultCursor("url('bird.cur')");   //设置地图默认的鼠标指针样式
function myFun(result){
  var cityName = result.name;
  map.setCenter(cityName);
}
var myCity = new BMap.LocalCity(); //根据ip设置坐标初始位置
myCity.get(myFun);

//var map = new BMap.Map("allmap");
//map.centerAndZoom("重庆",12);
map.enableScrollWheelZoom();   //启用滚轮放大缩小，默认禁用
map.enableContinuousZoom();    //启用地图惯性拖拽，默认禁用
//单击获取点击的经纬度
map.addEventListener("click",function(e){

  map.clearOverlays();
  var x = e.point.lng;
  var y = e.point.lat;
  $("#location").val(x+","+y)
  var ggPoint = new BMap.Point(x,y);
  map.addControl(new BMap.NavigationControl());
  var markergg = new BMap.Marker(ggPoint);
  map.addOverlay(markergg); //添加GPS坐标
  //var labelgg = new BMap.Label("",{offset:new BMap.Size(2,2)});
  //markergg.setLabel(labelgg); //添加GPS label
  map.setDefaultCursor("url('bird.cur')");     //设置鼠标形状为手型
});

$(function(){
  $("#location").on("click",function () {
    //alert(1)
    $('#projectModal').modal('toggle').find('.modal-body').html();
  })
})

//$('#projectModal').modal('toggle').find('.modal-body').html(html);
</script>
</body>
</html>
