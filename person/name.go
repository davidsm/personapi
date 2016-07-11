package person

import (
	"github.com/davidsm/personapi/data"
	"math/rand"
	"time"
)

type gender int

const (
	GenderFemale gender = iota
	GenderMale
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

	firstName := firstNameSet[rand.Intn(len(firstNameSet))]
	lastName := data.LastNames[rand.Intn(len(data.LastNames))]
	return NewName(firstName, lastName)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
