package person

import (
	"math/rand"
	"time"
)

type Person struct {
	Name          *Name          `json:"name"`
	Age           int            `json:"age"`
	BirthDate     *BirthDate     `json:"birthDate"`
	IdNumber      string         `json:"idNumber"`
	Gender        gender         `json:"gender"`
	PostalAddress *PostalAddress `json:"postalAddress"`
}

var randgen *rand.Rand

func init() {
	randgen = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}
