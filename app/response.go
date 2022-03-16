package app

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string          `json:"status"`
	Error  *ResponseError  `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}
type ResponseError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

const (
	StatusOK   = "ok"
	StatusFail = "nok"
)

func Send(w http.ResponseWriter, status int, result interface{}) {
	rj, err := json.Marshal(result)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	r := &Response{
		Status: StatusOK,
		Result: rj,
	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}

// Fail ends an unsuccessful JSON response with the stardard failure format.
func Fail(w http.ResponseWriter, status, errCode int, details ...string) {
	// msg, ok := frErrMap[errCode]
	// if !ok {
	// 	errCode = status
	// 	msg = http.StatusText(status)
	// }

	errCode = status
	msg := http.StatusText(status)

	r := &Response{
		Status: StatusFail,
		Error: &ResponseError{
			Code:    errCode,
			Message: msg,
			Details: details,
		},
	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}
