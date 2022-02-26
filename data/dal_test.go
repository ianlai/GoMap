package data

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// 1. decide the query
// 2. insert beforehand
// 3. t.Fatal  t.Error
// 4. mock.ExpectationsWereMet()

func TestInsertRecord(t *testing.T) {
	params := Record{
		Uid: "78eab4ccbdd98fa911e",
		Val: "1886",
	}
	query := `INSERT INTO map (uid, val)`
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
			gotErr := InsertRecord(db, params.Uid, params.Val)
			if (err != nil) != tt.wantErr {
				t.Errorf("gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetRecordsSortedByVal(t *testing.T) {
	query := `SELECT 
				uid, val 
			FROM 
				map 
			ORDER BY 
				val 
			DESC LIMIT $1`
	tests := []struct {
		name        string
		mockClosure func(mock sqlmock.Sqlmock)
		k           int
		want        []Record
		wantErr     bool
	}{
		{
			name: "Case1: Successful",
			mockClosure: func(mock sqlmock.Sqlmock) {
				result := sqlmock.NewResult(2, 2)
				mock.ExpectExec(query).
					WillReturnResult(result)
			},
			k: 2,
			want: []Record{
				{
					Uid: "78eab4ccbdd98fa911e",
					Val: "1886",
				},
				{
					Uid: "78eab4ccbdd98fa911e",
					Val: "1886",
				},
			},
			wantErr: false,
		},
		{
			name: "Case2: Failed",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WillReturnError(fmt.Errorf("error_occured"))
			},
			k:       2,
			want:    nil,
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
			gotRecords, gotErr := GetRecordsSortedByVal(db, tt.k)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
			for i, got := range gotRecords {
				if got != tt.want[i] {
					t.Errorf("got = %v, want %v", got, tt.want[i])
				}
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
