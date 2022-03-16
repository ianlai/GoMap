package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

//TODO:
// add commit id
// add database ping
// return in JSON
// handle fail request (done)
func (s *Server) HandleShowStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "!! The current time is: %s\n", time.Now())
}
func (s *Server) HandleShowRecord(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	result, err := s.Repo.GetRecordByUid(uid)
	if err != nil {
		Fail(w, 404, 404, "failed...")
		log.Printf("%v", errors.Wrap(err, "HandleShowRecord error"))
		return
	}
	log.Printf("%v", result)
	Send(w, http.StatusOK, result)
}

//TODO:
//handle the header
//add the send function (done)
func (s *Server) HandleListRecords(w http.ResponseWriter, r *http.Request) {

	isSort := r.URL.Query().Get("sort")
	result, err := s.Repo.ListRecords(20, false)
	if isSort == "true" {
		result, err = s.Repo.ListRecords(20, true)
	}
	if err != nil {
		return
	}
	Send(w, http.StatusOK, result)
}
