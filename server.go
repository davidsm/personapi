package main

import (
	"encoding/json"
	"fmt"
	"github.com/davidsm/personapi/person"
	"net/http"
)

func handleRequest(res http.ResponseWriter, req *http.Request) {
	gender := person.RandomGender()
	name := person.RandomName(gender)
	address := person.RandomAddress()
	p := person.Person{Name: name, Age: 25, Gender: gender, PostalAddress: address}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(res).Encode(p); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
