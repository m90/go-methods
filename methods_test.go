package methods

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAllow(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		methods            []string
		handler            http.HandlerFunc
		expectedStatusCode int
	}{
		{
			"default",
			http.MethodPost,
			[]string{http.MethodPost},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK!"))
			},
			200,
		},
		{
			"no args",
			http.MethodDelete,
			[]string{},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK!"))
			},
			405,
		},
		{
			"block",
			http.MethodDelete,
			[]string{http.MethodGet, http.MethodPost},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK!"))
			},
			405,
		},
		{
			"error propagation",
			http.MethodGet,
			[]string{http.MethodGet},
			func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "zalgo", http.StatusInternalServerError)
			},
			500,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := Allow(test.methods...)(test.handler)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(test.method, "/", nil)
			handler.ServeHTTP(w, r)
			if w.Code != test.expectedStatusCode {
				t.Errorf("Expected status code %v, got %v", test.expectedStatusCode, w.Code)
			}
		})
	}
}
func TestDisallow(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		methods            []string
		handler            http.HandlerFunc
		expectedStatusCode int
	}{
		{
			"default",
			http.MethodPost,
			[]string{http.MethodGet},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK!"))
			},
			200,
		},
		{
			"no args",
			http.MethodPost,
			[]string{},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK!"))
			},
			200,
		},
		{
			"no args",
			http.MethodDelete,
			[]string{},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK!"))
			},
			200,
		},
		{
			"block",
			http.MethodDelete,
			[]string{http.MethodDelete, http.MethodPost},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK!"))
			},
			405,
		},
		{
			"error propagation",
			http.MethodGet,
			[]string{http.MethodPost},
			func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "zalgo", http.StatusInternalServerError)
			},
			500,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := Disallow(test.methods...)(test.handler)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(test.method, "/", nil)
			handler.ServeHTTP(w, r)
			if w.Code != test.expectedStatusCode {
				t.Errorf("Expected status code %v, got %v", test.expectedStatusCode, w.Code)
			}
		})
	}
}
