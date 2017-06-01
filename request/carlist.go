package request

type CarList struct {
	UserId    string `url:"userid"`
	AppKey    string `url:"appkey"`
	DeviceId  string `url:"deviceid"`
	Timestamp int64  `url:"timestamp"`
	Token     string `url:"token"`
	AppSource string `url:"appsource"`
}
