package appController

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gofrendi/structureExample/appModel"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4/middleware"
)

func TestPersonLogin(t *testing.T) {
	e, pc := initTestEcho()
	person1 := appModel.Person{Name: "dono", Email: "dono@warkop.id", Password: "rahasia"}
	person1.ID = uint(1)
	pc.model.Add(person1)
	person2 := appModel.Person{Name: "kasino", Email: "kasino@warkop.id", Password: "rahasia"}
	person2.ID = uint(2)
	pc.model.Add(person2)

	// login request
	loginInfo, err := json.Marshal(LoginInfo{Email: "dono@warkop.id", Password: "rahasia"})
	if err != nil {
		t.Errorf("marshalling new person failed")
	}
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")

	// send request
	if err := pc.Login(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if rec.Code != 200 {
		t.Errorf("should return 200, get: %d", rec.Code)
	}

	// compare response
	var p appModel.Person
	if err := json.Unmarshal(rec.Body.Bytes(), &p); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if p.Token == "" {
		t.Errorf("token expected")
	}
}

func TestGetAll(t *testing.T) {
	e, pc := initTestEcho()
	person1 := appModel.Person{Name: "dono", Email: "dono@warkop.id", Password: "rahasia"}
	person1.ID = uint(1)
	pc.model.Add(person1)
	person2 := appModel.Person{Name: "kasino", Email: "kasino@warkop.id", Password: "rahasia"}
	person2.ID = uint(2)
	pc.model.Add(person2)

	// login request
	loginInfo, err := json.Marshal(LoginInfo{Email: "dono@warkop.id", Password: "rahasia"})
	if err != nil {
		t.Errorf("marshalling new person failed")
	}
	loginReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRec := httptest.NewRecorder()
	loginContext := e.NewContext(loginReq, loginRec)
	loginContext.SetPath("/login")

	// send request
	if err := pc.Login(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginRec.Code != 200 {
		t.Errorf("should return 200, get: %d", loginRec.Code)
	}

	// compare response
	var p appModel.Person
	if err := json.Unmarshal(loginRec.Body.Bytes(), &p); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if p.Token == "" {
		t.Errorf("token expected")
	}
	token := p.Token

	// get all request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/persons")

	// get all
	if err := middleware.JWT([]byte(pc.jwtSecret))(pc.GetAll)(context); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if rec.Code != 200 {
		t.Errorf("should return 200, get: %d", rec.Code)
	}

	var pList []appModel.Person
	if err := json.Unmarshal(rec.Body.Bytes(), &pList); err != nil {
		t.Errorf("unmarshalling returned person list failed")
	}

	expectedPListLength := 2
	if len(pList) != expectedPListLength {
		t.Errorf("expecting pList's length to be %d, get %d, data: %#v", expectedPListLength, len(pList), pList)
	}
}
