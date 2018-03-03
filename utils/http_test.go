package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
}

func TestGetJSONWrongResponseStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer ts.Close()

	nsqdURL := ts.URL

	err := GetJSON(nsqdURL, "")
	if err == nil {
		t.Errorf("failed request")
	}
}

func TestGetJSONReturnJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defer ts.Close()
	resp := new(Response)
	expected := Response{Page: 0, TotalResults: 0}
	GetJSON(ts.URL, resp)
	if *resp != expected {
		t.Errorf("test failed: %+v\n != %+v\n", expected, resp)
	}
}
