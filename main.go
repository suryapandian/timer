package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"timer.com/handlers"
	"timer.com/utils"
)

func main() {
	utils.SetupLog()

	logrus.Infof("Listening on port %s", "3000")

	http.ListenAndServe(":"+"3000", handlers.GetRouter())
}
