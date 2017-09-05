package enterbj

import (
	"testing"
)

func TestEnterBJ_CarList(t *testing.T) {
	if r, err := ob.CarList(); err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestClient_Login(t *testing.T) {
	if r, err := ob.Login(phone, code); err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestClient_GetPersonInfo(t *testing.T) {
	if r, err := ob.GetPersonInfo(); err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestClient_CheckEnvGrade(t *testing.T) {
	if r, err := ob.CheckEnvGrade(carId, licenseNo, carModel, carRegTime); err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestClient_SubmitPaper(t *testing.T) {
	if r, err := ob.SubmitPaper(licenseNo, engineNo, carTypeCode); err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}
