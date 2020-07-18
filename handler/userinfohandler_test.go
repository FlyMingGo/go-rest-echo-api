package handler

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"testproject/model"
)

func TestApiGetUserInfo(t *testing.T) {
	jsonTest := model.UserInfo{Id: 1, Name: "Frank", Age: 36, City: "Burnaby"}
	var jsonReturn model.UserInfo
	e := echo.New()
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	if assert.NoError(t, ApiGetUserInfo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		err := json.Unmarshal(rec.Body.Bytes(), &jsonReturn)
		if err != nil {
			t.Error(err)
			return
		}
		assert.Equal(t, jsonTest, jsonReturn)
	}
}

func TestApiGetAllUserInfo(t *testing.T) {
	var jsonTest []model.UserInfo
	jsonTest = append(jsonTest, model.UserInfo{Id: 1, Name: "Frank", Age: 36, City: "Burnaby"})
	jsonTest = append(jsonTest, model.UserInfo{Id: 2, Name: "Tina", Age: 31, City: "Port Moody"})
	jsonTest = append(jsonTest, model.UserInfo{Id: 3, Name: "Hope", Age: 5, City: "Vancouver"})
	jsonTest = append(jsonTest, model.UserInfo{Id: 4, Name: "Tom", Age: 18, City: "White Rock"})
	var jsonReturn []model.UserInfo
	e := echo.New()
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	if assert.NoError(t, ApiGetAllUserInfo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		err := json.Unmarshal(rec.Body.Bytes(), &jsonReturn)
		if err != nil {
			t.Error(err)
			return
		}
		assert.Equal(t, jsonTest, jsonReturn)
	}
}

func TestApiCreatUser(t *testing.T) {
	e := echo.New()
	var jsonStr = []byte(`{"name":"Ian","age":"15","city":"New York"}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, ApiCreatUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestApiUpdateUser(t *testing.T) {
	e := echo.New()
	var jsonStr = []byte(`{"id":"7","name":"Ian","age":"15","city":"Victoria"}`)
	req := httptest.NewRequest("PUT", "/", bytes.NewBuffer(jsonStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, ApiUpdateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestApiDeleteUserInfo(t *testing.T) {
	jsonTest := model.UserInfo{Id: 8, Name: "Ian", Age: 15, City: "Victoria"}
	var jsonReturn model.UserInfo
	e := echo.New()
	req := httptest.NewRequest("DELETE", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("8")
	if assert.NoError(t, ApiGetUserInfo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		err := json.Unmarshal(rec.Body.Bytes(), &jsonReturn)
		if err != nil {
			t.Error(err)
			return
		}
		assert.Equal(t, jsonTest, jsonReturn)
	}
}
