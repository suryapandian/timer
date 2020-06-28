package handlers

import (
	"net/http"
	"strconv"

	"timer.com/dtos"
	"timer.com/services"

	"github.com/go-chi/chi"
)

func setTimerRoutes(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Post("/_create", CreateTimer)
		r.Get("/_render", DisplayTimer)
		r.Get("/_check", CheckTimer)
		r.Post("/_clear/{id}", DeleteTimer)
		r.Post("/_pause/{id}", PauseTimer)
	})

}

func DisplayTimer(w http.ResponseWriter, r *http.Request) {
	//r, rd := logAndGetRequestData(w, r)
}

func CheckTimer(w http.ResponseWriter, r *http.Request) {
	r, rd := logAndGetRequestData(w, r)
	id := r.URL.Query().Get("id")
	if id != "" {
		timer, err := services.NewTimer(rd.logger).GetByID(id)
		switch err {
		case nil:
			writeJSONStruct(timer, http.StatusOK, rd)
			return
		case services.ErrTimerNotFound:
			writeJSONMessage(err.Error(), http.StatusNotFound, rd)
			return
		default:
			writeJSONMessage("Failed to get timer "+id, http.StatusInternalServerError, rd)
			return
		}
	}

	timers := services.NewTimer(rd.logger).GetAll()
	writeJSONStruct(timers, http.StatusOK, rd)
}

func CreateTimer(w http.ResponseWriter, r *http.Request) {
	r, rd := logAndGetRequestData(w, r)
	timer := &dtos.Timer{}
	startVal := r.URL.Query().Get("startVal")
	timer.Counter, _ = strconv.ParseFloat(startVal, 64)
	if badRequestIfNotMandatoryParams("Start Value", startVal, rd) {
		return
	}

	stepTime := r.URL.Query().Get("stepTime")
	timer.StepTime, _ = strconv.ParseFloat(stepTime, 64)
	if badRequestIfNotMandatoryParams("step Time", stepTime, rd) {
		return
	}

	id := services.NewTimer(rd.logger).Create(timer)
	writeJSONStruct(&responseMessage{id, "Successfully created timer", http.StatusOK},
		http.StatusOK, rd)
}

func DeleteTimer(w http.ResponseWriter, r *http.Request) {
	r, rd := logAndGetRequestData(w, r)
	id := chi.URLParam(r, "id")
	if badRequestIfNotMandatoryParams("id", id, rd) {
		return
	}
	err := services.NewTimer(rd.logger).Delete(id)
	switch err {
	case nil:
		writeJSONStruct(&responseMessage{id, "Successfully deleted timer", http.StatusOK},
			http.StatusOK, rd)
	case services.ErrTimerNotFound:
		writeJSONMessage(err.Error(), http.StatusNotFound, rd)
	default:
		writeJSONMessage("Failed to delete timer "+id, http.StatusInternalServerError, rd)
	}
}

func PauseTimer(w http.ResponseWriter, r *http.Request) {
	r, rd := logAndGetRequestData(w, r)
	id := chi.URLParam(r, "id")
	if badRequestIfNotMandatoryParams("id", id, rd) {
		return
	}
	timer, err := services.NewTimer(rd.logger).Pause(id)
	switch err {
	case nil:
		writeJSONStruct(timer, http.StatusOK, rd)
	case services.ErrTimerNotFound:
		writeJSONMessage(err.Error(), http.StatusNotFound, rd)
	default:
		writeJSONMessage("Failed to pause timer "+id, http.StatusInternalServerError, rd)
	}
}
