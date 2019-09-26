$(function () {
    // 跳出提示后，关闭提示框
    $("#alert").bind('DOMNodeInserted', function () {
        window.setTimeout(function () {
            $("#alert").html(null)
        }, 1000);
    });

    // 鼠标点击小心心
    var click_cnt = 0;
    var $html = document.getElementsByTagName("html")[0];
    var $body = document.getElementsByTagName("body")[0];
    $html.onclick = function (e) {
        var $elem = document.createElement("b");
        $elem.style.color = "#E94F06";
        $elem.style.zIndex = 9999;
        $elem.style.position = "absolute";
        $elem.style.select = "none";
        var x = e.pageX;
        var y = e.pageY;
        $elem.style.left = (x - 10) + "px";
        $elem.style.top = (y - 20) + "px";
        clearInterval(anim);
        switch (++click_cnt) {
            case 10:
                $elem.innerText = "OωO";
                break;
            case 20:
                $elem.innerText = "(๑•́ ∀ •̀๑)";
                break;
            case 30:
                $elem.innerText = "(๑•́ ₃ •̀๑)";
                break;
            case 40:
                $elem.innerText = "(๑•̀_•́๑)";
                break;
            case 50:
                $elem.innerText = "（￣へ￣）";
                break;
            case 60:
                $elem.innerText = "(╯°口°)╯(┴—┴";
                break;
            case 70:
                $elem.innerText = "૮( ᵒ̌皿ᵒ̌ )ა";
                break;
            case 80:
                $elem.innerText = "╮(｡>口<｡)╭";
                break;
            case 90:
                $elem.innerText = "( ง ᵒ̌皿ᵒ̌)ง⁼³₌₃";
                break;
            case 100:
            case 101:
            case 102:
            case 103:
            case 104:
            case 105:
                $elem.innerText = "(ꐦ°᷄д°᷅)";
                break;
            default:
                $elem.innerText = "❤";
                break;
        }
        $elem.style.fontSize = Math.random() * 10 + 8 + "px";
        var increase = 0;
        var anim;
        setTimeout(function () {
            anim = setInterval(function () {
                if (++increase == 150) {
                    clearInterval(anim);
                    $body.removeChild($elem);
                }
                $elem.style.top = y - 20 - increase + "px";
                $elem.style.opacity = (150 - increase) / 120;
            }, 8);
        }, 70);
        $body.appendChild($elem);
    };
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
    // 隐藏“网络图片”tab
    editor.customConfig.showLinkImg = false
    // base64插入图片
    editor.customConfig.uploadImgShowBase64 = true
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
    editor.customConfig.uploadImgMaxLength = 9 // 单次最大上传数量
    editor.customConfig.uploadImgMaxSize = 10 * 1024 * 1024 // 图片大小限制
    editor.customConfig.uploadImgParams = { // 上传参数补充
        _xsrf: XSRF
    }
    button_submit_text = ""
    editor.customConfig.uploadImgHooks = {
        before: function (xhr, editor, files) {
            button_submit_text= $('button[type="submit"]').text()
            $('button[type="submit"]').text('图片上传中,稍加等待...')
            $('button[type="submit"]').attr("disabled", "true")
        },
        success: function (xhr, editor, result) {  
            $("#alert").html('<div class="text-center alert alert-success" role="alert">上传完成</div>')
            $('button[type="submit"]').text(button_submit_text)
            $('button[type="submit"]').removeAttr("disabled")
        },
        fail: function (xhr, editor, result) {   
            $("#alert").html('<div class="text-center alert alert-danger" role="alert">上传失败</div>')       
            $('button[type="submit"]').text(button_submit_text)
            $('button[type="submit"]').removeAttr("disabled")
        },
        error: function (xhr, editor) {    
            $("#alert").html('<div class="text-center alert alert-danger" role="alert">上传出错</div>')       
            $('button[type="submit"]').text(button_submit_text)
            $('button[type="submit"]').removeAttr("disabled")    
        },
        timeout: function (xhr, editor) {
            $("#alert").html('<div class="text-center alert alert-danger" role="alert">上传超时</div>')       
            $('button[type="submit"]').text(button_submit_text)
            $('button[type="submit"]').removeAttr("disabled")
        },    
        customInsert: function (insertImg, result) {  // 返回值设置
            var url = result.data.path
            insertImg(url)
        }
    }
    editor.customConfig.onchange = function (html) { // 内容填充
        $('input[name="editor"]').val(html)
    }

    return editor;
}