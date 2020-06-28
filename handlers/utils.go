package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"timer.com/utils"

	"github.com/sirupsen/logrus"
)

type requestData struct {
	logger *logrus.Entry
	//db     *db.Conn
	w     http.ResponseWriter
	r     *http.Request
	start time.Time
}

type responseMessage struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

const jsonContentType = "application/json; charset=utf-8"

func badRequestIfNotMandatoryParams(key string, value string, rd *requestData) bool {
	if strings.TrimSpace(value) == "" {
		writeJSONMessage(fmt.Sprintf("%s is mandatory", key), http.StatusBadRequest, rd)
		return true
	}
	return false
}

func logAndGetRequestData(w http.ResponseWriter, r *http.Request) (*http.Request, *requestData) {
	ctx := r.Context()
	l, ok := ctx.Value("request-logger").(logrus.Entry)
	if !ok {
		l = *logrus.WithField("reference", utils.NewUUID())
		ctx = context.WithValue(ctx, "request-logger", l)
		r = r.WithContext(ctx)
		l.Info("Serving Request: ", r.RequestURI, ", Method:", r.Method)
	}

	//dbConn := db.NewConn(&l)

	return r, &requestData{
		logger: &l,
		//db:     dbConn,
		w:     w,
		r:     r,
		start: time.Now(),
	}
}

func writeJSONMessage(msg string, code int, rd *requestData) {
	writeJSONStruct(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		code,
		msg,
	}, code, rd)
}

func writeJSONStruct(v interface{}, code int, rd *requestData) {
	d, err := json.Marshal(v)
	if err != nil {
		writeResponse([]byte("Unable to marshal data. Error: "+err.Error()), http.StatusInternalServerError, jsonContentType, rd)
		return
	}
	writeResponse(d, code, jsonContentType, rd)
}

func writeResponse(d []byte, code int, contentType string, rd *requestData) {
	rd.logger = rd.logger.WithField("responseStatus", code).WithField("time", time.Since(rd.start))
	if code != http.StatusOK {
		rd.logger.WithField("response", string(d)).Info("Failure")
	} else {
		rd.logger.Info("Success")
	}

	rd.w.Header().Set("Access-Control-Allow-Origin", "*")
	rd.w.Header().Set("Content-Type", contentType)
	rd.w.WriteHeader(code) // WriteHeader should always be called last after setting all headers.
	rd.w.Write(d)
}
