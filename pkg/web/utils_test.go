package web

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNewSimpleResponse(t *testing.T) {
	// create a new echo instance
	e := echo.New()
	// create a new http request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// create a new http response recorder
	rec := httptest.NewRecorder()

	_ = NewSimpleResponse(e.NewContext(req, rec), http.StatusOK, "test message")

	expected := `{"code":200,"message":"test message"}`
	if expected != strings.TrimRight(rec.Body.String(), "\n") {
		t.Errorf(EXPECT_CONTENT, expected, rec.Body.String())
	}
}

func TestNewError(t *testing.T) {
	// create a new echo instance
	e := echo.New()
	// create a new http request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// create a new http response recorder
	rec := httptest.NewRecorder()

	err := errors.New("test error")
	_ = NewError(e.NewContext(req, rec), http.StatusInternalServerError, err)

	expectedCode := http.StatusInternalServerError
	if expectedCode != rec.Code {
		t.Errorf("expected %d but got %d", expectedCode, rec.Code)
	}

	expectedBody := `{"code":500,"message":"test error"}`
	if expectedBody != strings.TrimRight(rec.Body.String(), "\n") {
		t.Errorf(EXPECT_CONTENT, expectedBody, rec.Body.String())
	}
}
