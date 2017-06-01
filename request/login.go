package request

import "time"

type Login struct {
	DeviceType string    `json:"devicetype"`
	Sign       string    `json:"sign"`
	Lon        float32   `json:"lon"`
	Phone      string    `json:"phone"`
	Timestamp  time.Time `json:"timestamp"`
	Source     string    `json:"source"`
	Lat        float32   `json:"lat"`
	Token      string    `json:"token"`
	DeviceId   string    `json:"deviceid"`
	AppKey     string    `json:"appkey"`
	ValiCode   string    `json:"valicode"`
	VerType    string    `json:"vertype"`
	Method     string    `json:"method"`
}
