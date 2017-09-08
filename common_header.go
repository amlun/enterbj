package enterbj

import (
	"net/http"
)

var commonHeader = http.Header{
	"Host":             []string{"enterbj.zhongchebaolian.com"},
	"Accept":           []string{"*/*"},
	"X-Requested-With": []string{"XMLHttpRequest"},
	"Accept-Language":  []string{"zh-cn"},
	"Content-Type":     []string{"application/x-www-form-urlencoded; charset=UTF-8"},
	"Origin":           []string{"https://enterbj.zhongchebaolian.com"},
	"User-Agent":       []string{"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89"},
	"Referer":          []string{"https://enterbj.zhongchebaolian.com/enterbj/jsp/enterbj/index.html"},
	"Cookie":           []string{"JSESSIONID=9E56E86F02184BF5E1D8BC9C05C5D76C"},
}
