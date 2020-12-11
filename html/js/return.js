// 禁止双指放大
document.documentElement.addEventListener('touchstart', function (event) {
    if (event.touches.length > 1) {
        event.preventDefault();
    }
}, false);
// 禁止双击放大
var lastTouchEnd = 0;
document.documentElement.addEventListener('touchend', function (event) {
    var now = Date.now();
    if (now - lastTouchEnd <= 300) {
        event.preventDefault();
    }
    lastTouchEnd = now;
}, false);
document.addEventListener('gesturestart', function(event) {
    event.preventDefault();
});

window.addEventListener("popstate", function(e) {
    //alert("我监听到了浏览器的返回按钮事件啦");//根据自己的需求实现自己的功能
}, false);
$(document).ready(function(e) {
    var counter = 0;
    if (window.history && window.history.pushState) {
        $(window).on('popstate', function () {
            window.history.pushState('forward', null, '#');
            window.history.forward(1);
            /*layer.open({
                content: '不可回退！',
                style: 'background-color:#EF7F00; color:#fff; border:none;',
                time: 2
            });*/
        });
    }

    window.history.pushState('forward', null, '#'); //在IE中必须得有这两行  
    window.history.forward(1);
});
/*var count = 0 ;
window.history.pushState(null, null, "#");
window.addEventListener("popstate", function(e) {
    window.history.pushState(null, null, "#");
    alert("用户点击返回" + (++count))

})*/