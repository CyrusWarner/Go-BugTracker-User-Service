package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/user/register", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := executeRequest(req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserRegisterHandler)
	handler.ServeHTTP(rr, req)

	return rr
}
