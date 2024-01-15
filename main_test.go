package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRegisterShortenedUrl(t *testing.T) {
	tt := []struct {
		input    string
		expected URLs
	}{
		{input: "http://example.com/page/", expected: URLs{
			LongURL:  "http://example.com/page/",
			ShortURL: "https://localhost:8000/3b6qaefl67BhWe4aGZo8F5"}},
	}

	for _, tc := range tt {
		data, err := json.Marshal(tc.input)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/person", bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/person", RegisterShortenedUrl).Methods("POST")
		router.ServeHTTP(rr, req)

		if rr.Body.String() != tc.expected {
			t.Errorf("wrong response body for param %v: got %v want %v",
				tc.input, rr.Body.String(), tc.expected)
		}
	}
}
