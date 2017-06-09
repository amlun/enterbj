package response

// envGrade == 1 黄标车
// envGrade == 2 国1国2
// envGrade == 3 其它

type CheckEnvGrade struct {
	EnvGrade int `json:"envgrade"`
	Res
}
