package enterbj

import (
	"net/http"
)

var commonHeader = http.Header{
	"Host":             []string{"api.jinjingzheng.zhongchebaolian.com"},
	"accept":           []string{"*/*"},
	"x-requested-with": []string{"XMLHttpRequest"},
	"accept-language":  []string{"zh-cn"},
	"content-type":     []string{"application/x-www-form-urlencoded"},
	"origin":           []string{"https://api.jinjingzheng.zhongchebaolian.com"},
	"user-agent":       []string{"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89"},
	"referer":          []string{"https://api.jinjingzheng.zhongchebaolian.com/enterbj/jsp/enterbj/index.jsp"},
	"cookie":           []string{"JSESSIONID=9E56E86F02184BF5E1D8BC9C05C5D76C"},
}