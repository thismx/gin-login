var handleResponseError = function (response) {
    if (response.status == 401) {
        window.location.href = '/login';
    }
    var msg = response.status == 500 ? response.statusText : $.response_shift(response);
    $.messager.alert("操作失败", msg);
    return false;
};

var handleResponseMsg = function(response) {
    var msg = '';
    try{
        if (response.status == 401) {
            window.location.href = '/login';
            return false;
        }else if(response.status == 403) {
            msg = '没有权限访问该页面';
        }else if(response.status == 404) {
            return false;
        }else{
            msg = response.responseJSON == undefined ? response.json() : response.responseJSON;
            $.each(msg.errors, function(i, n){
                msg = typeof(n) == "string" ? n : n[0];
                return false;
            })
        }
    }
    catch(err)
    {
        msg = '未知错误信息, 请联系网站管理员';
    }

    return msg;
};
var AJAXErrorByLayer = function (response) {
    if(response.status == 404){
        return false;
    }
    layer.alert(handleResponseMsg(response), {icon: 2});
    return false;
};