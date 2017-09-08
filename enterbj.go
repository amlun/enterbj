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

const (
	DEFAULT_TIME_OUT    = 100
	CARLIST_URL         = "https://api.jinjingzheng.zhongchebaolian.com/enterbj/platform/enterbj/entercarlist"
	LOGIN_URL           = "https://bjjj.zhongchebaolian.com/industryguild_mobile_standard_self2.1.2/mobile/standard/login"
	SUBMIT_PAPER_URL    = "https://api.jinjingzheng.zhongchebaolian.com/enterbj/platform/enterbj/submitpaper"
	PERSON_INFO_URL     = "https://api.accident.zhongchebaolian.com/industryguild_mobile_standard_self2.1.2/mobile/standard/getpersonalinfor?"
	CHECK_ENV_GRADE_URL = "https://api.jinjingzheng.zhongchebaolian.com/enterbj/platform/enterbj/checkenvgrade"
	//LOAD_OTHER_DRIVERS_URL = "https://api.jinjingzheng.zhongchebaolian.com/enterbj/platform/enterbj/loadotherdrivers"
)

var (
	initialized bool
	client      *Client
	conf        *Config
	httpClient  *http.Client
	mutex       sync.Mutex
)

type Client struct {
}

// New 返回一个 Enterbj Client
func New(config *Config) *Client {
	mutex.Lock()
	defer mutex.Unlock()

	if initialized {
		return client
	}
	// 默认100ms
	if config.Timeout == 0 {
		config.Timeout = DEFAULT_TIME_OUT
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

func (e *Client) GetPersonInfo(userId string) (*response.PersonInfo, error) {
	req := personInfoRequest(userId)
	if req == nil {
		return nil, errors.New("generate person info request with error")
	}
	var repBody *response.PersonInfo
	if _, err := sendRequest(req, &repBody); err != nil {
		return nil, err
	}
	return repBody, nil
}

func (e *Client) CarList(userId string) (*response.CarList, error) {
	req := carListRequest(userId)
	if req == nil {
		return nil, errors.New("generate car list request with error")
	}

	var repBody *response.CarList
	if _, err := sendRequest(req, &repBody); err != nil {
		return nil, err
	}
	return repBody, nil

}

func (e *Client) CheckEnvGrade(userId, carId, licenseNo, carModel, carRegTime string) (*response.CheckEnvGrade, error) {
	req := checkEnvGradeRequest(userId, carId, licenseNo, carModel, carRegTime)
	if req == nil {
		return nil, errors.New("generate check env grade request with error")
	}

	var repBody *response.CheckEnvGrade
	if _, err := sendRequest(req, &repBody); err != nil {
		return nil, err
	}
	return repBody, nil
}

// TODO 获取cookie
func (e *Client) LoadOtherDrivers() error {
	return nil
}

// TODO 处理参数,sign需要解决，通过客户端处理
func (e *Client) SubmitPaper(userId, licenseNo, engineNo, carTypeCode string) (*response.SubmitPaper, error) {
	req := applySubmitRequest(userId, licenseNo, engineNo, carTypeCode)
	if req == nil {
		return nil, errors.New("generate submit paper request with error")
	}

	var repBody *response.SubmitPaper
	if _, err := sendRequest(req, &repBody); err != nil {
		return nil, err
	}
	return repBody, nil
}
