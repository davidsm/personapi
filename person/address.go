package person

import (
	"strconv"

	"github.com/masenius/personapi/data"
)

type PostalAddress struct {
	Street       string `json:"streetAddress"`
	StreetNumber string `json:"streetNumber"`
	Code         string `json:"code"`
	Locality     string `json:"locality"`
}

type postalCode struct {
	Code     string `json:"code"`
	Locality string `json:"locality"`
}

type streetAddress struct {
	Street string `json:"street"`
	Number string `json:"number"`
}

func randomPostalCode() postalCode {
	pc := data.PostalCodes[randgen.Intn(len(data.PostalCodes))]
	return postalCode{pc.Code, pc.Locality}
}

func randomStreetAddress() streetAddress {
	return streetAddress{Street: randomStreetName(), Number: randomStreetNumber()}
}

func randomStreetName() string {
	return data.StreetNames[randgen.Intn(len(data.StreetNames))]
}

func randomStreetNumber() string {
	number := randgen.Intn(75)
	var letter string
	// Add a letter in 1/25 cases
	useLetterRoll := randgen.Intn(25)
	if useLetterRoll == 0 {
		// Use A-G as address letters.
		// Get the codepoint for one of these
		letterCode := randgen.Intn(7) + 65
		letter = string(letterCode)
	}
	return strconv.Itoa(number) + letter
}

func RandomAddress() *PostalAddress {
	pc := randomPostalCode()
	sa := randomStreetAddress()
	return &PostalAddress{Street: sa.Street, StreetNumber: sa.Number, Code: pc.Code, Locality: pc.Locality}
}
