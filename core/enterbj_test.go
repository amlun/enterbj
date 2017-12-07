package core

import (
	"testing"
)

func TestClient_CarList(t *testing.T) {
	setUpTest()
	if r, err := ob.CarList(userID); err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestClient_Login(t *testing.T) {
	setUpTest()
	if r, err := ob.Login(phone, code); err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestClient_GetPersonInfo(t *testing.T) {
	setUpTest()
	if r, err := ob.GetPersonInfo(userID); err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestClient_CheckEnvGrade(t *testing.T) {
	setUpTest()
	if r, err := ob.CheckEnvGrade(userID, carID, licenseNo, carModel, carRegTime); err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestClient_SubmitPaper(t *testing.T) {
	setUpTest()
	if r, err := ob.SubmitPaper(userID, licenseNo, engineNo, drivingPhotoPath, carPhotoPath, driverName, driverLicenseNo,
		driverPhotoPath, personPhotoPath, carID, carModel, carRegTime, envGrade); err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestClient_CheckServiceStatus(t *testing.T) {
	setUpTest()
	if err := ob.CheckServiceStatus(userID); err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}
