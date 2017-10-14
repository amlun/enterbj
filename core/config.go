package core

import "time"

// Config 进京证客户端需要的一些配置信息
type Config struct {
	AppKey    string
	AppSource string
	DeviceID  string
	Token     string
	Platform  string
	SignURL   string
	Timeout   time.Duration
}
