package db

import (
	"database/sql"
	"fmt"

	"timer.com/config"
	"timer.com/utils"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func init() {
	l := utils.LogEntryWithRef()
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s  dbname=%s sslmode=disable", config.DB_HOST, config.DB_NAME))
	if err != nil {
		l.WithError(err).Fatal("Failed to get DB connection")
	}

	db.SetMaxIdleConns(config.MAX_DB_CONNECTIONS)
	db.SetMaxOpenConns(config.MAX_DB_CONNECTIONS)
	logrus.Info("Successfully established database connection")

}

func GetDB() *sql.DB {
	return db
}
