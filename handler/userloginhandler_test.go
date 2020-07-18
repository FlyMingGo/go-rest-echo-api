package handler

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiLogin(t *testing.T) {
	e := echo.New()
	var jsonStr = []byte(`{"username":"admin","password":"admin"}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, ApiLogin(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
