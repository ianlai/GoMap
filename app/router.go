package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Server) SetRouter(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Get("/status", s.HandleShowStatus)
	r.Route("/records", func(r chi.Router) {
		r.Get("/", s.HandleListRecords)
		r.Get("/{uid}", s.HandleShowRecord)
	})
}

// func Send(w http.ResponseWriter, status int, result interface{}) {
// 	rj, err := json.Marshal(result)
// 	if err != nil {
// 		http.Error(
// 			w,
// 			http.StatusText(http.StatusInternalServerError),
// 			http.StatusInternalServerError,
// 		)
// 		return
// 	}
// 	r := &Response{
// 		Status: StatusOK,
// 		Result: rj,
// 	}
// 	j, err := json.Marshal(r)
// 	if err != nil {
// 		http.Error(
// 			w,
// 			http.StatusText(http.StatusInternalServerError),
// 			http.StatusInternalServerError,
// 		)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	w.Write(j)
// }

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
