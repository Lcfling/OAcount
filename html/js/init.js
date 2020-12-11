/*全局参数定义*/
// const hostUrl = "http://114.67.97.70:6756/api/";
const hostUrl = "http://wenjuan.webziti.com/api/";
const hostUrl2 = "http://wenjuan.webziti.com/";

// document.getElementsByTagName('html')[0].style.fontSize = document.body.clientWidth / 7.5 + 'px';
document.getElementsByTagName('html')[0].style.fontSize = document.documentElement.clientWidth / 7.5 + 'px';
window.addEventListener('resize' , function () {
    document.getElementsByTagName('html')[0].style.fontSize = document.documentElement.clientWidth / 7.5 + 'px';

});

$(function () {
    // JQ Alert
    var toast = '<div class="toast" ></div>';
    var loading = '<div class="loading" ></div>';
    var dialog = '<div class="mask close" ></div><div class="mask-box dialog border-radius-1 center"><p>温馨提示</p><div class="mask-msg pad-3"></div><div class="pad-t-3 flex space-ard"><span class="dia_cancel bg-gray pad-2-5 border-radius-1" >取消</span><span class="dia_confirm bg-blue color-white pad-2-5 border-radius-1">确定</span></div></div>';
    $("body").append(toast, loading, dialog);

    // 判断login
    var L = getUrlname(window.location.href);
    if(!(localStorage.user_id && localStorage.userpwd)){
        if(!(L == 'login.html'|| L == 'register.html')){
            // window.location.href = 'login.html';
        }
    }

});


// Ajax Request
(function () {
    function AjaxRequest(opts) {
        this.type = opts.type || 'post';
        this.url  = opts.url;
        this.dataType = opts.dataType || 'json';
        this.data= opts.data || {"userid": localStorage.user_id,"token": localStorage.token};
        this.isShowLoader = opts.isShowLoader || false;
        this.callback = opts.callback;
        this.error = opts.error;
        this.init();
    }
    AjaxRequest.prototype = {
        // 初始化
        init: function () {
            this.sendRequest();
        },
        //渲染loader
        showLoader: function () {
            if(this.isShowLoader){
                var loading = '<div class="loading none"></div>';
                $("body").append(loading);
                $(".loading").show();
            }
        },
        //隐藏loader
        hideLoader: function () {
            if(this.isShowLoader){
                $(".loading").hide();
            }
        },
        // 发送请求
        sendRequest: function () {
            var self = this;
            $.ajax({
                type: this.type,
                data: this.data,
                dateType: this.dataType,
                url: this.url,
                beforeSend: res => {
                    self.showLoader();
                    res.setRequestHeader("userid", localStorage.user_id);
                    res.setRequestHeader("token", localStorage.token);
                    // console.log("userid:"+localStorage.user_id+",token:"+localStorage.token);
                    // res.setRequestHeader('X-CSRF-TOKEN',$('#token').val())
                },
                success: res => {
                    self.hideLoader();
                    console.log(res);
                    if(res != null && res != ""){
                        // 登录状态失败
                        if(res.status == 10001){
                            localStorage.user_id = "";
                            localStorage.token = "";
                            window.location.href = 'login.html';
                        }
                        if(self.callback){
                            if (Object.prototype.toString.call(self.callback) === "[object Function]") {   //Object.prototype.toString.call方法--精确判断对象的类型
                                self.callback(res);
                            }else{
                                console.log(self.callback+"这也不是个函数!");
                            }
                        }
                    }
                },
                error: err => {
                    self.hideLoader();
                    toast('请求错误');
                    console.log(err);
                    if(self.error){
                        if (Object.prototype.toString.call(self.error) === "[object Function]") {   //Object.prototype.toString.call方法--精确判断对象的类型
                            self.error(err);
                        }else{
                            console.log(self.error+"这不是个函数!");
                        }
                    }
                }
            });
        },
    };
    window.AjaxRequest = AjaxRequest;
})();

// JQ toast
function toast(txt) {
    $(".toast").text(txt).show().delay(2000).hide(0);
}
// JQ dialog
function dialog(txt) {
    $(".mask-msg").text(txt);
    $(".mask").show();

    $(".mask-box").show();
}
// JQ dialog clickMask
function close() {
    $(".mask").hide();
    $(".mask-box").hide();
    $(".mask-bot").hide();
}
$(document).on("click",".close",function () {
    $(".mask").hide();
    $(".mask-box").hide();
    $(".mask-bot").hide();
});


//获取get变量
function getQueryString(name){
    var reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)', 'i');
    var r = window.location.search.substr(1).match(reg);
    if (r != null) {
        return unescape(r[2]);
    }
    return null;
}
function getUrlname(url){
    url=url.split('?')[0];
    var urlSlashCount=url.split('/').length;
    return url.split('/')[urlSlashCount-1].toLowerCase();
}
function getUrlParams(url) {
    var params;
    url.search ? params = url.search.substr(1) : params = '';
    params.indexOf('&') > -1 ? params = params.split('&') : '';  //返回数组
    return params;
}
function getUrlParam(url,param) {
    if(url.indexOf('?') == -1 || url.indexOf(param) == -1)
        return '';
    var arr = url.split('?')[1].split('&');
    var result = {};
    arr.forEach(item => {
        result[decodeURIComponent(item.split('=')[0])] = decodeURIComponent(item.split('=')[1]);
    });
    return result[param];
}
function parseUrlParams() {
    if (window.location.search.length <= 0) return false;
    var info = window.location.search.slice(1);
    var result = {};
    info.split('&').forEach(item => {
        result[decodeURIComponent(item.split('=')[0])] = decodeURIComponent(item.split('=')[1]);
    });
    return result;
}

// CopyText -- click event
function getcopy(copyid){
    var txt = document.getElementById(copyid).value;
    var oinput = document.createElement('input');
    oinput.value = txt;
    document.body.appendChild(oinput);
    oinput.select();
    document.execCommand('copy');
    document.body.removeChild(oinput);
    toast("复制成功!");
}