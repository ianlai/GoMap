package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/ianlai/GoMap/data"
	"github.com/pkg/errors"
)

//TODO:
// add commit id
// add database ping
// return in JSON
func (s *Server) HandleShowStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current time: %s\n", time.Now())
}

//TODO:
// handle fail request (done)
func (s *Server) HandleShowRecord(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	result, err := s.Repo.GetRecordByUid(uid)
	if err != nil {
		Fail(w, http.StatusNotFound, "HandleShowRecord failed", "Can't find the resource...")
		log.Printf("%v", errors.Wrap(err, "HandleShowRecord error"))
		return
	}
	Send(w, http.StatusOK, result)
}

//TODO:
// handle the header
// add the send function (done)
func (s *Server) HandleListRecords(w http.ResponseWriter, r *http.Request) {
	isSort := r.URL.Query().Get("sort")
	var result []data.Record
	var err error
	if isSort == "true" {
		result, err = s.Repo.ListRecords(20, true)
	} else {
		result, err = s.Repo.ListRecords(20, false)
	}
	if err != nil {
		Fail(w, http.StatusNotFound, "HandleListRecords failed", "Something happened...")
		log.Printf("%v", errors.Wrap(err, "HandleListRecords error"))
		return
	}
	Send(w, http.StatusOK, result)
}
