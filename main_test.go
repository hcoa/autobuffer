package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"crypto/rand"
	"strconv"
)
const (
	testSz = 1000
	testFilename = "testout.mkv"
)
var (
	testData = make([]byte, testSz)
)

func TestNewVideoStream(t *testing.T) {
	_, err := rand.Read(testData)
	if err != nil {
		t.Fatal(err)
	}	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Length", strconv.Itoa(testSz))
		_, err := w.Write(testData)
		if err != nil {
			panic(err)
		}
	}))
	defer ts.Close()

	vs, err := NewVideoStream(ts.URL, "testout.mkv")
	if err != nil {
		t.Fatal(err)
	}
	if vs.Size != testSz {
		t.Fatalf("VideoStream created with wrong size, got %v wanted %v\n", vs.Size, testSz)
	}
}