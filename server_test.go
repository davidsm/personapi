package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/masenius/personapi/app"
)

type personResponse struct {
	res     *http.Response
	persons *app.PersonResponse
}

func getPerson(url string) (*personResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var persons app.PersonResponse
	if err := json.NewDecoder(res.Body).Decode(&persons); err != nil {
		return nil, err
	}
	return &personResponse{res, &persons}, nil
}

func testHeader(header, expected string, headers http.Header, t *testing.T) {
	actual := headers.Get(header)
	if actual != expected {
		t.Errorf("Header %s was %s, expected %s", header, actual, expected)
	}
}

func testAmount(persons *app.PersonResponse, expected int, t *testing.T) {
	if persons.Amount != expected {
		t.Errorf("Amount field was %d, expected %d", persons.Amount, expected)
	}

	if len(persons.Result) != expected {
		t.Errorf("Number of results was %d, expected %d",
			len(persons.Result), expected)
	}
}

func TestGetPerson(t *testing.T) {
	server := httptest.NewServer(app.Create())
	defer server.Close()

	response, err := getPerson(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	testHeader("Content-Type", "application/json; charset=UTF-8", response.res.Header, t)

	// Default amount is 10
	testAmount(response.persons, 10, t)

	response, err = getPerson(server.URL + "?amount=20")
	if err != nil {
		t.Fatal(err)
	}
	testAmount(response.persons, 20, t)
}
