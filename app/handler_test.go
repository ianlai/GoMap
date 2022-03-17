package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-chi/chi"
	"github.com/ianlai/GoMap/data"
)

var recordFixtures = map[string]*data.Record{
	"c1a812dda818edee076": &data.Record{
		ID:  1,
		Uid: "c1a812dda818edee076",
		Val: "1618",
	},
	"78eab4ccbdd98fa911e": &data.Record{
		ID:  2,
		Uid: "78eab4ccbdd98fa911e",
		Val: "1886",
	},
	"0c2a9c4b0566e49721a": &data.Record{
		ID:  3,
		Uid: "0c2a9c4b0566e49721a",
		Val: "877",
	},
}

// ReadResponse reads the response from ResponseRecorder and return a Response.
func ReadResponse(w *httptest.ResponseRecorder) (*Response, error) {
	r := &Response{}
	if err := json.Unmarshal(w.Body.Bytes(), r); err != nil {
		return nil, err
	}
	return r, nil
}

func AddChiCtx(r *http.Request, urlParams map[string]string) *http.Request {
	rctx := chi.NewRouteContext()
	for k, v := range urlParams {
		rctx.URLParams.Add(k, v)
	}
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func MustJsonMarshal(recordUid string) []byte {
	//fmt.Printf("marshal: %v", recordFixtures[recordUid])
	res, err := json.Marshal(recordFixtures[recordUid])
	if err != nil {
		fmt.Errorf("json marshal failed")
	}
	return res
}

func TestHandleShowRecord(t *testing.T) {
	s := &Server{
		Repo: &data.MockDB{
			MockGetRecordByUid: func(uid string) (*data.Record, error) {
				if uid == "c1a812dda818edee076" {
					record := &data.Record{
						ID:  1,
						Uid: "c1a812dda818edee076",
						Val: "1618",
					}
					return record, nil
				} else {
					return nil, fmt.Errorf("error_occured")
				}
			},
		},
		Name: "GoMap-Test-Server",
	}
	tests := []struct {
		name string
		uid  string
		want *Response
	}{
		{
			name: "Case1: Success (Found)",
			uid:  "c1a812dda818edee076",
			want: &Response{
				Status: StatusOK,
				Error:  nil,
				Result: MustJsonMarshal("c1a812dda818edee076"),
			},
		},
		{
			name: "Case2: Failed (Not found)",
			uid:  "xxxxx",
			want: &Response{
				Status: StatusFail,
				Error: &ResponseError{
					Code: http.StatusNotFound,
				},
				Result: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := "/records"
			r, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal("failed to create request", err)
			}
			r = AddChiCtx(r, map[string]string{"uid": tt.uid})

			w := httptest.NewRecorder()
			http.HandlerFunc(s.HandleShowRecord).ServeHTTP(w, r)
			response, err := ReadResponse(w)
			if err != nil {
				t.Error("failed to read the response from ResponseRecorder", err)
			}
			if tt.want.Status == StatusOK {
				got := &data.Record{}
				err = json.Unmarshal(response.Result, got)
				if err != nil {
					t.Error("failed to unmarshal []byte (got)", err)
				}
				wantRecord := &data.Record{}
				err = json.Unmarshal(tt.want.Result, wantRecord)
				if err != nil {
					t.Error("failed to unmarshal []byte (want)", err)
				}
				if !reflect.DeepEqual(got, wantRecord) {
					t.Errorf("got => %v, want => %v", got, wantRecord)
				}
			} else {
				gotErr := response.Error.Code
				if gotErr != tt.want.Error.Code {
					t.Errorf("gotErr => %v, wantErr => %v", gotErr, tt.want.Error.Code)
				}
			}
		})
	}
}
