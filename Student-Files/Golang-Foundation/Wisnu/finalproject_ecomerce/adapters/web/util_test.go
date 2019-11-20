package web

import (
	"bytes"
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/gorilla/mux"
)

type testCase struct {
	name               string
	url                string
	method             string
	body               []byte
	expectedStatusCode int
	expectedBody       []byte
}

func runHandlerTestCases(testCases []testCase, h http.Handler, t *testing.T) {
	for _, tc := range testCases {
		reqBody := bytes.NewReader(tc.body)
		r, err := http.NewRequest(tc.method, tc.url, reqBody)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		h.ServeHTTP(w, r)

		if w.Code != tc.expectedStatusCode {
			t.Errorf("%s %s (%s) expected status code %v got %v", tc.method, tc.url, tc.name, tc.expectedStatusCode, w.Code)
			t.Logf("%v", w.Body)
		}
	}
}

func TestQparam(t *testing.T) {
	req, err := http.NewRequest("GET", "/?first=a&2nd=b&5nd=e&last=z", nil)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		param    string
		expected string
	}{
		{"first", "a"},
		{"2nd", "b"},
		{"5nd", "e"},
		{"last", "z"},
		{"notexists", ""},
	}

	for _, test := range tests {
		qp := queryValue(test.param, req)
		if qp != test.expected {
			t.Errorf("Expected qParam(%q) to be %v got %v", test.param, test.expected, qp)
		}
	}

}

func TestMuxVarMustInt(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	mr := mux.NewRouter()
	mr.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := urlParamMustInt("id", r)
		if id != 1 {
			t.Errorf("Expected %d got %v", 1, id)
		}
	}).Methods("GET")

	mr.ServeHTTP(rr, req)
}
