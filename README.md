# enterbj

## 公告

**本应用为测试项目，请勿使用在线上**

## 应用介绍

辅助办理进京证，可以接入第三方API做到消息通知等

## 使用说明

### 准备

安装 [Glide](https://glide.sh/)

```bash
mkdir -p $GOPATH/src/github.com/amlun
git clone https://github.com/amlun/enterbj $GOPATH/src/github.com/amlun
cd $GOPATH/src/github.com/amlun/enterbj
glide install
cd example
cp config.ini.example config.ini
```

### 配置

**目前SIGN处于测试中，还不对外开放**
修改config.ini文件的signUrl，（如果你可以找到的话😊）

```ini
[test]
userId = ABCDEFGHIJKLMNOPQRSTUVWXYZ

[enterbj]
appKey = kkk
appSource = bjjj
signUrl = http://127.0.0.1:8080/gensign/%s%s
timeOut = 2000

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
INFO[0002] 车辆 冀A66666 已经申请到进京证，时间为 2018-01-01 到 2018-01-07 
```

车辆可以申请进京证：
```
WARN[0002] 该车辆 冀A88888 当前可以申请，请立即申请！
```

## 版本记录

### 当前开发版本

完成基本接口的对接：

- [x] 获取用户信息
- [x] 获取车辆列表
- [x] 获取车辆环保信息
- [ ] 登录
- [ ] 验证码
- [ ] 获取其他驾驶人员列表
- [ ] 添加车辆
- [ ] 添加其他驾驶员
- [ ] 提交申请
