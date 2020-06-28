package daos

import (
	"database/sql"

	"timer.com/dtos"
)

type Timer struct {
	dbConn *sql.DB
}

func NewTimer(dbConn *sql.DB) *Timer {
	return &Timer{dbConn: dbConn}
}

func (t *Timer) Insert(timer *dtos.Timer) (id string, err error) {
	statement, err := t.dbConn.Prepare(`INSERT INTO timers (id, step_time, counter, created_at, updated_at)  VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return "", err
	}

	defer statement.Close()
	_, err = statement.Exec(timer.ID, timer.StepTime, timer.Counter, timer.CreatedAt, timer.ModifiedAt)
	return id, err
}

func (t *Timer) Update(timer *dtos.Timer) (err error) {
	statement, err := t.dbConn.Prepare(`UPDATE timers SET counter=$1, updated_at = $2 WHERE id=$3`)
	if err != nil {
		return
	}

	defer statement.Close()
	_, err = statement.Exec(timer.Counter, timer.ModifiedAt, timer.ID)
	return err
}

func (t *Timer) Delete(id string) (err error) {
	statement, err := t.dbConn.Prepare(`DELETE FROM timers WHERE id=$1`)
	if err != nil {
		return
	}

	defer statement.Close()
	_, err = statement.Exec(id)

	return err
}
