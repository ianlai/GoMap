package data

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUpdateMap(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("open a stub database failed: %s", err)
	}
	defer db.Close()
	rows := mock.NewRows([]string{"uid", "id", "val", "updated_at"}).
		AddRow(1, "c1a812dda818edee076", 1618, time.Now())
		//AddRow(2, "0c2a9c4b0566e49721a", 377, time.Now())
		//AddRow(3, "78eab4ccbdd98fa911e", 1886, time.Now())

	mock.ExpectQuery("INSERT INTO map (uid, val) VALUES ($1, $2)").WithArgs("78eab4ccbdd98fa911e", 1886).WillReturnRows(rows)
	err = UpdateMap(db, "https://bucket-ian-1.s3.amazonaws.com/data_short.txt")
	if err != nil {
		t.Errorf("Expected no error, but error occured: %s", err)
	}
}

func TestGetTopKth(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("open a stub database failed: %s", err)
	}
	defer db.Close()
	rows := mock.NewRows([]string{"uid", "id", "val", "updated_at"}).
		AddRow(1, "c1a812dda818edee076", 1618, time.Now())
		//AddRow(2, "0c2a9c4b0566e49721a", 377, time.Now())
		//AddRow(3, "78eab4ccbdd98fa911e", 1886, time.Now())

	mock.ExpectQuery("SELECT uid, val FROM map").WithArgs(2).WillReturnRows(rows)

	got := GetTopKth(db, 2)
	expected := "c1a812dda818edee076"
	if got != expected {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}
