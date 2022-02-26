package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

// test http get (really need a url to download?)
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
			//if !reflect.DeepEqual(got, tt.expected) {
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("got = %v, want %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
			}
		})
	}
}

// can't create error
func TestGetLinesFromReader(t *testing.T) {
	tests := []struct {
		name    string
		r       io.Reader
		want    []string
		wantErr bool
	}{
		{
			name: "Case1: Success",
			r:    bytes.NewReader([]byte("aaa \nbbb\nccc")),
			want: []string{
				"aaa", "bbb", "ccc",
			},
			wantErr: false,
		},
		{
			name:    "Case2: Failed",
			r:       bytes.NewReader([]byte("")),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLinesFromReader(tt.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
			//if !reflect.DeepEqual(got, tt.expected) {
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("got = %v, want %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
			}
		})
	}
}

func TestInsertRecords(t *testing.T) {

}

func TestGetTopKthVal(t *testing.T) {

}
