
function closeWin() {
    if (navigator.userAgent.indexOf("Firefox") != -1 || navigator.userAgent.indexOf("Chrome") != -1) {
        window.location.href = "about:blank";
        window.close();
    } else {
        window.opener = null;
        window.open("", "_self");
        window.close();
    }
}

function GetQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]); return null;
}

$(document).ready(function () {
    if (GetQueryString("clickId") == null) {
        clickid = GetQueryString("aff_click")
    } else {
        clickid = GetQueryString("clickId")
    }
    $.ajax({
        type: "GET",
        url: "http://cpx2.allcpx.com/tn/uk/lpreq",
        data: {
            type: "game_w",
            affName: GetQueryString("affName"),
            clickId: clickid,
            pubId: GetQueryString("pubId"),
            proId: GetQueryString("proId"),
        }
    }).done(function (result) {
        ptxid = result
    })
});

$('#submit_btn_mobile').click(function () {
    $.ajax({
        type: "GET",
        url: "http://cpx2.allcpx.com/tn/uk/lpreq",
        data: {
            trackId: ptxid,
        }
    }).done(function (result) {
        if (result != "false") {
            window.location.href = result
        } else {
            window.location.href = "http://www.gogamehub.com/lp/tn/uk/opps/index.html"
        }
    })
});

	// $('#submit_btn_mobile_cancel').click(function () {
	// 	var url = "http://<broadcaster_url>/partner/apiv3/delete/?" +
	// 		"login=" + "CPX" +
	// 		"&rubric_id=" + "cpx" +
	// 		"&signature=" + "cpx" +
	// 		"&msisdn=" + "xcpx"
	// 	alert(url)
	// 	// window.location.href = url
	// });
