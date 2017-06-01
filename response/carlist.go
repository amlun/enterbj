package response

type CarList struct {
	DataList []Car `json:"datalist"`
	Res
}

//status = 0 //审核失败
//status = 1 //审核成功
//status = 2 //审核中

type CarApply struct {
	ApplyId         string `json:"applyid"`
	CarId           string `json:"carid"`
	CarType         string `json:"cartype"`
	EngineNo        string `json:"engineno"`
	EnterBjEnd      string `json:"enterbjend"`
	EnterBjStart    string `json:"enterbjstart"`
	ExistPaper      string `json:"existpaper"`
	LicenseNo       string `json:"licenseno"`
	LoadPaperMethod string `json:"loadpapermethod"`
	Remark          string `json:"remark"`
	Status          string `json:"status"`
	SysCode         string `json:"syscode"`
	SysCodeDesc     string `json:"syscodedesc"`
	UserId          string `json:"userid"`
}

//applyflag = 0 //不能申请
//applyflag = 1 //可以申请
type Car struct {
	CarId       string     `json:"carid"`
	UserId      string     `json:"userid"`
	LicenseNo   string     `json:"licenseno"`
	ApplyFlag   string     `json:"applyflag"`
	ApplyId     string     `json:"applyid"`
	CarApplyArr []CarApply `json:"carapplyarr"`
}
