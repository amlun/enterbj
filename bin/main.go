package main

import (
	"github.com/robfig/cron"
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/amlun/enterbj"
	"gopkg.in/gcfg.v1"
)

var conf *Config
var eClient *enterbj.Client

type Config struct {
	EnterBj enterbj.Config
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
	eClient = &enterbj.Client{}

	// TODO
	c := cron.New()
	c.AddFunc("@daily", checkCar)
	c.Run()
}

// TODO
// 每天运行一次，检查是否过期
func checkCar() {
	if info, err := eClient.CarList(conf.EnterBj.UserId); err != nil {
		log.Error("Get car list error", err)
	} else {
		log.Info(info)
	}
}
