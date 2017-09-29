package main

import (
	"github.com/amlun/enterbj"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gcfg.v1"
	"os"
)

var conf *Config
var eClient *enterbj.Client

type Test struct {
	UserId string
}

type Config struct {
	EnterBj enterbj.Config
	Test    Test
}

func InitConfig(confPath string) (*Config, error) {
	if conf != nil {
		return conf, nil
	}
	conf = &Config{}
	if err := gcfg.ReadFileInto(conf, confPath); err != nil {
		return nil, err
	}
	return conf, nil
}

func main() {
	// 必须指定配置文件
	if len(os.Args) < 2 {
		log.Error("You must specify a config file")
		return
	}
	_, err := InitConfig(os.Args[1])
	if err != nil {
		log.Error("Init config file error", err)
		return
	}
	eClient = enterbj.New(&conf.EnterBj)
	checkStatus()

	// TODO
	//c := cron.New()
	//c.AddFunc("@daily", checkCar)
	//c.Run()
}

func checkStatus() {
	if err := eClient.CheckServiceStatus(); err != nil {
		log.Error("当前服务不可用")
	} else {
		log.Info("当前服务可用，请尽快处理")
	}
}

// TODO
// 每天运行一次，检查是否过期
func checkCar() {
	if info, err := eClient.CarList(conf.Test.UserId); err != nil {
		log.Error("Get car list error", err)
	} else {
		for _, car := range info.DataList {
			if car.ApplyFlag == "1" {
				// TODO send notice
				log.Warnf("该车辆 %s 当前可以申请，请立即申请！", car.LicenseNo)
				//eClient.SubmitPaper(conf.Test.UserId, car.LicenseNo, car.)
			} else {
				var format string
				for _, apply := range car.CarApplyArr {
					switch apply.Status {
					case "1":
						format = "车辆 %s 申请成功，时间为 %s 到 %s"
					case "2":
						format = "车辆 %s 正在审核，时间为 %s 到 %s"
					default:
						format = "车辆 %s 审核失败，时间为 %s 到 %s"
					}
					log.Infof(format, apply.LicenseNo, apply.EnterBjStart, apply.EnterBjEnd)
				}
			}
		}
	}
}
