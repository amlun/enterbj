package enterbj

import "time"

type Config struct {
	AppKey    string
	AppSource string
	DeviceId  string
	Token     string
	Platform  string
	SignUrl   string
	Timeout   time.Duration
}
