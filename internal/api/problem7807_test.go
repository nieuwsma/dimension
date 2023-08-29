package api

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Problem7807TS struct {
	suite.Suite
}

func (suite *Problem7807TS) TestProblem7807Equals_HappyPath() {

	var err1 Problem7807
	var err2 Problem7807

	err1 = GetFormattedErrorMessage(nil, 500)
	err2 = GetFormattedErrorMessage(nil, 500)

	suite.True(err1.Equals(err2))
	suite.True(err2.Equals(err1))
	suite.True(err1.Equals(err1))
	suite.True(err2.Equals(err2))

	errMsg := errors.New("Whoops, you slipped on a chip!")

	err1 = GetFormattedErrorMessage(errMsg, 500)
	err2 = GetFormattedErrorMessage(errMsg, 500)

	suite.True(err1.Equals(err2))
	suite.True(err2.Equals(err1))
	suite.True(err1.Equals(err1))
	suite.True(err2.Equals(err2))

}

func (suite *Problem7807TS) TestProblem7807Equals_NotEqual() {

	var err1 Problem7807
	var err2 Problem7807
	errMsg := errors.New("Whoops, you slipped on a chip!")

	err1 = GetFormattedErrorMessage(errMsg, 500)
	err2 = GetFormattedErrorMessage(nil, 500)

	suite.False(err1.Equals(err2))
	suite.False(err2.Equals(err1))
	suite.True(err1.Equals(err1))
	suite.True(err2.Equals(err2))

}

func TestModelProblem7807Suite(t *testing.T) {

	suite.Run(t, new(Problem7807TS))
}
