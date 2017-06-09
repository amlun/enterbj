package request

type SubmitPaper struct {
	AppSource         string `url:"appsource"`         // bjjj
	HiddenTime        string `url:"hiddentime"`        // 当前时间(2017-06-09 14:24:37)
	InbjEntranceCode1 int    `url:"inbjentrancecode1"` // 进入口(编码） 05 朝阳
	InbjEntranceCode  int    `url:"inbjentrancecode"`  // 进入口(编码)  12 京哈
	InbjDuration      int    `url:"inbjduration"`      // 进入时长 最多7天，最少2天
	InbjTime          string `url:"inbjtime"`          // 进入时间 明天-4天内 (2017-6-10)
	AppKey            string `url:"appkey"`            // 留空
	Deviceid          string `url:"deviceid"`          // 留空
	Token             string `url:"token"`             // 留空
	Timestamp         string `url:"timestamp"`         // 留空
	UserId            string `url:"userid"`            // 用户ID
	LicenseNo         string `url:"licenseno"`         // 车牌号
	EngineNo          string `url:"engineno"`          // 发动机编号
	CarTypeCode       string `url:"cartypecode"`       // 号牌类型 02
	VehicleType       string `url:"vehicletype"`       // 机动车类型 11微型客车 12小型汽车
	DrivingPhoto      string `url:"drivingphoto"`      // 车辆行驶证 base64(照片）
	CarPhoto          string `url:"carphoto"`          // 车辆正面照 base64(照片）
	DriverName        string `url:"drivername"`        // 驾驶人姓名
	DriverLicenseno   string `url:"driverlicenseno"`   // 驾驶人驾照编号(身份证号)
	DriverPhoto       string `url:"driverphoto"`       // 驾驶人证件 base64(照片）
	Personphoto       string `url:"personphoto"`       // 驾驶员手持身份证 base64(照片）
	GpsLon            string `url:"gpslon"`            // 地理位置信息
	GpsLat            string `url:"gpslat"`            // 地理位置信息
	PhoneNo           string `url:"phoneno"`           // 留空
	Imei              string `url:"imei"`              // 留空
	Imsi              string `url:"imsi"`              // 留空
	CarId             string `url:"carid"`             // 车辆注册编号
	CarModel          string `url:"carmodel"`          // 车辆型号
	CarRegTime        string `url:"carregtime"`        // 车辆注册时间
	EnvGrade          string `url:"envGrade"`          // 环保标准
	Code              string `url:"code"`              // 留空
	Sign              string `url:"sign"`              // 需要从客户端获取
}

/*
CarTypeCode
<li data="02">小型汽车</li>
<li data="03">使馆汽车</li>
<li data="04">领馆汽车</li>
<li data="05">境外汽车</li>
<li data="06">外国汽车</li>
<li data="26">香港入出境车</li>
<li data="27">澳门入出境车</li>
 */

// Sign 需要通过客户端获取
// var cookieval = getCookie("enterbj_sign_session_id");
// bridgewai.callHandler('loadingCameraImage', {'imageId':cookieval}, function(response) {
// 	$('#cookieval').val(response.imageString);
// }
