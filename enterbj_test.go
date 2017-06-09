package enterbj

import (
	"testing"
)

func TestEnterBJ_CarList(t *testing.T) {
	r, _ := ob.CarList()
	t.Log(r)
}

func TestClient_Login(t *testing.T) {
	r, _ := ob.Login("18688888888", "666666")
	t.Log(r)
}

func TestClient_GetPersonInfo(t *testing.T) {
	r, _ := ob.GetPersonInfo()
	t.Log(r)
}

func TestClient_CheckEnvGrade(t *testing.T) {
	r, _ := ob.CheckEnvGrade(carId, LicenseNo, carModel, carRegTime)
	t.Log(r)
}
