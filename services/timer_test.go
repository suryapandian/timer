package services

import (
	"testing"

	"timer.com/dtos"
	"timer.com/utils"

	"github.com/stretchr/testify/suite"
)

type TimerTestSuite struct {
	suite.Suite
	t *Timer
}

func (t *TimerTestSuite) SetupTest() {
	t.t = NewTimer(utils.LogEntryWithRef())
}

func TestTimerSuite(t *testing.T) {
	suite.Run(t, new(TimerTestSuite))
}

func (t *TimerTestSuite) TestCreateValidTimer() {
	timer := &dtos.Timer{
		StepTime: 5,
		Counter:  10,
	}
	id := t.t.Create(timer)
	t.NotEmpty(id)
}

func (t *TimerTestSuite) TestGetTimer() {
	timer := &dtos.Timer{
		StepTime: 5,
		Counter:  10,
	}
	id := t.t.Create(timer)
	t.NotEmpty(id)

	timers := t.t.GetAll()
	t.NotEmpty(timers)

	timer, err := t.t.GetByID(id)
	t.NotEmpty(timer)
	t.NoError(err)

	//Invalid Timer
	timer, err = t.t.GetByID(utils.NewUUID())
	t.Empty(timer)
	t.Equal(ErrTimerNotFound, err)

}
