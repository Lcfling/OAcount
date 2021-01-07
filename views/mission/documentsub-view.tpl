<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/js/bootstrap-datepicker/css/datepicker-custom.css" rel="stylesheet" />
  <style>
    .menu{
      height: 100%;
      overflow-y: scroll;
    }
    .menu >li.active{
      /*border: 1px solid #5AFCFB;*/
    }
    .menu >li:before{
      content: '';
      position: absolute;
      /*background: ;*/
      width: 15px;
      height: 15px;
      left: 0;
      top: 0;
      bottom: 0;
      margin: auto;
    }
    .menu >li ul{
      position: absolute;
      width: 100%;
      left: 0;
      top: 100%;
    }
    .menu >li li.active{
      background: #F1F1F1;
    }
    .menu .search{
      width: 120px;
      margin: 10px;
      border: 1px solid #5C5C5C;
      border-radius: 3px;
      padding: 5px;
      display: -webkit-flex;
      align-items: center;
      position: relative;
    }
    .menu .search div{
      flex:1;
      height: 1.5em;
    }
    .menu .search ul{
      position: absolute;
      width: 100%;
      left: 0;
      top: 100%;
      background: #f4f5fa;
      z-index: 100;
      padding: 5px;
      display: none;
      list-style: none;

    }
    .menulist li{
      position: relative;
      line-height: 1.8;
      text-overflow: ellipsis;
      overflow: hidden;
      white-space: nowrap;
      list-style: none;
    }
    .menulist p:before{
      content: '';
      position: absolute;
      background: url("../static/img/15.png")no-repeat 10px 7px;
      background-size: 45% 45%;
      width: 20px;
      height: 20px;
      left: 0;
    }
    .menulist li p{
      padding-left: 20px;
    }
    .menulist ul{
      display: none;
    }
    .menulist ul li{
      padding-left: 15px;
    }
    .menulist ul li.active{
    }
    .menulist p:hover{
    }
  </style>
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
            <div class="panel-body" >
              <div class="col-lg-4">
                <article class="menu">
                  <div class="search">
                    <div><span>请选择</span><ul></ul></div>
                    <img src="img/a.png" alt="" width="10">
                  </div>
                  <ul class="menulist">

                  </ul>

                </article>
              </div>
              <div class="col-lg-8" id="subform"></div>
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

  let list,detail;
  $.ajax({
    type:"GET",
    url: '/document/subtree',
    dataType: 'json',
    data: {
      pid: 0
    },
    success: function (res) {
      console.log(res)
      let tItem='',id;
      list = res.data? res.data: '';

      if(list.length){
        for (let i=0;i<list.length;i++){
          // recursive($('.menulist'), list[i])
          tItem += '<li data-id="'+list[i].Id+'">'+list[i].Title+'</li>'
        }
        $('.search ul').html(tItem);
        // $('.menulist').html(item);
        // selectli(id,$('#newsbox'));
      }else{
        // $('.menulist').html(item);
      }



    }
  });

  //递归函数
  function recursive(el, obj) {
    // console.log(el, obj)
    // console.log(obj.Child)
    if (obj.Child == null) {
      //最终节点
      var li = $("<li>");
      var p = $("<p>").text(obj.Title).attr("postId", obj.Id).attr("Pid", obj.Pid).attr("postShow",obj.Show).attr("postUp",0);
      li.append(p);
      el.append(li);
    } else {
      var li = $("<li>");
      var p = $("<p>").text(obj.Title).attr("postId", obj.Id).attr("postShow",obj.Show).attr("postUp",0);
      li.append(p);
      el.append(li);
      var ul = $("<ul>");
      p.after(ul);
      for (var i=0;i<obj.Child.length;i++) {
        recursive(ul, obj.Child[i]);
      }

    }
  }

  $(".search div").click(function () {
    $(".search ul").slideToggle();
  });
  $(document).on('click','.search li',function () {
    let index = $(this).index();
    if(list[index].Child){
      console.log(list[index])
      $(".search div span").html(list[index].Title)
      $('.menulist').html('')
      for (var i=0;i<list[index].Child.length;i++) {
        recursive($('.menulist'), list[index].Child[i]);
      }
    }else{
      $('.menulist').html('')
    }
  });
  $(document).on('click','.menulist p',function () {
    let id = $(this).attr('postId');
    let show = $(this).attr('postShow');
    let up = $(this).attr('postUp');

    let pid=$(this).attr('Pid');
    if(show == 1){
      //selectli(id,$('#newsbox ul'))
      getsubform(id,pid)
    }
    up==0?up=1:up=0;
    $(this).attr('postUp',up);
    if(up==1){
      $(this).siblings('ul').show();

      var beforeStyle = window.getComputedStyle(this, ":before");
      beforeStyle.background = url('img/16.png')

    }else{
      $(this).siblings('ul').hide().parent().find('ul').hide();
      $(this).siblings('ul').hide().parent().find('p').attr('postUp',0);
    }
  });

  function selectli(aid,obj) {
    new AjaxRequest({
      url: hostUrl +'document/detail',
      data: {
        // lastid: 0,
        id: aid,
      },
      callback: function (res) {
        let item ='';
        detail = res.data? res.data: '';
        if(detail){

          //let url = 'https://view.officeapps.live.com/op/view.aspx?src=' + encodeURI(fileUrl + newsList[0].Path)
          // obj.find('iframe').attr('src',url);

          obj.find('.name').html(detail.Name);
          obj.find('.content').html(detail.Content);

        }else{
          // obj.find('iframe').attr('src','');

          obj.find('.name').html('标题');
          obj.find('.content').html('');
        }
      }
    })
  }



  function getsubform(id,pid) {
    var data={}
    data.pid=2
    data.id=270

    $.ajax({
      type: 'GET',
      url: '/document/sub',
      data: data,
      dataType: 'html',
      success: function (res) {
        console.log(res);
        $("#subform").html(res)
      }
    })
  }

})
</script>
</body>
</html>
