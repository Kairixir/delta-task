package main

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestHandler(test *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		test.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "^Hello World!\n.*"

	re := regexp.MustCompile(expected)

	if !re.MatchString(rr.Body.String()) {
		test.Errorf("handler returned unexpected body: got %v expected to match %v",
			rr.Body.String(), expected)
	}
}
