$(function () {
    // 跳出提示后，关闭提示框
    $("#alert").bind('DOMNodeInserted', function () {
        window.setTimeout(function () {
            $("#alert").html(null)
        }, 1000);
    });

});

// 获取form数据
function getFormJson(frm) {
    var o = {};
    var a = $(frm).serializeArray()
    $.each(a, function () {
        o[this.name] = this.value || ""
    })
    return o;
}

// ajax上传form数据
function ajaxSubmit(frm, fn) {
    var dat = getFormJson(frm)
    return $.ajax({
        url: frm.action,
        type: frm.method,
        data: dat,
        success: fn
    })
}

// 将base64转换为文件
function dataURLtoFile(dataurl, filename) {
    var arr = dataurl.split(','),
        mime = arr[0].match(/:(.*?);/)[1],
        bstr = atob(arr[1]),
        n = bstr.length,
        u8arr = new Uint8Array(n);
    while (n--) {
        u8arr[n] = bstr.charCodeAt(n);
    }
    return new File([u8arr], filename, {
        type: mime
    });
}

// 随机字符串
function randomString(len = 30) {
    len = len || 32;
    var $chars = 'ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678'; /****默认去掉了容易混淆的字符oOLl,9gq,Vv,Uu,I1****/
    var maxPos = $chars.length;
    var pwd = '';
    for (i = 0; i < len; i++) {
        pwd += $chars.charAt(Math.floor(Math.random() * maxPos));
    }
    return pwd;
}

// 富文本配置
function editorInit(editor) {
    // 菜单栏目
    editor.customConfig.menus = [
        'bold',
        'italic',
        'link',
        'list',
        'emoticon',
        'image',
    ]
    // 表情
    editor.customConfig.emotions = [
        {
            title: 'emoji',
            type: 'emoji',
            content: ["😀", "😃", "😄", "😁", "😆", "😅", "😂", "😊", "😇", "🙂", "🙃", "😉", "😓", "😪", "😴", "🙄", "🤔", "😬", "🤐"]
        }
    ]
    editor.customConfig.uploadImgServer = '/upload' // 上传url
    editor.customConfig.uploadFileName = 'file' // 上传图片name
    editor.customConfig.uploadImgMaxLength = 1 // 单次最大上传数量
    editor.customConfig.uploadImgMaxSize = 10 * 1024 * 1024 // 图片大小限制
    editor.customConfig.uploadImgParams = { // 上传参数补充
        _xsrf: XSRF
    }
    editor.customConfig.uploadImgHooks = {
        customInsert: function (insertImg, result) {  // 返回值设置
            insertImg(result.data.path)
        }
    }
    editor.customConfig.onchange = function (html) { // 内容填充
        $('input[name="editor"]').val(html)
    }

    return editor;
}