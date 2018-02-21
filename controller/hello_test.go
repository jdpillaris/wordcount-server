package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	// In case there is an error in forming the request, we fail and stop the test
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(Hello)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler 'Hello' returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `Welcome to Ulventech word counter!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("Handler 'Hello' returned unexpected body: got %v want %v", actual, expected)
	}
}
