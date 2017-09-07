package enterbj

import (
	"time"
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	"io/ioutil"
)

type SignResponse struct {
	SourceStr string `json:"ostr"`
	Sign      string `json:"sign"`
	Status    string `json:"status:"`
}

const (
	SIGN_GENERATING = "generating"
	SIGN_OK         = "ok"
)

func GetSign(token, ts string, try int, sleep time.Duration, signUrl string) (sign string, err error) {
	for i := 0; i < try; i++ {
		sign, err = getSign(token, ts, signUrl)
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

// 注意，必须要设定SIGN_URL，否则会报错，目前该接口不对外开放 !important!
func getSign(token, ts, signUrl string) (string, error) {
	signUrl = fmt.Sprintf(signUrl, token, ts)
	resp, err := http.Get(signUrl)
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

	if signResp.Status == SIGN_GENERATING {
		return "", nil
	}

	if signResp.Status == SIGN_OK {
		return signResp.Sign, nil
	}

	return "", errors.New("generate sign error")

}
