$(document).click(function () {
    $.ajax({
        url : "/db_info",//请求地址
        dataType : "json",//数据格式
        type : "post",//请求方式
        async : false,//是否异步请求
        success : function(data) { //如何发送成功
            html= JSON.stringify(data);
            $("#test").html(html);
        },
    })
})