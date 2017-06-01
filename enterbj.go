package enterbj

import (
	"bytes"
	"encoding/json"
	"github.com/amlun/enterbj/response"
	"github.com/google/go-querystring/query"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	CARLIST_URL     = "https://api.jinjingzheng.zhongchebaolian.com/enterbj/platform/enterbj/entercarlist"
	LOGIN_URL       = "https://api.accident.zhongchebaolian.com/industryguild_mobile_standard_self2.1.2/mobile/standard/login"
	//SUBMITPAPER_URL = "https://api.jinjingzheng.zhongchebaolian.com/enterbj/platform/enterbj/submitpaper"
	PERSONINFO_URL  = "https://api.accident.zhongchebaolian.com/industryguild_mobile_standard_self2.1.2/mobile/standard/getpersonalinfor?"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

type Client struct {
	session *Session
	app     *App
}

func (e *Client) Verify(phone string) (*response.Verify, error) {
	reqBody := e.verifyRequest(phone)
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
	r, err := json.Marshal(reqBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("login request body is [%s]", string(r))
	req, _ := http.NewRequest("POST", LOGIN_URL, bytes.NewBuffer(r))
	req.Header = commonHeader
	res, err := http.DefaultClient.Do(req)
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

func (e *Client) GetPersonInfo(userid string) (*response.PersonInfo, error) {
	reqBody := e.personInfoRequest(userid)
	r, err := query.Values(reqBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("get person info request body is [%s]", r.Encode())
	req, _ := http.NewRequest("GET", PERSONINFO_URL+r.Encode(), nil)
	req.Header = commonHeader
	res, err := http.DefaultClient.Do(req)
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

func (e *Client) CarList() (*response.CarList, error) {
	reqBody := e.carListRequest()
	r, err := query.Values(reqBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("car list request body is [%s]", r.Encode())
	req, _ := http.NewRequest("POST", CARLIST_URL, bytes.NewBufferString(r.Encode()))
	req.Header = commonHeader
	res, err := http.DefaultClient.Do(req)
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

func (e *Client) SubmitPaper() {

}
