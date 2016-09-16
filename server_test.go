package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/masenius/personapi/app"
	"github.com/masenius/personapi/reqlog"
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

func testGender(persons *app.PersonResponse, expected string, t *testing.T) {
	for _, p := range persons.Result {
		if string(p.Gender) != expected {
			t.Errorf("Gender was %s, expected %s", p.Gender, expected)
		}
	}
}

func getPersonTest(url string, t *testing.T) *personResponse {
	t.Log("Requesting", url)
	response, err := getPerson(url)
	if err != nil {
		t.Fatal(err)
	}
	return response
}

func TestSetSeed(t *testing.T) {
	// Test that you get the same "random" generation all the time
	// when setting the same seed.
	// A bit wobbly, as it would NOT produce reproducible results in the case of failure...
	var seed int64 = 12345
	appOpts := app.Options{
		Seed: &seed,
	}
	server := httptest.NewServer(app.Create(&appOpts))
	response1 := getPersonTest(server.URL, t)
	server.Close()

	server = httptest.NewServer(app.Create(&appOpts))
	response2 := getPersonTest(server.URL, t)
	server.Close()

	if !reflect.DeepEqual(response1.persons, response2.persons) {
		t.Errorf("Expected responses to be equal, got %v and %v", response1.persons, response2.persons)
	}
}

func TestGetPerson(t *testing.T) {
	var seed int64 = 12345
	appOpts := app.Options{
		Seed: &seed,
	}
	server := httptest.NewServer(app.Create(&appOpts))
	defer server.Close()

	response := getPersonTest(server.URL, t)
	testHeader("Content-Type", "application/json; charset=UTF-8", response.res.Header, t)
	testHeader("Access-Control-Allow-Origin", "*", response.res.Header, t)
	testHeader("Cache-Control", "no-cache", response.res.Header, t)

	// Default amount is 10
	testAmount(response.persons, 10, t)

	response = getPersonTest(server.URL+"?amount=20", t)
	testAmount(response.persons, 20, t)

	response = getPersonTest(server.URL+"?amount=1001", t)
	testAmount(response.persons, 1000, t)

	response = getPersonTest(server.URL+"?amount=-1", t)
	testAmount(response.persons, 1, t)

	response = getPersonTest(server.URL+"?gender=male", t)
	testGender(response.persons, "male", t)

	response = getPersonTest(server.URL+"?gender=female", t)
	testGender(response.persons, "female", t)
}

func TestFileLog(t *testing.T) {
	logFile, err := ioutil.TempFile("", "logtest")
	defer os.Remove(logFile.Name())

	if err != nil {
		t.Fatal(err)
	}

	logger, err := reqlog.File(logFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	appOpts := app.Options{
		Logger: logger,
	}

	server := httptest.NewServer(app.Create(&appOpts))
	defer server.Close()

	_ = getPersonTest(server.URL, t)

	logContents, err := ioutil.ReadFile(logFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if len(logContents) == 0 {
		t.Error("Logfile was empty")
	}
}
