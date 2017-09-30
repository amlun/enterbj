package main

import (
	"fmt"
	"github.com/amlun/enterbj"
	"github.com/jordan-wright/email"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"gopkg.in/gcfg.v1"
	"net/smtp"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	conf        *Config
	eClient     *enterbj.Client
	smtpAuth    smtp.Auth
	mailMutex   sync.Mutex
	statusMutxt sync.Mutex
	checkMutex  sync.Mutex
	eMail       = email.NewEmail()
	check       bool
	lastCheck   = LastCheck{}
)

type LastCheck struct {
	Online   bool
	lastTime int64
}

type Test struct {
	UserId string
	Email  string
}

type MailConfig struct {
	UserName string
	PassWord string
	SmtpHost string
	SmtpPort string
}

type Config struct {
	EnterBj enterbj.Config
	Test    Test
	Mail    MailConfig
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
	logrus.Info("begin ...")
	// 必须指定配置文件
	if len(os.Args) < 2 {
		logrus.Error("You must specify a config file")
		return
	}
	// 读取配置信息
	_, err := InitConfig(os.Args[1])
	if err != nil {
		logrus.Error("Init config file error", err)
		return
	}
	// 监听信号
	quit := signals()
	// 初始化enterbj客户端
	eClient = enterbj.New(&conf.EnterBj)
	// smtp服务起认证
	smtpAuth = smtp.PlainAuth("", conf.Mail.UserName, conf.Mail.PassWord, conf.Mail.SmtpHost)
	// 邮箱客户端初始化配置
	eMail.From = fmt.Sprintf("Enterbj Notice <%s>", conf.Mail.UserName)
	eMail.To = []string{conf.Test.Email}
	checkServiceStatus()
	// 定时任务
	c := cron.New()
	c.AddFunc("@daily", checkCar)
	c.AddFunc("@hourly", checkServiceStatus)
	logrus.Info("cron start ...")
	c.Start()
	defer func() {
		c.Stop()
		logrus.Info("cron stop ...")
	}()
	<-quit
	logrus.Info("quit ...")
}

func sendMail(subject, text string) {
	mailMutex.Lock()
	defer mailMutex.Unlock()

	logrus.Infof("sendMail(%s, %s)", subject, text)
	eMail.Subject = subject
	eMail.Text = []byte(text)
	if err := eMail.Send(conf.Mail.SmtpHost+":"+conf.Mail.SmtpPort, smtpAuth); err != nil {
		logrus.Error(err)
	}
}

func checkServiceStatus() {
	statusMutxt.Lock()
	defer statusMutxt.Unlock()

	// 不需要检查
	if !check {
		return
	}

	if err := eClient.CheckServiceStatus(); err != nil {
		if lastCheck.Online || lastCheck.lastTime == 0 {
			sendMail("进京证办理服务检查", "当前服务不可用")
		}
		lastCheck.Online = false
		logrus.Error("当前服务不可用")
	} else {
		if !lastCheck.Online || lastCheck.lastTime == 0 {
			sendMail("进京证办理服务检查", "当前服务可用，请尽快处理")
		}
		lastCheck.Online = true
		logrus.Info("当前服务可用，请尽快处理")
	}
	lastCheck.lastTime = time.Now().Unix()
}

// TODO
// 每天运行一次，检查是否过期
func checkCar() {
	checkMutex.Lock()
	defer checkMutex.Unlock()

	if info, err := eClient.CarList(conf.Test.UserId); err != nil {
		logrus.Error("Get car list error", err)
	} else {
		for _, car := range info.DataList {
			if car.ApplyFlag == "1" {
				check = true
				text := fmt.Sprintf("该车辆 %s 当前可以申请，请立即申请！", car.LicenseNo)
				logrus.Warn(text)
				sendMail("进京证办理提醒", text)
				// TODO 自动申请
				//eClient.SubmitPaper(conf.Test.UserId, car.LicenseNo, car.)
			} else {
				check = false
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
					logrus.Infof(format, apply.LicenseNo, apply.EnterBjStart, apply.EnterBjEnd)
				}
			}
		}
	}
}

// Signal Handling
func signals() <-chan bool {
	quit := make(chan bool)
	go func() {
		signals := make(chan os.Signal)
		defer close(signals)
		signal.Notify(signals, syscall.SIGQUIT, syscall.SIGTERM, os.Interrupt)
		defer signalStop(signals)
		<-signals
		quit <- true
	}()
	return quit
}

// Stops signals channel.
func signalStop(c chan<- os.Signal) {
	signal.Stop(c)
}
