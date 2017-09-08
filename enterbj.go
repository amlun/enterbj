package enterbj

import (
	"encoding/json"
	"errors"
	"github.com/amlun/enterbj/response"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

// DefaultTimeOut 默认的超时时间，单位为 ms
const DefaultTimeOut = 100

var (
	initialized bool
	client      *Client
	conf        *Config
	httpClient  *http.Client
	mutex       sync.Mutex
)

// Client 进京证办理客户端
type Client struct {
}

// New 返回一个 Enterbj Client对象
func New(config *Config) *Client {
	mutex.Lock()
	defer mutex.Unlock()

	if initialized {
		return client
	}
	// 默认100ms
	if config.Timeout == 0 {
		config.Timeout = DefaultTimeOut
	}
	// http client
	httpClient = &http.Client{
		Timeout: config.Timeout * time.Millisecond,
	}
	// config
	conf = config
	// enterbj client
	client = &Client{}
	initialized = true
	return client
}

// Verify 验证手机号
func (e *Client) Verify(phone string) (*response.Verify, error) {
	reqBody := verifyRequest(phone)
	if reqBody == nil {
		return nil, errors.New("generate verify request with error")
	}
	r, err := json.Marshal(reqBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("verify request body is [%s]", string(r))

	return nil, nil
}

// Login 用户登录
func (e *Client) Login(phone string, valicode string) (*response.Login, error) {
	req := loginRequest(phone, valicode)
	if req == nil {
		return nil, errors.New("generate login request with error")
	}
	var repBody *response.Login
	if _, err := sendRequest(req, &repBody); err != nil {
		return nil, err
	}
	return repBody, nil
}

// GetPersonInfo 获取用户信息
func (e *Client) GetPersonInfo(userID string) (*response.PersonInfo, error) {
	req := personInfoRequest(userID)
	if req == nil {
		return nil, errors.New("generate person info request with error")
	}
	var repBody *response.PersonInfo
	if _, err := sendRequest(req, &repBody); err != nil {
		return nil, err
	}
	return repBody, nil
}

// CarList 获取用户的车辆列表
func (e *Client) CarList(userID string) (*response.CarList, error) {
	req := carListRequest(userID)
	if req == nil {
		return nil, errors.New("generate car list request with error")
	}

	var repBody *response.CarList
	if _, err := sendRequest(req, &repBody); err != nil {
		return nil, err
	}
	return repBody, nil

}

// CheckEnvGrade 检查用户的车辆环保信息
func (e *Client) CheckEnvGrade(userID, carID, licenseNo, carModel, carRegTime string) (*response.CheckEnvGrade, error) {
	req := checkEnvGradeRequest(userID, carID, licenseNo, carModel, carRegTime)
	if req == nil {
		return nil, errors.New("generate check env grade request with error")
	}

	var repBody *response.CheckEnvGrade
	if _, err := sendRequest(req, &repBody); err != nil {
		return nil, err
	}
	return repBody, nil
}

// LoadOtherDrivers 加载其他驾驶人信息 TODO
func (e *Client) LoadOtherDrivers() error {
	return nil
}

// SubmitPaper 提交进京证申请 TODO
func (e *Client) SubmitPaper(userID, licenseNo, engineNo, carTypeCode string) (*response.SubmitPaper, error) {
	req := applySubmitRequest(userID, licenseNo, engineNo, carTypeCode)
	if req == nil {
		return nil, errors.New("generate submit paper request with error")
	}

	var repBody *response.SubmitPaper
	if _, err := sendRequest(req, &repBody); err != nil {
		return nil, err
	}
	return repBody, nil
}
