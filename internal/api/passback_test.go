package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PassbackTS struct {
	suite.Suite
}

func (suite *PassbackTS) TestBuildErrorPassback() {
	var err error
	err = fmt.Errorf("ERROR: TestBuildError")
	pb := BuildErrorPassback(1200, err)
	suite.True(pb.IsError)
	suite.True(pb.StatusCode == 1200)
}

func (suite *PassbackTS) TestBuildSuccessPassback() {
	obj := []string{"a1", "a2", "a3"}
	pb := BuildSuccessPassback(200, obj)
	suite.False(pb.IsError)
	suite.True(pb.StatusCode == 200)
}

func TestPassbackSuite(t *testing.T) {

	suite.Run(t, new(PassbackTS))
}
