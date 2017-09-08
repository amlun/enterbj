package enterbj

import (
	"bytes"
	"encoding/json"
	"github.com/amlun/enterbj/response"
	"github.com/google/go-querystring/query"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
	"errors"
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

var client *Client

type Client struct {
	Conf *Config
	Http *http.Client
}

func New(config *Config) *Client {
	if client != nil {
		return client
	}
	// 默认100ms
	if config.Timeout == 0 {
		config.Timeout = DEFAULT_TIME_OUT
	}
	// http client
	httpClient := http.Client{
		Timeout: config.Timeout * time.Millisecond,
	}
	// enterbj client
	client = &Client{
		Conf: config,
		Http: &httpClient,
	}
	return client
}

func (e *Client) Verify(phone string) (*response.Verify, error) {
	reqBody := e.verifyRequest(phone)
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
	reqBody := e.loginRequest(phone, valicode)
	if reqBody == nil {
		return nil, errors.New("generate login request with error")
	}
	r, err := json.Marshal(reqBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("login request body is [%s]", string(r))
	req, _ := http.NewRequest("POST", LOGIN_URL, bytes.NewBuffer(r))
	req.Header = commonHeader
	res, err := e.Http.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("login response body is [%s]", body)
	var repBody response.Login
	err = json.Unmarshal(body, &repBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &repBody, nil
}

func (e *Client) GetPersonInfo(userId string) (*response.PersonInfo, error) {
	reqBody := e.personInfoRequest(userId)
	if reqBody == nil {
		return nil, errors.New("generate person info request with error")
	}
	r, err := query.Values(reqBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("get person info request body is [%s]", r.Encode())
	req, _ := http.NewRequest("GET", PERSON_INFO_URL+r.Encode(), nil)
	req.Header = commonHeader
	res, err := e.Http.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("get person info response body is [%s]", body)
	var repBody response.PersonInfo
	err = json.Unmarshal(body, &repBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &repBody, nil

}

func (e *Client) CarList(userId string) (*response.CarList, error) {
	reqBody := e.carListRequest(userId)
	if reqBody == nil {
		return nil, errors.New("generate car list request with error")
	}
	r, err := query.Values(reqBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("car list request body is [%s]", r.Encode())
	req, _ := http.NewRequest("POST", CARLIST_URL, bytes.NewBufferString(r.Encode()))
	req.Header = commonHeader
	res, err := e.Http.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("car list response body is [%s]", body)
	var repBody response.CarList
	err = json.Unmarshal(body, &repBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &repBody, nil
}

func (e *Client) CheckEnvGrade(userId, carId, licenseNo, carModel, carRegTime string) (*response.CheckEnvGrade, error) {
	reqBody := e.checkEnvGradeRequest(userId, carId, licenseNo, carModel, carRegTime)
	if reqBody == nil {
		return nil, errors.New("generate check env grade request with error")
	}
	r, err := query.Values(reqBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("check env grade request body is [%s]", r.Encode())
	req, _ := http.NewRequest("POST", CHECK_ENV_GRADE_URL, bytes.NewBufferString(r.Encode()))
	req.Header = commonHeader
	res, err := e.Http.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("check env grade response body is [%s]", body)
	var repBody response.CheckEnvGrade
	err = json.Unmarshal(body, &repBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &repBody, nil
}

// TODO 获取cookie
func (e *Client) LoadOtherDrivers() error {
	return nil
}

// TODO 处理参数,sign需要解决，通过客户端处理
func (e *Client) SubmitPaper(userId, licenseNo, engineNo, carTypeCode string) (*response.SubmitPaper, error) {
	reqBody := e.applySubmitRequest(userId, licenseNo, engineNo, carTypeCode)
	if reqBody == nil {
		return nil, errors.New("generate submit paper request with error")
	}
	r, err := query.Values(reqBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("submit paper request body is [%s]", r.Encode())
	req, _ := http.NewRequest("POST", SUBMIT_PAPER_URL, bytes.NewBufferString(r.Encode()))
	req.Header = commonHeader
	res, err := e.Http.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("submit paper response body is [%s]", body)
	var repBody response.SubmitPaper
	err = json.Unmarshal(body, &repBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &repBody, nil
}
