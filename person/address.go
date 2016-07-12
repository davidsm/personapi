package person

import (
	"github.com/davidsm/personapi/data"
)

type PostalAddress struct {
	StreetAddress string `json:"streetAddress"`
	Code          string `json:"code"`
	Locality      string `json:"locality"`
}

type postalCode struct {
	Code     string `json:"code"`
	Locality string `json:"locality"`
}

func RandomPostalCode() postalCode {
	pc := data.PostalCodes[randgen.Intn(len(data.PostalCodes))]
	return postalCode{pc.Code, pc.Locality}
}

func RandomAddress() PostalAddress {
	pc := RandomPostalCode()
	return PostalAddress{StreetAddress: "Statiska gatan 1", Code: pc.Code, Locality: pc.Locality}
}
