package core

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/amlun/enterbj/request"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// SimpleDate 简单的日期格式
	SimpleDate = "2006-01-02"
	// SimpleDateTime 简单的日期时间格式
	SimpleDateTime = "2006-01-02 15:04:05"
	// CarListURL 车辆列表
	CarListURL = "https://enterbj.zhongchebaolian.com/enterbj/platform/enterbj/entercarlist"
	// LoginURL 登录
	LoginURL = "https://bjjj.zhongchebaolian.com/industryguild_mobile_standard_self2.1.2/mobile/standard/login"
	// SubmitPaperURL 提交进京证申请
	SubmitPaperURL = "https://enterbj.zhongchebaolian.com/enterbj-img/platform/enterbj/submitpaper_03"
	// PersonInfoURL 个人信息
	PersonInfoURL = "https://bjjj.zhongchebaolian.com/industryguild_mobile_standard_self2.1.2/mobile/standard/getpersonalinfor?"
	// CheckEnvGradeURL 检查环保信息
	CheckEnvGradeURL = "https://enterbj.zhongchebaolian.com/enterbj/platform/enterbj/checkenvgrade"
	//LoadOtherDriversUrl = "https://api.jinjingzheng.zhongchebaolian.com/enterbj/platform/enterbj/loadotherdrivers"
	CheckServiceStatusURL = "https://enterbj.zhongchebaolian.com/enterbj/platform/enterbj/curtime_03?userid="
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
		logrus.Error(err)
		return nil
	}
	req, _ := http.NewRequest("POST", LoginURL, bytes.NewBuffer(r))
	req.Header = commonHeader
	return req
}

func personInfoRequest(userID string) *http.Request {
	var reqBody request.PersonInfo
	reqBody.Appkey = ""
	reqBody.Dicver = ""
	reqBody.SN = ""
	reqBody.UserId = userID

	r, err := query.Values(reqBody)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	req, _ := http.NewRequest("GET", PersonInfoURL+r.Encode(), nil)
	req.Header = commonHeader

	return req
}

func carListRequest(userID string) *http.Request {
	var reqBody request.CarList
	reqBody.AppKey = "kkk"
	reqBody.AppSource = ""
	reqBody.DeviceId = "ddd"
	reqBody.Timestamp = time.Now().Format(SimpleDateTime)
	reqBody.Token = "922C90208F834084AF118EE49D6F522F"
	reqBody.UserId = userID
	reqBody.Platform = "02"
	sign, err := GetSign(reqBody.UserId, reqBody.Timestamp, 3, 2)
	if err != nil { // 处理Sign
		logrus.Error(err)
		return nil
	}
	reqBody.Sign = sign

	r, err := query.Values(reqBody)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	req, _ := http.NewRequest("POST", CarListURL, bytes.NewBufferString(r.Encode()))
	req.Header = commonHeader
	return req
}

func checkEnvGradeRequest(userID, carID, licenseNo, carModel, carRegTime string) *http.Request {
	var reqBody request.CheckEnvGrade
	reqBody.AppSource = "bjjj"
	reqBody.UserId = userID
	reqBody.CarId = carID
	reqBody.LicenseNo = licenseNo
	reqBody.CarModel = carModel
	reqBody.CarRegTime = carRegTime

	r, err := query.Values(reqBody)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	req, _ := http.NewRequest("POST", CheckEnvGradeURL, bytes.NewBufferString(r.Encode()))
	req.Header = commonHeader

	return req
}

// TODO
func applySubmitRequest(userID, licenseNo, engineNo, drivingPhotoPath, carPhotoPath, driverName, driverLicenseNo,
	driverPhotoPath, personPhotoPath, carID, carModel, carRegTime, envGrade string) *http.Request {
	var reqBody request.SubmitPaper
	reqBody.AppSource = "bjjj"
	now := time.Now().Format(SimpleDateTime)
	reqBody.Timestamp = now
	reqBody.HiddenTime = now
	reqBody.InbjEntranceCode1 = 05
	reqBody.InbjEntranceCode = 12
	reqBody.InbjDuration = 7
	reqBody.InbjTime = time.Now().AddDate(0, 0, 1).Format(SimpleDate)
	reqBody.CarTypeCode = "02"
	reqBody.VehicleType = "03"
	reqBody.UserId = userID                                  // 用户ID
	reqBody.LicenseNo = licenseNo                            // 车牌号
	reqBody.EngineNo = engineNo                              // 发动机编号
	reqBody.DrivingPhoto = readBase64Image(drivingPhotoPath) // 获取车辆行驶证 base64
	reqBody.CarPhoto = readBase64Image(carPhotoPath)         // 获取车辆正面照 base64
	reqBody.DriverName = driverName                          // 驾驶人姓名
	reqBody.DriverLicenseNo = driverLicenseNo                // 驾照编码
	reqBody.DriverPhoto = readBase64Image(driverPhotoPath)   // 驾驶人证件 base64
	reqBody.PersonPhoto = readBase64Image(personPhotoPath)   // 驾驶人手持身份证 base64
	reqBody.CarId = carID                                    // 车辆注册编号
	reqBody.CarModel = carModel                              // 车辆型号
	reqBody.CarRegTime = carRegTime                          // 车辆注册时间
	reqBody.EnvGrade = envGrade                              // 环保标准
	reqBody.VehicleType = "11"                               // 机动车类型
	reqBody.PlatForm = "02"                                  // 平台，默认使用安卓平台
	reqBody.ImageID = ""

	sign, err := GetSign(reqBody.UserId, reqBody.Timestamp, 3, 2)
	if err != nil { // 处理Sign
		logrus.Error(err)
		return nil
	}
	reqBody.Sign = sign

	r, err := query.Values(reqBody)
	if err != nil {
		return nil
	}
	req, _ := http.NewRequest("POST", SubmitPaperURL, bytes.NewBufferString(r.Encode()))
	req.Header = commonHeader

	return req
}

func checkServiceStatus(userID string) *http.Request {
	req, _ := http.NewRequest("POST", CheckServiceStatusURL+userID, bytes.NewBufferString(""))
	req.Header = commonHeader
	return req
}

func readBase64Image(fpath string) string {
	buf, err := ioutil.ReadFile(fpath)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(buf)
}

func sendRequest(req *http.Request, v interface{}) (resp *http.Response, err error) {
	logrus.Debugf("send request (%v)", req)
	resp, err = httpClient.Do(req)
	logrus.Debugf("receive response (%v)", resp)
	if err != nil {
		logrus.Errorf("send request error (%v)", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	logrus.Debugf("receive response body (%s)", body)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return resp, nil
}
