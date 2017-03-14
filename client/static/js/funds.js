$(function() {
    if (is_mobile()) {
        //$('body').removeClass('left-side-collapsed');
    }

    //左边菜单加选中状态
    var pre = location.pathname;
    var qstr = pre.split('/');
    if (qstr) {
        var lefthref = '/' + qstr[1] + '/' + qstr[2];
        $('.sub-menu-list a').filter(function() {
            return $(this).attr('href') == lefthref;
        }).parent().addClass('active').parents('.menu-list').addClass('nav-active');
    };

    $('.js_checkboxAll').on('click', function() {
        var that = $(this);
        var chk = that.parent().prev('table').find('input[type="checkbox"]');
        if (that.is(':checked')) {
            chk.prop('checked', true);
        } else {
            chk.prop('checked', false);
        }
    });

    // setInterval(function myInterval() {
    //     // alert('aaaaaaaaa');
    //     var url = '/my/fund';
    //     $('#buyfund-form').ajaxSubmit({
    //         url: url,
    //         type: 'GET',
    //         dataType: 'json',
    //         success: function(data) {
    //             var myaccount = $(this).document.getElementsByName("myaccount")[0];
    //             myaccount.setAttribute("value", data.myaccount);
    //         }
    //     });
    // }, 5000);

    $('#login-form').validate({
        ignore: '',
        rules: {
            username: { required: true },
            password: { required: true }
        },
        messages: {
            username: { required: '请填写用户名' },
            password: { required: '请填写密码' }
        },
        submitHandler: function(form) {
            var url = '/login';
            $(form).ajaxSubmit({
                url: url,
                type: 'POST',
                dataType: 'json',
                success: function(data) {
                    if (data.code) {
                        setTimeout(function() { window.location.href = "/my/index" }, 0);
                    } else {
                        dialogInfo(data.message)
                        setTimeout(function() { $('#dialogInfo').modal('hide'); }, 1000);
                    }
                }
            });
        }
    });

    $('#create-form').validate({
        ignore: '',
        submitHandler: function(form) {
            var url = '/currency/create';
            $(form).ajaxSubmit({
                url: url,
                type: 'POST',
                dataType: 'json',
                success: function(data) {
                    dialogInfo(data.message);
                    if (data.code) {
                        setTimeout(function() { window.location.href = "/my/index" }, 1000);
                    } else {
                        setTimeout(function() { $('#dialogInfo').modal('hide'); }, 1000);
                    }
                }
            });
        }
    });

    $('#release-form').validate({
        ignore: '',
        submitHandler: function(form) {
            var url = '/currency/release';
            $(form).ajaxSubmit({
                url: url,
                type: 'POST',
                dataType: 'json',
                success: function(data) {
                    dialogInfo(data.message);
                    if (data.code) {
                        setTimeout(function() { window.location.href = "/my/index" }, 1000);
                    } else {
                        setTimeout(function() { $('#dialogInfo').modal('hide'); }, 1000);
                    }
                }
            });
        }
    });

    $('#assign-form').validate({
        ignore: '',
        submitHandler: function(form) {
            var url = '/currency/assign';
            $(form).ajaxSubmit({
                url: url,
                type: 'POST',
                dataType: 'json',
                success: function(data) {
                    dialogInfo(data.message);
                    if (data.code) {
                        setTimeout(function() { window.location.href = "/my/index" }, 1000);
                    } else {
                        setTimeout(function() { $('#dialogInfo').modal('hide'); }, 1000);
                    }
                }
            });
        }
    });

    $('#exchange-form').validate({
        ignore: '',
        submitHandler: function(form) {
            var url = '/tx/exchange';
            $(form).ajaxSubmit({
                url: url,
                type: 'POST',
                dataType: 'json',
                success: function(data) {
                    dialogInfo(data.message);
                    if (data.code) {
                        setTimeout(function() { window.location.href = "/my/index" }, 1000);
                    } else {
                        setTimeout(function() { $('#dialogInfo').modal('hide'); }, 1000);
                    }
                }
            });
        }
    });

    $(".cancel-class").click(function(e) {
        var submitData = {
            id: $(this).attr("data-id"),
        };

        $.post('/tx/cancel', submitData,
            function(data) {
                dialogInfo(data.message);
                if (data.code) {
                    setTimeout(function() { window.location.href = "/my/index" }, 1000);
                } else {
                    setTimeout(function() { $('#dialogInfo').modal('hide'); }, 1000);
                }
            },
            "json");
        e.preventDefault();
    });

});

