package utils

import (
	"math/rand"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func SetupLog() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(TimestampFormatter())
}

func randString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

//LogEntryWithRef returns a logrus Entry with a random unique value for requestId field
func LogEntryWithRef() *logrus.Entry {
	return logrus.WithField("reference", randString(10))
}

//TimestampFormatter returns a custom logrus Formatter with timestamp enabled
func TimestampFormatter() logrus.Formatter {
	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.PrettyPrint = false

	return formatter
}
