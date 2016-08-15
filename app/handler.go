package app

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masenius/personapi/person"
)

type Persons []person.Person

type PersonResponse struct {
	Result Persons `json:"result"`
	Amount int     `json:"amount"`
}

func createPerson(reqOpts *requestOptions) person.Person {
	gender := person.RandomGender()
	name := person.RandomName(gender)
	address := person.RandomAddress()
	birthDate := person.RandomBirthDate(reqOpts.AgeFrom, reqOpts.AgeTo)
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

func handleRequest(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	params := req.URL.Query()
	reqOpts := handleParams(params)

	persons := make(Persons, 0, reqOpts.Amount)
	for i := 0; i < reqOpts.Amount; i++ {
		persons = append(persons, createPerson(reqOpts))
	}
	body := PersonResponse{Amount: reqOpts.Amount, Result: persons}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(res).Encode(body); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}