function dialogInfo(msg) {
    $('#dialogInfo').remove();
    var html = '';
    html = '<div class="modal fade" id="dialogInfo" tabindex="-1" role="dialog" aria-labelledby="dialogInfoTitle">';
    html += '<div class="modal-dialog" role="document">';
    html += '<div class="modal-content">';
    html += '<div class="modal-header">';
    html += '<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>';
    html += '<h4 class="modal-title" id="dialogInfoTitle">友情提示</h4>';
    html += ' </div>';
    html += '<div class="modal-body">';
    html += '<p>' + msg + '</p>';
    html += '</div>';
    //html += '<div class="modal-footer">';
    //html += ' <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>';
    //html += ' <button type="button" class="btn btn-primary">Send message</button>';
    //html += '</div>';
    html += '</div>';
    html += '</div>';
    html += '</div>';
    $('body').append(html);
    $('#dialogInfo').modal('show')
}

function dialogAlbum(id, title, summary, status) {
    $('#dialogAlbum').remove();
    var html = '';
    html += '<div class="modal fade in" id="dialogAlbum" tabindex="-1" role="dialog" aria-labelledby="dialogAlbumTitle">';
    html += '  <div class="modal-dialog" role="document">';
    html += '   <form id="album-form">';
    html += '    <div class="modal-content">';
    html += '      <div class="modal-header">';
    html += '        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">×</span></button>';
    html += '        <h4 class="modal-title" id="dialogAlbumTitle">编辑' + title + '</h4>';
    html += '      </div>';
    html += '      <div class="modal-body">';
    html += '          <div class="form-group">';
    html += '            <label for="recipient-name" class="control-label">标题:</label>';
    html += '           <input type="text" class="form-control" name="title" value="' + title + '">';
    html += '          </div>';

    html += '          <div class="form-group">';
    html += '            <label for="message-text" class="control-label">说明:</label>';
    html += '            <textarea class="form-control" name="summary">' + summary + '</textarea>';
    html += '          </div>';


    html += '<div class="form-group">';
    html += '<label class="radio-inline">';
    html += '<input type="radio" name="status" value="0" ' + (status == 0 ? 'checked' : '') + '> 屏蔽';
    html += '</label>';
    html += '<label class="radio-inline">';
    html += '<input type="radio" name="status" value="1" ' + (status == 1 ? 'checked' : '') + '> 正常';
    html += '</label>';
    html += '          </div>';



    html += '      </div>';
    html += '      <div class="modal-footer"><input type="hidden" name="id" value="' + id + '">';
    html += '        <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>';
    html += '        <button type="button" class="btn btn-primary js-album-submit">提交</button>';
    html += '      </div>';
    html += '    </div>';
    html += '   </form>';
    html += '  </div>';
    html += '</div>';
    $('body').append(html);
    $('#dialogAlbum').modal('show');
}

function workDay(started, ended) {
    var beginDate = new Date(started.replace(/-/g, "/"));
    //结束日期  
    var endDate = new Date(ended.replace(/-/g, "/"));
    //日期差值,即包含周六日、以天为单位的工时，86400000=1000*60*60*24.  
    var workDayVal = (endDate - beginDate) / 86400000 + 1;
    //工时的余数  
    var remainder = workDayVal % 7;
    //工时向下取整的除数  
    var divisor = Math.floor(workDayVal / 7);
    var weekendDay = 2 * divisor;

    //起始日期的星期，星期取值有（1,2,3,4,5,6,0）  
    var nextDay = beginDate.getDay();
    //从起始日期的星期开始 遍历remainder天  
    for (var tempDay = remainder; tempDay >= 1; tempDay--) {
        //第一天不用加1  
        if (tempDay == remainder) {
            nextDay = nextDay + 0;
        } else if (tempDay != remainder) {
            nextDay = nextDay + 1;
        }
        //周日，变更为0  
        if (nextDay == 7) {
            nextDay = 0;
        }

        //周六日  
        if (nextDay == 0 || nextDay == 6) {
            weekendDay = weekendDay + 1;
        }
    }
    //实际工时（天） = 起止日期差 - 周六日数目。  
    workDayVal = workDayVal - weekendDay;
    return workDayVal;
}

function is_mobile() {
    var regex_match = /(nokia|iphone|android|motorola|^mot-|softbank|foma|docomo|kddi|up.browser|up.link|htc|dopod|blazer|netfront|helio|hosin|huawei|novarra|CoolPad|webos|techfaith|palmsource|blackberry|alcatel|amoi|ktouch|nexian|samsung|^sam-|s[cg]h|^lge|ericsson|philips|sagem|wellcom|bunjalloo|maui|symbian|smartphone|midp|wap|phone|windows ce|iemobile|^spice|^bird|^zte-|longcos|pantech|gionee|^sie-|portalmmm|jigs browser|hiptop|^benq|haier|^lct|operas*mobi|opera*mini|320x320|240x320|176x220)/i;
    var u = navigator.userAgent;
    if (null == u) {
        return true;
    }
    var result = regex_match.exec(u);
    if (null == result) {
        return false
    } else {
        return true
    }
}