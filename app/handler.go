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
	var gender person.Gender
	if reqOpts.OnlyGender != nil {
		gender = *reqOpts.OnlyGender
	} else {
		gender = person.RandomGender()
	}
	name := person.RandomName(gender)
	address := person.RandomAddress()
	birthDate := person.RandomBirthDate(reqOpts.AgeFrom, reqOpts.AgeTo)
	idNumber := person.GenerateIdNumber(birthDate, gender)
	phoneNumber := person.RandomPhoneNumber()
	return person.Person{
		Name:        name,
		Age:         birthDate.Age(),
		BirthDate:   birthDate,
		IdNumber:    idNumber,
		Gender:      gender,
		Address:     address,
		PhoneNumber: phoneNumber,
	}
}

func createPersons(reqOpts *requestOptions) Persons {
	persons := make(Persons, 0, reqOpts.Amount)

	// Store generated id numbers to make sure that there are no duplicates
	idNums := make(map[string]bool)

	for i := 0; i < reqOpts.Amount; i++ {
		p := createPerson(reqOpts)

		// Recreate if the same idNumber has already been generated
		for idNums[p.IdNumber] {
			p = createPerson(reqOpts)
		}
		idNums[p.IdNumber] = true

		persons = append(persons, p)
	}
	return persons
}

func handleRequest(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	params := req.URL.Query()
	reqOpts := handleParams(params)

	persons := createPersons(reqOpts)
	body := PersonResponse{Amount: reqOpts.Amount, Result: persons}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(res).Encode(body); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}
