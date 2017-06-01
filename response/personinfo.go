package response

type PersonInfo struct {
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Sex             string `json:"sex"`
	HeadUrl         string `json:"headurl"`
	DriverLicenseNo string `json:"driverlicenseno"`
	Res
}
