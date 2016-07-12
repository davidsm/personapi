package person

import (
	"github.com/davidsm/personapi/data"
)

type Name struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewName(firstName, lastName string) Name {
	return Name{FirstName: firstName, LastName: lastName}
}

func RandomName(gender gender) Name {
	var firstNameSet []string
	switch gender {
	case GenderMale:
		firstNameSet = data.MaleNames
	case GenderFemale:
		firstNameSet = data.FemaleNames
	default:
		firstNameSet = data.FemaleNames
	}

	firstName := firstNameSet[randgen.Intn(len(firstNameSet))]
	lastName := data.LastNames[randgen.Intn(len(data.LastNames))]
	return NewName(firstName, lastName)
}
