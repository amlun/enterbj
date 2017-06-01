package response

type Login struct {
	AccessToken  interface{} `json:"accesstoken"`
	UserId       string      `json:"userid"`
	ProvinceCode interface{} `json:"provincecode"`
	CityCode     string      `json:"citycode"`
	PoliceNo     string      `json:"policeno"`
	Ssid         string      `json:"ssid"`
	ProvinceTiny interface{} `json:"provincetiny"`
	UserType     string      `json:"userType"`
	SSID         string      `json:"ssid"`
	Res
}
