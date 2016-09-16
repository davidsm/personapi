package app

import (
	"net/url"
	"strconv"

	"github.com/masenius/personapi/person"
)

type requestOptions struct {
	Amount     int
	AgeFrom    int
	AgeTo      int
	OnlyGender *person.Gender
}

const (
	paramAmount = "amount"
	paramMinAge = "minAge"
	paramMaxAge = "maxAge"
	paramGender = "gender"
)

const (
	minAmount     = 1
	maxAmount     = 1000
	defaultAmount = 10

	minAge = 0
	maxAge = 105
)

// intBetween takes a param value (string), tries to convert it to an integer,
// and returns it if successful. If the value is not between min and max, it is rounded up/down as needed
// If the parameter is empty or can't be converted to an integer, the default value is returned
func intBetween(value string, min, max, defaultVal int) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		return defaultVal
	}
	if num < min {
		return min
	} else if num > max {
		return max
	}
	return num
}

func spanBetween(minValue, maxValue string, min, max int) (int, int) {
	numMin := intBetween(minValue, min, max, min)
	numMax := intBetween(maxValue, min, max, max)
	if numMin > numMax {
		numMin = min
		numMax = max
	}
	return numMin, numMax
}

func handleParams(params url.Values) *requestOptions {
	amount := intBetween(params.Get(paramAmount), minAmount, maxAmount, defaultAmount)
	ageFrom, ageTo := spanBetween(
		params.Get(paramMinAge),
		params.Get(paramMaxAge),
		minAge,
		maxAge,
	)
	onlyGender := person.Gender(params.Get(paramGender))
	onlyGenderPtr := &onlyGender
	if onlyGender != person.GenderMale && onlyGender != person.GenderFemale {
		onlyGenderPtr = nil
	}

	return &requestOptions{
		Amount:     amount,
		AgeFrom:    ageFrom,
		AgeTo:      ageTo,
		OnlyGender: onlyGenderPtr,
	}
}
