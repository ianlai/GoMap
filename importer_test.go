package main

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/ianlai/GoMap/data"
)

func TestRetrieveData(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		removeLength int64
		want         []string
		wantErr      bool
	}{
		{
			name:         "Case1: Success",
			url:          "https://bucket-ian-1.s3.amazonaws.com/data_prefix.txt",
			removeLength: 719,
			want: []string{
				"e7483dfb5ea31ac0cb4 1654",
				"e0296c3f3d00ef78fc5 1888",
			},
			wantErr: false,
		},
		{
			name:         "Case2: Failed - URL Format",
			url:          "htt://www.example.com",
			removeLength: 10,
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Case3: Failed - removeLength ",
			url:          "htt://www.example.com",
			removeLength: 1000,
			want:         nil,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RetrieveData(tt.url, tt.removeLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
			for i, _ := range got {
				if got[i] != tt.want[i] {
					t.Errorf("got = %v, want = %v", got[i], tt.want[i])
				}
			}
		})
	}
}

func TestRemovePrefixData(t *testing.T) {
	tests := []struct {
		name         string
		r            io.Reader
		removeLength int64
		want         io.Reader
		wantErr      bool
	}{
		{
			name:         "Case1: Success",
			r:            bytes.NewReader([]byte("hellotest")),
			removeLength: 3,
			want:         bytes.NewReader([]byte("test")),
			wantErr:      false,
		},
		{
			name:         "Case2: Failed",
			r:            bytes.NewReader([]byte("hellotest")),
			removeLength: 100, //too long
			want:         nil,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemovePrefixData(tt.r, tt.removeLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("got = %v, want %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
			}
		})
	}
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("error")
}
func TestGetLinesFromReader(t *testing.T) {
	tests := []struct {
		name    string
		r       io.Reader
		want    []string
		wantErr bool
	}{
		{
			name: "Case1: Success",
			r:    bytes.NewReader([]byte("aaa\nbbb\nccc")),
			want: []string{
				"aaa", "bbb", "ccc",
			},
			wantErr: false,
		},
		{
			name:    "Case2: Failed",
			r:       &errorReader{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gots, err := GetLinesFromReader(tt.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
			for i, got := range gots {
				if got != tt.want[i] {
					t.Errorf("got = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}
func TestInsertRecords(t *testing.T) {
	tests := []struct {
		name    string
		repo    data.Repo
		records []string
		wantErr bool
	}{
		{
			name: "Case1: Success",
			repo: &data.MockDB{
				MockInsertRecord: func(string, string) error {
					return nil
				},
			},
			records: []string{
				"successful input",
			},
			wantErr: false,
		},
		{
			name: "Case2: Failed",
			repo: &data.MockDB{
				MockInsertRecord: func(string, string) error {
					return errors.New("Mock DB insert error")
				},
			},
			records: []string{
				"failed input",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InsertRecords(tt.repo, tt.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTopKRecords(t *testing.T) {
	tests := []struct {
		name    string
		repo    data.Repo
		k       int
		want    []data.Record
		wantErr bool
	}{
		{
			name: "Case1: Success",
			repo: &data.MockDB{
				MockListRecords: func(int, bool) ([]data.Record, error) {
					//records should have descending values
					records := []data.Record{
						{
							Uid: "c1a812dda818edee076",
							Val: "1618",
						},
						{
							Uid: "a8f962e650be6e090d0",
							Val: "1595",
						},
					}
					return records, nil
				},
			},
			k: 2,
			want: []data.Record{
				{
					Uid: "c1a812dda818edee076",
					Val: "1618",
				},
				{
					Uid: "a8f962e650be6e090d0",
					Val: "1595",
				},
			},
			wantErr: false,
		},
		{
			name: "Case2: Failed",
			repo: &data.MockDB{
				MockListRecords: func(int, bool) ([]data.Record, error) {
					return nil, errors.New("Mock DB get error")
				},
			},
			k:       2,
			want:    []data.Record{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTopKRecords(tt.repo, tt.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
