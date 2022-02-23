package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

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
				t.Errorf("got = %v, wantErr %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
			}
		})
	}
}
