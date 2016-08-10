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

func getPersonTest(url string, t *testing.T) *personResponse {
	response, err := getPerson(url)
	if err != nil {
		t.Fatal(err)
	}
	return response
}

func TestGetPerson(t *testing.T) {
	server := httptest.NewServer(app.Create())
	defer server.Close()

	response := getPersonTest(server.URL, t)
	testHeader("Content-Type", "application/json; charset=UTF-8", response.res.Header, t)
	// Default amount is 10
	testAmount(response.persons, 10, t)

	response = getPersonTest(server.URL+"?amount=20", t)
	testAmount(response.persons, 20, t)

	response = getPersonTest(server.URL+"?amount=150", t)
	testAmount(response.persons, 100, t)

	response = getPersonTest(server.URL+"?amount=-1", t)
	testAmount(response.persons, 1, t)
}
