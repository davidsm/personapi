package app

import (
	"net/url"
	"strconv"
)

type requestOptions struct {
	Amount int
}

const (
	paramAmount = "amount"
)

const (
	minAmount     = 1
	maxAmount     = 100
	defaultAmount = 10
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

func handleParams(params url.Values) *requestOptions {
	amount := intBetween(params.Get(paramAmount), minAmount, maxAmount, defaultAmount)
	return &requestOptions{Amount: amount}
}
