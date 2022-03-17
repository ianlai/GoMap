package data

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertRecord(t *testing.T) {
	params := Record{
		Uid: "78eab4ccbdd98fa911e",
		Val: "1886",
	}
	query := `INSERT INTO map`
	tests := []struct {
		name        string
		mockClosure func(mock sqlmock.Sqlmock)
		wantErr     bool
	}{
		{
			name: "Case1: Successful",
			mockClosure: func(mock sqlmock.Sqlmock) {
				result := sqlmock.NewResult(1, 1)
				mock.ExpectExec(query).
					WithArgs(params.Uid, params.Val).
					WillReturnResult(result)
			},
			wantErr: false,
		},
		{
			name: "Case2: Failed",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(params.Uid, params.Val).
					WillReturnError(fmt.Errorf("error_occured"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.mockClosure(mock)
			customDB := &DB{db}
			gotErr := customDB.InsertRecord(params.Uid, params.Val)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestListRecords(t *testing.T) {
	param_k := 2
	tests := []struct {
		name          string
		mockClosure   func(mock sqlmock.Sqlmock)
		k             int
		isSortedByVal bool
		want          []Record
		wantErr       bool
	}{
		{
			name: "Case1: Successful (sorted by val)",
			mockClosure: func(mock sqlmock.Sqlmock) {
				result := sqlmock.NewRows([]string{"ID", "uid", "val"}).
					AddRow("1", "78eab4ccbdd98fa911e", "1886").
					AddRow("2", "0c2a9c4b0566e49721a", "877").
					AddRow("3", "aabbccdd11223344556", "455")
				mock.ExpectQuery(regexp.QuoteMeta(`
				SELECT 
					* 
				FROM 
					map 
				ORDER BY
					val DESC
				LIMIT $1`)).
					WithArgs(param_k).WillReturnRows(result)
			},
			k:             param_k,
			isSortedByVal: true,
			want: []Record{
				{
					ID:  1,
					Uid: "78eab4ccbdd98fa911e",
					Val: "1886",
				},
				{
					ID:  2,
					Uid: "0c2a9c4b0566e49721a",
					Val: "877",
				},
				{
					ID:  3,
					Uid: "aabbccdd11223344556",
					Val: "455",
				},
			},
			wantErr: false,
		},
		{
			name: "Case2: Successful (unsorted)",
			mockClosure: func(mock sqlmock.Sqlmock) {
				result := sqlmock.NewRows([]string{"ID", "uid", "val"}).
					AddRow("1", "78eab4ccbdd98fa911e", "15").
					AddRow("2", "0c2a9c4b0566e49721a", "877").
					AddRow("3", "aabbccdd11223344556", "66")
				mock.ExpectQuery(regexp.QuoteMeta(`
				SELECT 
					* 
				FROM 
					map 
				LIMIT $1`)).
					WithArgs(param_k).WillReturnRows(result)
			},
			k:             param_k,
			isSortedByVal: false,
			want: []Record{
				{
					ID:  1,
					Uid: "78eab4ccbdd98fa911e",
					Val: "15",
				},
				{
					ID:  2,
					Uid: "0c2a9c4b0566e49721a",
					Val: "877",
				},
				{
					ID:  3,
					Uid: "aabbccdd11223344556",
					Val: "66",
				},
			},
			wantErr: false,
		},
		{
			name: "Case3: Failed",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`
				SELECT 
					* 
				FROM 
					map 
				LIMIT $1`)).
					WithArgs(param_k).
					WillReturnError(errors.New("error_occured"))
			},
			k:             param_k,
			isSortedByVal: false,
			want:          nil,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.mockClosure(mock)
			customDB := &DB{db}
			gotRecords, gotErr := customDB.ListRecords(tt.k, tt.isSortedByVal)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("gotErr => %v, wantErr => %v", gotErr, tt.wantErr)
			}
			for i, got := range gotRecords {
				if got != tt.want[i] {
					t.Errorf("got => %v, want => %v", got, tt.want[i])
				}
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
