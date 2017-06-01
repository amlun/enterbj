package request

type PersonInfo struct {
	Appkey string `url:"appkey"`
	Dicver string `url:"dicver"`
	SN     string `url:"sn"`
	UserId string `url:"userid"`
}
