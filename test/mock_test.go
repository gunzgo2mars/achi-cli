package mock_test

import (
	"testing"

	mock "github.com/gunzgo2mars/achi-cli/test"
	"github.com/stretchr/testify/suite"
)

type SumIntTestSuite struct {
	suite.Suite
}

func TestSumIntTestSuite(t *testing.T) {
	suite.Run(t, new(SumIntTestSuite))
}

func (ts *SumIntTestSuite) TestSumInt() {

	ts.T().Run("test sum int", func(t *testing.T) {

		actual := mock.SumInt(1, 3)
		ts.Assert().Equal(4, actual, "actual should be 4 equal to expected")

	})
}
