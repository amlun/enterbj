package request


type CarList struct {
	UserId    string `url:"userid"`
	AppKey    string `url:"appkey"`
	DeviceId  string `url:"deviceid"`
	Timestamp string `url:"timestamp"`
	Token     string `url:"token"`
	Sign	  string `url:"sign"`
	Platform  string `url:"platform"`		// 01: ios 02: android
	AppSource string `url:"appsource"`
}
