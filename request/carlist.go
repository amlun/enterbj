package request

type CarList struct {
	UserId    string `url:"userid"`
	AppKey    string `url:"appkey"`
	DeviceId  string `url:"deviceid"`
	Timestamp string `url:"timestamp"`
	Token     string `url:"token"`
	AppSource string `url:"appsource"`
	Platform  string `url:"platform"`
	Sign      string `url:"sign"`
}
