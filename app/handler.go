package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/ianlai/GoMap/data"
	"github.com/pkg/errors"
)

func (s *Server) HandleShowStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current time: %s\n", time.Now())
}

func (s *Server) HandleCreateRecord(w http.ResponseWriter, r *http.Request) {

	var record data.Record
	err := json.NewDecoder(r.Body).Decode(&record)
	result, err := s.Repo.InsertRecord(record.Uid, record.Val)
	if err != nil {
		log.Printf("%v", errors.Wrap(err, "HandleCreateRecord error"))
		return
	}
	Send(w, http.StatusCreated, result)
}

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
