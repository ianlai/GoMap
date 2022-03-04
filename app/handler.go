package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

func (s *Server) HandleShowStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current time is: %s\n", time.Now())
}
func (s *Server) HandleShowRecord(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	result, err := s.Repo.GetRecordByUid(uid)
	if err != nil {
		log.Printf("%v", errors.Wrap(err, "HandleShowRecord error"))
	}
	log.Printf("%v", result)

	rj, err := json.Marshal(result)
	if err != nil {
		log.Printf("%v", errors.Wrap(err, "HandleShowRecord error"))
	}
	w.Write(rj)
}
func (s *Server) HandleListRecords(w http.ResponseWriter, r *http.Request) {

	isSort := r.URL.Query().Get("sort")
	result, err := s.Repo.ListRecords(20, false)
	if isSort == "true" {
		result, err = s.Repo.ListRecords(20, true)
	}
	if err != nil {
		return
	}
	rj, err := json.Marshal(result)
	if err != nil {
		return
	}
	res := &Response{
		Status: "ok",
		Result: rj,
	}
	j, err := json.Marshal(res)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(status)
	w.Write(j)
}
