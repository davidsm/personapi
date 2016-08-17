package person

import (
	"math/rand"
	"time"
)

type Person struct {
	Name      *Name      `json:"name"`
	Age       int        `json:"age"`
	BirthDate *BirthDate `json:"birthDate"`
	IdNumber  string     `json:"idNumber"`
	Gender    Gender     `json:"gender"`
	Address   *Address   `json:"address"`
}

var randgen *rand.Rand

// Seed sets an explicit seed for the random generator.
// Useful for reproducible tests.
func Seed(seed int64) {
	randgen.Seed(seed)
}

func init() {
	randgen = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}
