package request

type CheckEnvGrade struct {
	AppSource  string `url:"appsource"`  // bjjj
	UserId     string `url:"userid"`     // 用户id
	CarId      string `url:"carid"`      // 车辆注册ID
	LicenseNo  string `url:"licenseno"`  // 车牌号
	CarModel   string `url:"carmodel"`   // 车辆型号
	CarRegTime string `url:"carregtime"` // 车辆注册时间
}
