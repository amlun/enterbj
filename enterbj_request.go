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
	reqBody.DeviceType = ""
	reqBody.Lon = 116.542162
	reqBody.Phone = phone
	reqBody.Timestamp = time.Now().Format("2006-01-01 00:00:00")
	reqBody.Source = "0"
	reqBody.Lat = 39.937293
	reqBody.Token = "922C90208F834084AF118EE49D6F522F"
	reqBody.DeviceId = ""
	reqBody.AppKey = ""
	reqBody.ValiCode = valicode
	reqBody.VerType = "1"
	reqBody.Method = "login"
	return &reqBody
}

func (e *Client) personInfoRequest(userId string) *request.PersonInfo {
	var reqBody request.PersonInfo
	reqBody.Appkey = ""
	reqBody.Dicver = ""
	reqBody.SN = ""
	reqBody.UserId = userId
	return &reqBody
}

func (e *Client) carListRequest(userId string) *request.CarList {
	var reqBody request.CarList
	reqBody.AppKey = "kkk"
	reqBody.AppSource = ""
	reqBody.DeviceId = "ddd"
	reqBody.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	reqBody.Token = "922C90208F834084AF118EE49D6F522F"
	reqBody.UserId = userId
	reqBody.Platform = "02"
	return &reqBody
}

func (e *Client) checkEnvGradeRequest(userId, carId, licenseNo, carModel, carRegTime string) *request.CheckEnvGrade {
	var reqBody request.CheckEnvGrade
	reqBody.AppSource = "bjjj"
	reqBody.UserId = userId
	reqBody.CarId = carId
	reqBody.LicenseNo = licenseNo
	reqBody.CarModel = carModel
	reqBody.CarRegTime = carRegTime
	return &reqBody
}

// TODO
func (e *Client) applySubmitRequest(userId, licenseNo, engineNo, carTypeCode string) *request.SubmitPaper {
	var reqBody request.SubmitPaper
	reqBody.AppSource = "bjjj"
	reqBody.HiddenTime = time.Now().Format("2006-01-02 15:04:05")
	reqBody.InbjEntranceCode1 = 05
	reqBody.InbjEntranceCode = 12
	reqBody.InbjDuration = 7
	reqBody.InbjTime = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	reqBody.UserId = userId
	reqBody.LicenseNo = licenseNo
	reqBody.EngineNo = engineNo
	reqBody.CarTypeCode = carTypeCode
	reqBody.VehicleType = "11"

	return &reqBody
}
