package app

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masenius/personapi/person"
)

const numberOfResults = 10

type Persons []person.Person

type PersonResponse struct {
	Result Persons `json:"result"`
	Amount int     `json:"amount"`
}

func CreatePerson() person.Person {
	gender := person.RandomGender()
	name := person.RandomName(gender)
	address := person.RandomAddress()
	birthDate := person.RandomBirthDate()
	idNumber := person.GenerateIdNumber(birthDate, gender)
	return person.Person{
		Name:          name,
		Age:           birthDate.Age(),
		BirthDate:     birthDate,
		IdNumber:      idNumber,
		Gender:        gender,
		PostalAddress: address,
	}
}

func HandleRequest(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	persons := make(Persons, 0, numberOfResults)
	for i := 0; i < numberOfResults; i++ {
		persons = append(persons, CreatePerson())
	}
	body := PersonResponse{Amount: numberOfResults, Result: persons}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(res).Encode(body); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}
