package main

import (
	"encoding/json"
	"github.com/masenius/personapi/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHeader(header, expected string, headers http.Header, t *testing.T) {
	actual := headers.Get(header)
	if actual != expected {
		t.Errorf("Expected %s to be %s, was %s", header, expected, actual)
	}
}

func TestDefaultReply(t *testing.T) {
	server := httptest.NewServer(app.Create())
	defer server.Close()

	res, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	testHeader("Content-Type", "application/json; charset=UTF-8", res.Header, t)

	var personResponse app.PersonResponse
	err = json.NewDecoder(res.Body).Decode(&personResponse)
	if err != nil {
		t.Fatal(err)
	}

	// 10 is temporarily hardcoded as the amount
	if personResponse.Amount != 10 {
		t.Error("Expected amount field to be 10, was", personResponse.Amount)
	}

	if len(personResponse.Result) != personResponse.Amount {
		t.Errorf("Expected number of results to be the same as the amount field (%d), was %d",
			personResponse.Amount, len(personResponse.Result))
	}
}
