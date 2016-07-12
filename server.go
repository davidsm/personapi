package main

import (
	"encoding/json"
	"fmt"
	"github.com/masenius/personapi/person"
	"net/http"
)

const Amount = 10

type Persons []person.Person

type PersonResponse struct {
	Result Persons `json:"result"`
	Amount int     `json:"amount"`
}

func CreatePerson() person.Person {
	gender := person.RandomGender()
	name := person.RandomName(gender)
	address := person.RandomAddress()
	return person.Person{Name: name, Age: 25, Gender: gender, PostalAddress: address}
}

func handleRequest(res http.ResponseWriter, req *http.Request) {
	persons := make(Persons, 0, Amount)
	for i := 0; i < Amount; i++ {
		persons = append(persons, CreatePerson())
	}
	body := PersonResponse{Amount: Amount, Result: persons}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(res).Encode(body); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
