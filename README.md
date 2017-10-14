# enterbj

## 公告

**本应用为测试项目，请勿使用在线上**

## 应用介绍

辅助办理进京证，可以接入第三方API做到消息通知等

## 使用说明

### 准备

首先先安装[Go](https://golang.org/)并配置好环境变量 $GOROOT 和 $GOPATH

然后安装Go的包管理[Glide](https://glide.sh/)

之后就checkout代码

```bash
mkdir -p $GOPATH/src/github.com/amlun
git clone https://github.com/amlun/enterbj $GOPATH/src/github.com/amlun/enterbj
cd $GOPATH/src/github.com/amlun/enterbj
glide install
cp config.ini.example config.ini
```

### 配置

**目前SIGN处于测试中，还不对外开放**

修改userId为你自己的userId 😼

配置你要接收信息的邮箱地址 email 📮

然后配置邮件服务 [mail]

修改signUrl 😊

```ini
[test]
userId = ABCDEFGHIJKLMNOPQRSTUVWXYZ
email  = user@example.com

[enterbj]
appKey = kkk
appSource = bjjj
signUrl = http://127.0.0.1:8080/gensign/%s%s
timeOut = 2000

[mail]
userName = user@example.com
passWord = password
smtpHost = smtp.example.com
smtpPort = 25

```

### 运行

```bash
go run main.go config.ini
```
或者
```bash
make build
./enterbj config.ini
```

### 运行结果

车辆已有进京证：
```
INFO[0001] 车辆 冀A66666 已经申请到进京证，时间为 2018-01-01 到 2018-01-07 
```

车辆可以申请进京证：
```                        
WARN[0001] 该车辆 冀A88888 当前可以申请，请立即申请！
```

服务可用
```
INFO[0001] 当前服务可用，请尽快处理
```

服务不可用
```
ERRO[0001] 当前服务不可用
```

邮件提醒
```
INFO[0001] sendMail(进京证办理服务检查, 当前服务不可用)
INFO[0001] sendMail(进京证办理服务检查, 当前服务可用，请尽快处理)
INFO[0001] sendMail(进京证办理提醒, 该车辆 冀A66666 当前可以申请，请立即申请！)
```

## 版本记录

### 当前开发版本

完成基本接口的对接：

- [x] 获取用户信息
- [x] 获取车辆列表
- [x] 获取车辆环保信息
- [x] 服务可用状态
- [ ] 登录
- [ ] 验证码
- [ ] 获取其他驾驶人员列表
- [ ] 添加车辆
- [ ] 添加其他驾驶员
- [ ] 提交申请

## 其它

[讨论](https://github.com/amlun/enterbj/issues)

[群聊](https://t.me/joinchat/F9pB0w5VRUE3TC-pB5w_NQ)

[公告](https://t.me/enterbj)

