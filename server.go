package main

import (
	"encoding/json"
	"fmt"
	"github.com/davidsm/personapi/person"
	"net/http"
)

func handleRequest(res http.ResponseWriter, req *http.Request) {
	name := person.RandomName(person.GenderMale)
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(res).Encode(name); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
