package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"timer.com/utils"

	"github.com/stretchr/testify/suite"
)

type TimerTestSuite struct {
	suite.Suite
}

func TestTimerSuite(t *testing.T) {
	suite.Run(t, new(TimerTestSuite))
}

func (t *TimerTestSuite) TestCreateInvalidTimer() {
	request := httptest.NewRequest(http.MethodPost, "/create?stepTime=7", nil)
	w := httptest.NewRecorder()
	CreateTimer(w, request)
	response := w.Result()
	t.Equal(http.StatusBadRequest, response.StatusCode)
}

func (t *TimerTestSuite) TestCreateValidTimer() {
	request := httptest.NewRequest(http.MethodPost, "/create?startVal=10&stepTime=7", nil)
	w := httptest.NewRecorder()
	CreateTimer(w, request)
	response := w.Result()
	t.Equal(http.StatusOK, response.StatusCode)
}

func (t *TimerTestSuite) TestGetAllTimer() {
	r := httptest.NewRequest(http.MethodGet, "/check", nil)
	w := httptest.NewRecorder()
	CheckTimer(w, r)
	response := w.Result()
	t.Equal(http.StatusOK, response.StatusCode)
}

func (t *TimerTestSuite) TestInvalidGetTimerByID() {
	r := httptest.NewRequest(http.MethodGet, "/check?id="+utils.NewUUID(), nil)
	w := httptest.NewRecorder()
	CheckTimer(w, r)
	response := w.Result()
	t.Equal(http.StatusNotFound, response.StatusCode)
}
