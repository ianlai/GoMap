package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetRouter(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Get("/status", handleShowStatus)
	r.Route("/records", func(r chi.Router) {
		r.Get("/{id}", handleShowRecord)
		r.Get("/", handleListRecord)
	})
}

func handleShowStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current time is: %s\n", time.Now())
}
func handleShowRecord(w http.ResponseWriter, r *http.Request) {
	val := chi.URLParam(r, "name")
	if val != "" {
		fmt.Fprintf(w, "Hello %s!", val)
	} else {
		fmt.Fprintf(w, "Hello no name.")
	}
}
func handleListRecord(w http.ResponseWriter, r *http.Request) {
	val := chi.URLParam(r, "name")
	if val != "" {
		fmt.Fprintf(w, "Hello %s!", val)
	} else {
		fmt.Fprintf(w, "Hello no name.")
	}
}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "root endpoint")
// })
// http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "status endpoint")
// })
// http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
// 	u := &UserInfo{
// 		Name: "Yoyo",
// 		Age:  28,
// 	}
// 	b, err := json.Marshal(u)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	w.Write(b)
// 	//fmt.Fprintf(w, "showjson endpoint")
// })
// http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
// 	users := []UserInfo{
// 		{
// 			Name: "Yoyo",
// 			Age:  28,
// 		}, {
// 			Name: "Quan",
// 			Age:  19,
// 		},
// 	}
// 	b, err := json.Marshal(users)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	w.Write(b)
// 	//fmt.Fprintf(w, "showjson endpoint")
// })
