package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// SignResponse 签名的返回结构
type SignResponse struct {
	SourceStr string `json:"ostr"`
	Sign      string `json:"sign"`
	Status    string `json:"status:"`
}

const (
	// SignGenerating 签名正在生成中
	SignGenerating = "generating"
	// SignOK 签名计算OK
	SignOK = "ok"
)

// GetSign 返回请求的签名
func GetSign(token, ts string, try int, sleep time.Duration) (sign string, err error) {
	for i := 0; i < try; i++ {
		sign, err = getSign(token, ts)
		if err != nil {
			return "", err
		}
		if sign != "" {
			return sign, nil
		}
		time.Sleep(sleep * time.Second)
	}
	return "", errors.New("too many times when get sign")
}

// 注意 `signUrl，否则会报错，目前该接口不对外开放` !important!
func getSign(token, ts string) (string, error) {
	resp, err := http.Get(fmt.Sprintf(conf.SignURL, token, ts))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var signResp SignResponse
	err = json.Unmarshal(body, &signResp)
	if err != nil {
		return "", err
	}

	if signResp.Status == SignGenerating {
		return "", nil
	}

	if signResp.Status == SignOK {
		return signResp.Sign, nil
	}

	return "", errors.New("generate sign error")

}
