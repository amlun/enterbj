package enterbj

import (
	"github.com/amlun/enterbj/request"
	"time"
)

func (e *Client) verifyRequest(phone string) *request.Verification {
	var reqBody request.Verification
	reqBody.Phone = phone
	reqBody.Regist = "1"
	return &reqBody
}

func (e *Client) loginRequest(phone string, valicode string) *request.Login {
	var reqBody request.Login
	reqBody.DeviceType = e.session.DeviceType
	reqBody.Lon = 116.542162
	reqBody.Phone = phone
	reqBody.Timestamp = time.Now().Format("2006-01-01 00:00:00")
	reqBody.Source = "0"
	reqBody.Lat = 39.937293
	reqBody.Token = e.session.Token
	reqBody.DeviceId = e.session.DeviceId
	reqBody.AppKey = e.app.Key
	reqBody.ValiCode = valicode
	reqBody.VerType = "1"
	reqBody.Method = "login"
	return &reqBody
}

func (e *Client) personInfoRequest() *request.PersonInfo {
	var reqBody request.PersonInfo
	reqBody.Appkey = e.app.Key
	reqBody.Dicver = ""
	reqBody.SN = ""
	reqBody.UserId = e.session.UserId
	return &reqBody
}

func (e *Client) carListRequest() *request.CarList {
	var reqBody request.CarList
	reqBody.AppKey = "kkk"
	reqBody.AppSource = ""
	reqBody.DeviceId = "ddd"
	reqBody.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	reqBody.Token = e.session.Token
	reqBody.UserId = e.session.UserId
	reqBody.Platform = "02"
	return &reqBody
}

func (e *Client) checkEnvGradeRequest(carId string, licenseNo string, carModel string, carRegTime string) *request.CheckEnvGrade {
	var reqBody request.CheckEnvGrade
	reqBody.AppSource = "bjjj"
	reqBody.UserId = e.session.UserId
	reqBody.CarId = carId
	reqBody.LicenseNo = licenseNo
	reqBody.CarModel = carModel
	reqBody.CarRegTime = carRegTime
	return &reqBody
}

// TODO
func (e *Client) applySubmitRequest(licenseNo, engineNo, carTypeCode string) *request.SubmitPaper {
	var reqBody request.SubmitPaper
	reqBody.AppSource = "bjjj"
	reqBody.HiddenTime = time.Now().Format("2006-01-02 15:04:05")
	reqBody.InbjEntranceCode1 = 05
	reqBody.InbjEntranceCode = 12
	reqBody.InbjDuration = 7
	reqBody.InbjTime = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	reqBody.UserId = e.session.UserId
	reqBody.LicenseNo = licenseNo
	reqBody.EngineNo = engineNo
	reqBody.CarTypeCode = carTypeCode
	reqBody.VehicleType = "11"

	return &reqBody
}
