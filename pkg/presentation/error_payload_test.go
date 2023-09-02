package presentation

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ErrorPayloadTS struct {
	suite.Suite
}

func (suite *ErrorPayloadTS) TestErrorPayload_Equality() {

	errMsg := errors.New("Whoops, you slipped on a chip!")
	var err1 Problem7807
	err1.Status = http.StatusNotFound
	err1.Type_ = "about:blank"
	err1.Instance = ""
	err1.Detail = errMsg.Error()
	err1.Title = http.StatusText(http.StatusNotFound)

	err2 := GetFormattedErrorMessage(errMsg, http.StatusNotFound)

	suite.True(err1.Equals(err2))
	suite.True(err1.Equals(err1))
	suite.True(err2.Equals(err2))
	suite.True(err2.Equals(err1))

}

func (suite *ErrorPayloadTS) TestErrorPayload_Nil() {

	errMsg := errors.New("unknown error - could not parse from error object")
	var err1 Problem7807
	err1.Status = http.StatusNotFound
	err1.Type_ = "about:blank"
	err1.Instance = ""
	err1.Detail = errMsg.Error()
	err1.Title = http.StatusText(http.StatusNotFound)

	err2 := GetFormattedErrorMessage(nil, http.StatusNotFound)

	suite.True(err1.Equals(err2))
	suite.True(err1.Equals(err1))
	suite.True(err2.Equals(err2))
	suite.True(err2.Equals(err1))

}

func TestModelErrorPayloadSuite(t *testing.T) {

	suite.Run(t, new(ErrorPayloadTS))
}
