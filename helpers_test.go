package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router().ServeHTTP(rr, req)
	return rr
}

func clearTables() {
	tables = make([]*table, 0)
}
