package enterbj

import (
	"bytes"
	"encoding/json"
	"github.com/amlun/enterbj/request"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"time"
)

const SimpleDate = "2006-01-02"
const SimpleDateTime = "2006-01-02 15:04:05"

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

func verifyRequest(phone string) *request.Verification {
	var reqBody request.Verification
	reqBody.Phone = phone
	reqBody.Regist = "1"
	return &reqBody
}

func loginRequest(phone string, valicode string) *http.Request {
	var reqBody request.Login
	reqBody.DeviceType = "0"
	reqBody.Lon = 116.542162
	reqBody.Phone = phone
	reqBody.Timestamp = time.Now().Format(SimpleDateTime)
	reqBody.Source = "0"
	reqBody.Lat = 39.937293
	reqBody.Token = ""
	reqBody.Platform = "01"
	reqBody.DeviceId = "dbf55511b74c4380c460faf3cc1f3bb7f51fec56"
	reqBody.CityCode = "1101"
	reqBody.AppKey = "0791682354"
	reqBody.ValiCode = valicode
	reqBody.VerType = "1"

	r, err := json.Marshal(reqBody)
	if err != nil {
		return nil
	}
	req, _ := http.NewRequest("POST", LOGIN_URL, bytes.NewBuffer(r))
	req.Header = commonHeader
	return req
}

func personInfoRequest(userId string) *http.Request {
	var reqBody request.PersonInfo
	reqBody.Appkey = ""
	reqBody.Dicver = ""
	reqBody.SN = ""
	reqBody.UserId = userId

	r, err := query.Values(reqBody)
	if err != nil {
		return nil
	}
	req, _ := http.NewRequest("GET", PERSON_INFO_URL+r.Encode(), nil)
	req.Header = commonHeader

	return req
}

func carListRequest(userId string) *http.Request {
	var reqBody request.CarList
	reqBody.AppKey = "kkk"
	reqBody.AppSource = ""
	reqBody.DeviceId = "ddd"
	reqBody.Timestamp = time.Now().Format(SimpleDateTime)
	reqBody.Token = "922C90208F834084AF118EE49D6F522F"
	reqBody.UserId = userId
	reqBody.Platform = "02"
	if sign, err := GetSign(reqBody.UserId, reqBody.Timestamp, 3, 2); err != nil { // 处理Sign
		return nil
	} else {
		reqBody.Sign = sign
	}

	r, err := query.Values(reqBody)
	if err != nil {
		return nil
	}
	req, _ := http.NewRequest("POST", CARLIST_URL, bytes.NewBufferString(r.Encode()))
	req.Header = commonHeader
	return req
}

func checkEnvGradeRequest(userId, carId, licenseNo, carModel, carRegTime string) *http.Request {
	var reqBody request.CheckEnvGrade
	reqBody.AppSource = "bjjj"
	reqBody.UserId = userId
	reqBody.CarId = carId
	reqBody.LicenseNo = licenseNo
	reqBody.CarModel = carModel
	reqBody.CarRegTime = carRegTime

	r, err := query.Values(reqBody)
	if err != nil {
		return nil
	}
	req, _ := http.NewRequest("POST", CHECK_ENV_GRADE_URL, bytes.NewBufferString(r.Encode()))
	req.Header = commonHeader

	return req
}

// TODO
func applySubmitRequest(userId, licenseNo, engineNo, carTypeCode string) *http.Request {
	var reqBody request.SubmitPaper
	reqBody.AppSource = "bjjj"
	now := time.Now().Format(SimpleDateTime)
	reqBody.Timestamp = now
	reqBody.HiddenTime = now
	reqBody.InbjEntranceCode1 = 05
	reqBody.InbjEntranceCode = 12
	reqBody.InbjDuration = 7
	reqBody.InbjTime = time.Now().AddDate(0, 0, 1).Format(SimpleDate)
	reqBody.UserId = userId
	reqBody.LicenseNo = licenseNo
	reqBody.EngineNo = engineNo
	reqBody.CarTypeCode = carTypeCode
	reqBody.VehicleType = "11"
	if sign, err := GetSign(reqBody.UserId, reqBody.Timestamp, 3, 2); err != nil { // 处理Sign
		return nil
	} else {
		reqBody.Sign = sign
	}

	r, err := query.Values(reqBody)
	if err != nil {
		return nil
	}
	req, _ := http.NewRequest("POST", SUBMIT_PAPER_URL, bytes.NewBufferString(r.Encode()))
	req.Header = commonHeader

	return req
}

func sendRequest(req *http.Request, v interface{}) (resp *http.Response, err error) {
	resp, err = httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
