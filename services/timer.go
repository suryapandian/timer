package services

import (
	"context"
	"errors"
	"time"
	"timer.com/dtos"
	"timer.com/utils"

	"github.com/sirupsen/logrus"
)

type Timer struct {
	l *logrus.Entry
}

var timers = make(map[string]*dtos.Timer)

func NewTimer(l *logrus.Entry) *Timer {
	return &Timer{l: l}
}

func (t *Timer) Create(timer *dtos.Timer) string {
	ctx, cancel := context.WithCancel(context.Background())
	//timer := dtos.Timer{}
	timer.ID = utils.NewUUID()
	timer.Context = ctx
	timer.Cancel = cancel
	createdAt := time.Now().UTC()
	timer.CreatedAt = &createdAt
	timer.ModifiedAt = &createdAt
	timers[timer.ID] = timer
	t.l.WithFields(
		logrus.Fields{
			"timerID":  timer.ID,
			"stepTime": timer.StepTime,
			"counter":  timer.Counter,
		},
	).Info("starting Timer for timerID")
	go t.startTimer(timer)
	return timer.ID
}

func (t *Timer) GetByID(id string) (*dtos.Timer, error) {
	err := t.IsPresent(id)
	if err != nil {
		return nil, err
	}
	return timers[id], err
}

func (t *Timer) GetAll() (result []*dtos.Timer) {
	for _, timer := range timers {
		result = append(result, timer)
	}
	return
}

func (t *Timer) Delete(id string) error {
	t.l.Info("Deleting Timer")
	err := t.IsPresent(id)
	if err != nil {
		return err
	}

	timers[id].Cancel()
	delete(timers, id)
	return nil

}

func (t *Timer) Pause(id string) (*dtos.Timer, error) {
	t.l.Info("Pausing Timer")
	err := t.IsPresent(id)
	if err != nil {
		return nil, err
	}
	timers[id].Cancel()
	modifiedAt := time.Now().UTC()
	timers[id].ModifiedAt = &modifiedAt
	return timers[id], err
}

var ErrTimerNotFound = errors.New("Timer Not found")

func (t *Timer) IsPresent(id string) error {
	_, ok := timers[id]
	if !ok {
		return ErrTimerNotFound
	}
	return nil
}

func (t *Timer) startTimer(timer *dtos.Timer) {
	for true {
		t.l.Info("Inside timer - Counter: ", timer.Counter)
		timer.Counter = timer.Counter + 1
		time.Sleep(time.Duration(timer.StepTime) * time.Second)
		manualBreak := timer.Counter > timer.StepTime*5
		select {
		case <-timer.Context.Done():
			t.l.Info("Break through code")
			return
		default:
			if manualBreak {
				t.l.Info("Break Maually")
				return
			}
		}

	}
}
