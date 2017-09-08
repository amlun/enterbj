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
	checkCar()

	// TODO
	//c := cron.New()
	//c.AddFunc("@daily", checkCar)
	//c.Run()
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
			} else {
				for _, apply := range car.CarApplyArr {
					log.Infof("车辆 %s 已经申请到进京证，时间为 %s 到 %s", apply.LicenseNo, apply.EnterBjStart, apply.EnterBjEnd)
				}
			}
		}
	}
}
