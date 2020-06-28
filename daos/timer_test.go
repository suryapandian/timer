package daos

import (
	"testing"

	"timer.com/db"
	"timer.com/dtos"
	"timer.com/utils"

	"github.com/stretchr/testify/suite"
)

type TimerTestSuite struct {
	suite.Suite
	dao *Timer
}

func (t *TimerTestSuite) SetupTest() {
	t.dao = NewTimer(db.GetDB())
}

func (t *TimerTestSuite) TestSave() {
	timer := &dtos.Timer{}
	timer.ID = utils.NewUUID()
	id, err := t.dao.Insert(timer)
	t.NoError(err)
	t.NotNil(id)
}

func TestTimerSuite(t *testing.T) {
	suite.Run(t, new(TimerTestSuite))
}
